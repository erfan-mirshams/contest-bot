# Math Contest Bot

A Telegram bot for hosting math contests, written in Go.

## Features (Planned)

- User login and authentication
- Problems with different difficulty categories
- Purchase system for problems
- Reward system for correct solutions
- PDF-based problem statements and submissions
- Admin panel for managing contests, problems, and users
- Judge panel for evaluating submissions
- YAML import/export for contest data

## Getting Started

### Prerequisites

- Go 1.24 or higher

### Running the Echo Server

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/contest-bot.git
   cd contest-bot
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Run the server:
   ```sh
   go run cmd/bot/main.go
   ```

4. Test the echo endpoint:
   ```sh
   curl -X POST -d "Hello, world!" http://localhost:8080/echo
   ```

## Project Structure

```
mathcontestbot/
├── cmd/                 # Application entry points
│   └── bot/             # Telegram bot main application
├── internal/            # Private application code
│   ├── config/          # Configuration using Viper
│   ├── db/              # Database access with SQLC
│   ├── handler/         # Telegram bot command handlers
│   ├── model/           # Domain models
│   ├── service/         # Business logic
│   ├── storage/         # PDF storage 
│   └── telegram/        # Telegram bot setup
├── pkg/                 # Public libraries
│   ├── logger/          # Logging with slog
│   └── yaml/            # YAML import/export
```

## Configuration

The application can be configured using a `config.yaml` file in the root directory:

```yaml
server:
  port: 8080
log:
  level: info
```

Environment variables can also be used to override the configuration file settings:

```sh
CONTEST_BOT_SERVER_PORT=9090
CONTEST_BOT_LOG_LEVEL=debug
```

Environment variables take precedence over the YAML configuration.

## License

This project is licensed under the MIT License - see the LICENSE file for details. 