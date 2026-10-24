package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// loadEnv reads key=value pairs from a .env file and sets them in the environment.
func loadEnv(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			val := strings.TrimSpace(parts[1])
			val = strings.Trim(val, `"'`)
			if os.Getenv(key) == "" {
				os.Setenv(key, val)
			}
		}
	}
	return scanner.Err()
}

// loadHistory reads saved past conversations from a file.

func loadHistory() ([]Message, error) {
	data, err := os.ReadFile(historyFile)

	if os.IsNotExist(err) {
		return []Message{}, nil
	}

	if err != nil {
		return nil, err
	}

	var history []Message
	if err := json.Unmarshal(data, &history); err != nil {
		return nil, err
	}

	return history, nil
}

func saveHistory(history []Message) error {
	data, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(historyFile, data, 0644)
}

// askGemini sends the chat history to the Gemini API and gets a reply.
func askGemini(ctx context.Context, history []Message) (string, error) {

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return "Error: GEMINI_API_KEY not set", fmt.Errorf("missing API key")
	}

	// Prepare the data using the structs from models.go
	body, err := json.Marshal(APIRequest{Contents: history})
	if err != nil {
		return "Error: Failed to prepare request", err
	}

	model := os.Getenv("GEMINI_MODEL")
	if model == "" {
		model = modelName
	}
	url := "https://generativelanguage.googleapis.com/v1beta/models/" + model + ":generateContent"

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return "Error: Failed to create HTTP request", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-goog-api-key", apiKey)

	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Error: No response found", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	var result APIResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", err
	}

	if len(result.Candidates) == 0 {
		return "", fmt.Errorf("Gemini returned no response candidates")
	}

	var answer strings.Builder
	for _, part := range result.Candidates[0].Content.Parts {
		answer.WriteString(part.Text)
	}
	if answer.Len() == 0 {
		return "", fmt.Errorf("Gemini returned no text")
	}

	return answer.String(), nil

}
