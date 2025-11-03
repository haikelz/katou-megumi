<p align="center">
<img src="pkg/assets/katou-megumi.gif" alt="katou megumi gif" />
</p>

<h1 align="center">Katou Megumi Discord Bot</h1>

<p align="center">A simple Discord Bot</p>

## Features

### Commands

- **`!info`** - Display bot information and command list
- **`!ping`** - Check bot response time
- **`!salam`** - Send greetings to users
- **`!ask <question>`** - Ask questions using Google Gemini AI (2.5 pro)
- **`!jadwalsholat <city>`** - Get prayer times for a specific city
- **`!doa`** - Display daily prayers and supplications
- **`!asmaulhusna`** - Show the 99 Names of Allah (Asmaul Husna)
- **`!jokes`** - Share random jokes
- **`!quote`** - Display inspirational quotes
- **`!editbackground`** - Edit profile background images
- **`!shutdown`** - Turn off the bot

## Getting Started

### Prerequisites

- Go 1.23.1 or higher
- Discord Bot Token
- Google Gemini API Key

### Installation

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd katou-megumi
   ```

2. **Install dependencies**

   ```bash
   go mod download
   ```

3. **Set up environment variables**
   Create a `.env` file in the root directory:

   ```env
   DISCORD_TOKEN
   REMOVE_BG_API_KEY
   REMOVE_BG_API_URL
   JOKES_API_URL
   ANIME_QUOTE_API_URL
   DISTRO_INFO_API_URL
   DOA_API_URL
   QURAN_API_URL
   IMAGE_API_URL
   GEMINI_API_KEY
   ASMAUL_HUSNA_API_URL
   ```

4. **Run the bot**

   ```bash
   air -c .air.toml
   ```

### Discord Bot Setup

1. Create a Discord application at [Discord Developer Portal](https://discord.com/developers/applications)
2. Create a bot for your application
3. Copy the bot token and add it to your `.env` file
4. Invite the bot to your server with appropriate permissions

### Google Gemini AI Setup

1. Get an API key from [Google AI Studio](https://makersuite.google.com/app/apikey)
2. Add the API key to your `.env` file

## 🛠️ Development

### Running the Project Using Air

```
air -c .air.toml
```

### Building the Project

1. `go build`

```bash
go build -o bin/bot cmd/app/main.go
```

2. Docker

```bash
docker buildx bake && docker compose up -d
```

### Running Tests

```bash
go test ./...
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.

**Created by:** [haikelz](https://github.com/haikelz/)
