# MyPasswords

A secure CLI password manager written in Go with PAM authentication and encrypted storage.

## Features

- **Secure Authentication**: Uses PAM (Pluggable Authentication Modules) for system-level authentication
- **Encrypted Storage**: SQLite database with SQLCipher encryption for password storage
- **Dual Interface**: Support for both command-line arguments and interactive TUI mode
- **Session Management**: Automatic session timeout for enhanced security (15-minute sessions)
- **Local Storage**: Passwords stored locally in `~/.local/mypasswords/`

## Prerequisites

### System Dependencies
```bash
sudo apt-get install libpam0g-dev libsqlcipher-dev
```

### Go Requirements
- Go 1.23.6 or later
- CGO enabled (required for PAM and SQLCipher)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/PeterSaletra/MyPasswords.git
cd MyPasswords
```

2. Install system dependencies:
```bash
chmod +x setup.sh
./setup.sh
```

3. Build the application:
```bash
export CGO_ENABLED=1
go mod tidy
go build -o mypasswords .
```

## Usage

MyPasswords supports two modes of operation:

### Command Line Mode

Add a password:
```bash
./mypasswords add <service> <username> <password>
```

List stored passwords:
```bash
./mypasswords list
```

### Interactive Mode

Run without arguments to enter interactive mode:
```bash
./mypasswords
```

This launches a TUI (Text User Interface) with the following commands:
- `exit` - Exit the application
- `clear` - Clear the screen

## Authentication

MyPasswords integrates with your system's PAM authentication. You'll be prompted to enter your system password when:
- Starting the application for the first time
- When your session expires (after 15 minutes of inactivity)

## Database Location

- Database: `~/.local/mypasswords/db/mypasswords.db`
- Session files: `~/.local/mypasswords/session/`

## Security Features

- **PAM Integration**: Leverages system authentication
- **SQLCipher Encryption**: Database is encrypted at rest
- **Session Timeout**: Automatic logout after 15 minutes
- **Secure File Permissions**: Database and session files use restrictive permissions (0600)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
