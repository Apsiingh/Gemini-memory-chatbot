# 🤖 Gemini Memory Chatbot

A lightweight, zero-dependency command-line chatbot built with Go that connects to the Google Gemini API. It maintains persistent chat memory across sessions using a local JSON file.

---

## ✨ Features

- **Persistent Conversation Memory**: Automatically saves chat history to `chat_history.json` so context is preserved between runs.
- **Context-Aware Responses**: Passes previous exchanges to the Gemini API for continuous multi-turn dialogue.
- **Zero Third-Party Dependencies**: Built purely with Go's standard library (`net/http`, `bufio`, `encoding/json`).
- **Automatic `.env` Loading**: Seamlessly parses `.env` on launch without external libraries.
- **Built-in Commands**:
  - `/reset` — Clears conversation history and starts a fresh context.
  - `/exit` — Gracefully exits the application.

---

## 📋 Prerequisites

- [Go](https://go.dev/dl/) installed (version 1.20 or newer).
- A [Google Gemini API Key](https://aistudio.google.com/app/apikey).

---

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/Apsiingh/Gemini-memory-chatbot.git
cd Gemini-memory-chatbot
```

### 2. Configure Environment Variables

Create a `.env` file in the project root (you can copy `.env.example`):

```bash
cp .env.example .env
```

Open `.env` and add your Gemini API key:

```env
GEMINI_API_KEY=your_gemini_api_key_here
GEMINI_MODEL=gemini-3.5-flash-lite
```

> **Note:** If `GEMINI_MODEL` is not specified, it defaults to `gemini-3.8-flash`.

---

## 💻 Usage

### Run Directly
To start the chatbot, run all `.go` files in the directory:

```bash
go run .
```

### Build Binary
To build a standalone executable:

**Windows:**
```powershell
go build -o gemini-chatbot.exe .
.\gemini-chatbot.exe
```

**macOS / Linux:**
```bash
go build -o gemini-chatbot .
./gemini-chatbot
```

---

## 💬 Chat Controls

| Command | Description |
| :--- | :--- |
| Any prompt | Sends your message to Gemini with previous context |
| `/reset` | Clears conversation history and resets memory |
| `/exit` | Exits the chat session |

---

## 📁 Project Structure

```
├── .env.example         # Template for environment variables
├── .gitignore           # Ignores .env, binaries, and chat history
├── chat_history.json    # Local JSON storage for conversation memory (auto-generated)
├── go.mod               # Go module definition
├── helper.go            # History storage, .env parser, and Gemini API client
├── main.go              # CLI loop, user input handling, and command processing
├── models.go            # Request/response structs and configuration constants
└── README.md            # Project documentation
```

---

## 🔒 Security

- Never commit your `.env` file containing your secret API key.
- `chat_history.json` and `.env` are excluded via `.gitignore` to prevent sensitive credentials or chat history from being pushed to version control.

---

## 📄 License

This project is open source and available under the [MIT License](LICENSE).
