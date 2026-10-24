package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if err := loadEnv(".env"); err != nil {
		fmt.Println("Warning: could not load .env file:", err)
	}

	history, err := loadHistory()
	if err != nil {
		fmt.Println("Could not load history:", err)
		return
	}

	fmt.Println("Gemini Memory Chatbot")
	fmt.Println("Commands: /exit to quit, /reset to clear history")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("You: ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "/exit":
			return
		case "/reset":
			history = []Message{}
			if err := saveHistory(history); err != nil {
				fmt.Println("Save error:", err)
			} else {
				fmt.Println("Conversation history cleared.")
			}
			continue
		case "":
			continue
		}

		history = append(history, Message{
			Role:  "user",
			Parts: []Part{{Text: input}},
		})

		ctx, cancel := context.WithTimeout(context.Background(), 65*time.Second)
		answer, err := askGemini(ctx, history)
		cancel()

		if err != nil {
			fmt.Println("Request failed:", err)
			history = history[:len(history)-1] // Remove failed message
			continue
		}

		fmt.Println("\nGemini:", answer)

		history = append(history, Message{
			Role:  "model",
			Parts: []Part{{Text: answer}},
		})

		if err := saveHistory(history); err != nil {
			fmt.Println("Warning: could not save history:", err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Input error:", err)
	}

}
