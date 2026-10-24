package main

// Constants to set up the Gemini model, file name for saving chats, and the API web address.
const (
	modelName   = "gemini-3.8-flash"
	historyFile = "chat_history.json"
	apiURL      = "https://generativelanguage.googleapis.com/v1beta/models/" + modelName + ":generateContent"
)

// Part holds the actual text of a message.
type Part struct {
	Text string `json:"text"`
}

// Message holds who is speaking ("user" or "model") and their text.
type Message struct {
	Role  string `json:"role"`
	Parts []Part `json:"parts"`
}

// APIRequest is the full package of data we send to Gemini.
type APIRequest struct {
	Contents []Message `json:"contents"`
}

// Candidate holds a single response option from Gemini.
type Candidate struct {
	Content Message `json:"content"`
}

// APIResponse is the full package of data Gemini sends back to us.
type APIResponse struct {
	Candidates []Candidate `json:"candidates"`
}
