# MyPasswords CLI

MyPasswords is a Go-based command-line password manager that uses PAM authentication for system integration, GORM with SQLite for encrypted data storage, and provides both argument-based and interactive TUI modes using the Cobra CLI framework and readline.

**Always reference these instructions first and fallback to search or bash commands only when you encounter unexpected information that does not match the info here.**

## Working Effectively

### Prerequisites and Setup
- Install system dependencies:
  ```bash
  sudo apt-get update && sudo apt-get install -y libpam0g-dev libsqlcipher-dev
  ```
- Ensure Go 1.23+ is installed (`go version` should show 1.23 or later)
- **CRITICAL**: Always set `export CGO_ENABLED=1` before any Go build operations

### Building and Testing
- Download dependencies:
  ```bash
  go mod download
  ```
- **Build the application** (NEVER CANCEL - allow 2+ minutes for clean builds):
  ```bash
  export CGO_ENABLED=1
  go build -v .
  ```
  - Clean build time: ~51 seconds - NEVER CANCEL, set timeout to 120+ seconds
  - Incremental build time: ~1 second
  - Creates `mypasswords` executable binary
- **Run tests**:
  ```bash
  go test ./...
  ```
  - Currently no test files exist in the codebase
- **Format and lint code**:
  ```bash
  go fmt ./...
  go vet ./...
  ```
- **Advanced linting** (install if needed):
  ```bash
  go install golang.org/x/tools/cmd/goimports@latest
  go install honnef.co/go/tools/cmd/staticcheck@latest
  ~/go/bin/goimports -d .
  ~/go/bin/staticcheck ./...
  ```

### Application Architecture
- **Entry Point**: `main.go` - detects argument vs interactive mode
- **Authentication**: `auth/` - PAM-based authentication with 15-minute session caching
- **Database**: `store/` - GORM models (User, Password) with SQLite backend
- **CLI Framework**: `commands/` - Cobra command structure
- **Interactive Shell**: `cli/` - readline-based TUI with shell commands
- **Data Location**: `~/.local/mypasswords/db/mypasswords.db` (SQLite database)
- **Session Cache**: `~/.local/mypasswords/session/` (PAM session tokens)

### Application Modes
1. **Argument Mode**: `./mypasswords <command> <args>`
   - `./mypasswords add <service> <user> <password>` - Add password entry
   - `./mypasswords list` - List all password entries
2. **Interactive Mode**: `./mypasswords` (no arguments)
   - Launches TUI with readline interface
   - Available commands: `exit`, `clear`
   - Uses `@username -> ` prompt format

### Key Files and Locations
```
/home/runner/work/MyPasswords/MyPasswords/
├── main.go                    # Application entry point
├── go.mod                     # Go module dependencies
├── setup.sh                   # System dependency installation script
├── auth/                      # PAM authentication
│   ├── pam.go                # PAM integration and password prompts
│   └── session.go            # Session caching (15-min expiry)
├── store/                     # Database layer
│   ├── database.go           # GORM database connection and config
│   ├── passwords.go          # Password model and CRUD operations
│   └── user.go               # User model and authentication
├── commands/                  # CLI command structure
│   ├── root.go               # Main cobra command setup
│   └── shell/                # Interactive shell commands
│       ├── shell.go          # Shell command framework
│       └── basic.go          # Basic commands (exit, clear)
└── cli/                       # Interactive TUI
    ├── cli.go                # Main CLI loop and readline integration
    └── menu.go               # CLI configuration and input filtering
```

## Validation and Testing

### Manual Validation Requirements
**CRITICAL**: The application requires PAM authentication for ALL operations, including help commands. This makes traditional testing challenging in automated environments.

- **Authentication Test**: Run `./mypasswords --help` to verify PAM authentication prompts appear
- **Build Validation**: Successful compilation creates a ~13MB `mypasswords` binary
- **Directory Creation**: First run creates `~/.local/mypasswords/` directory structure
- **Database Setup**: Application creates SQLite database on first connection

### Complete Validation Workflow
After making any changes to the codebase, run this complete validation sequence:

```bash
# 1. Clean build with timing validation (NEVER CANCEL - allow 2+ minutes)
export CGO_ENABLED=1
rm -f mypasswords
time go build -v .

# 2. Verify binary creation and size
ls -lh mypasswords  # Should show ~13MB binary

# 3. Format and lint validation
go fmt ./...
go vet ./...

# 4. Advanced linting (install tools if needed)
go install golang.org/x/tools/cmd/goimports@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
~/go/bin/goimports -d .
~/go/bin/staticcheck ./...

# 5. Test compilation chain
export CGO_ENABLED=1 && go build -v . && echo "Build validation: PASSED"
```

### Expected Validation Results
- **Clean build time**: 51 seconds (on fresh environment)
- **Incremental build time**: ~1 second
- **Binary size**: ~13MB (ELF 64-bit LSB executable)
- **go fmt**: No output (already formatted)
- **go vet**: No output (no issues found)
- **staticcheck**: May report code quality issues (e.g., "error strings should not be capitalized")

### Common Validation Steps
- **After making changes**: Always run build, format, and lint:
  ```bash
  export CGO_ENABLED=1
  go build -v . && go fmt ./... && go vet ./...
  ```
- **Check for unused functions**: Run `staticcheck` to identify dead code
- **Database migration**: Note that `Migrate()` function exists in `store/database.go` but is not currently called
- **Error handling**: All database operations return proper error types

### Known Issues and Limitations
- **No existing tests**: The codebase currently has no test files
- **PAM dependency**: Cannot test application functionality without valid PAM authentication  
- **Database encryption**: SQLCipher support exists but is currently disabled in favor of plain SQLite
- **Migration not called**: `store/database.go` has a `Migrate()` function that is defined but never invoked
- **Code quality**: `staticcheck` identifies issues like improper error string capitalization in `cli/cli.go:29`

## Common Development Tasks

### Adding New Commands
- Interactive commands: Add to `commands/shell/shell.go` in `PrepareCommands()`
- Argument commands: Modify switch statement in `main.go`
- Command implementations: Add functions to `commands/shell/basic.go` or create new files

### Database Changes
- Models: Modify structs in `store/user.go` or `store/passwords.go`  
- Operations: Add CRUD functions following existing patterns
- **Remember**: Migration function exists but isn't called - may need manual invocation

### Authentication Modifications
- PAM config: Modify `auth/pam.go` for authentication behavior
- Session duration: Change `sessionDuration` constant in `auth/session.go` (currently 15 minutes)
- Session storage: Files stored in `~/.local/mypasswords/session/`

### CLI Interface Changes
- Prompt format: Modify in `cli/menu.go` `GetConfig()` function
- Input filtering: Modify `filterInput()` in `cli/menu.go`
- Readline behavior: Configure in `cli/cli.go` `NewCli()` function

## Build Timing and Resource Requirements
- **Clean build**: 51 seconds (set timeouts to 120+ seconds)
- **Incremental build**: 1 second  
- **Dependencies download**: ~30 seconds first time
- **Linting**: ~2 seconds
- **Binary size**: ~13MB after compilation
- **Memory usage**: Minimal during build, PAM authentication requires system integration

**NEVER CANCEL builds or long-running commands - builds may take up to 2 minutes on clean environments.**

## Troubleshooting Common Issues

### Build Failures
- **"CGO_ENABLED not set"**: Always run `export CGO_ENABLED=1` before building
- **"libpam0g-dev not found"**: Install with `sudo apt-get install -y libpam0g-dev libsqlcipher-dev`
- **"multiple definition" errors**: Usually indicates conflicting SQLite drivers - check imports
- **"syntax error"**: Check for duplicate or malformed code lines, especially in database functions

### Runtime Issues  
- **"Authentication failed"**: PAM authentication is required for all operations
- **"Failed to connect to database"**: Database will be created automatically in `~/.local/mypasswords/`
- **Application hangs**: PAM authentication is waiting for password input

### Code Quality Issues
- **"error strings should not be capitalized"**: Fix error messages in format functions
- **"declared and not used"**: Remove unused variables and imports
- **"failed to migrate database"**: The `Migrate()` function exists but is not called automatically

## Repository-Specific Knowledge

### Output from Common Commands
```bash
# Repository structure
ls -la
total 60
drwxr-xr-x  8 runner docker  4096 Aug 16 23:34 .
drwxr-xr-x  3 runner docker  4096 Aug 16 23:17 ..
drwxr-xr-x  8 runner docker  4096 Aug 16 23:34 .git
drwxr-xr-x  2 runner docker  4096 Aug 16 23:34 .github
-rw-r--r--  1 runner docker    50 Aug 16 23:17 .gitignore
-rw-r--r--  1 runner docker  1071 Aug 16 23:17 LICENSE
-rw-r--r--  1 runner docker    26 Aug 16 23:17 README.md
drwxr-xr-x  2 runner docker  4096 Aug 16 23:17 auth
drwxr-xr-x  2 runner docker  4096 Aug 16 23:17 cli
drwxr-xr-x  3 runner docker  4096 Aug 16 23:17 commands
-rw-r--r--  1 runner docker   720 Aug 16 23:24 go.mod
-rw-r--r--  1 runner docker  2884 Aug 16 23:17 go.sum
-rw-r--r--  1 runner docker  1008 Aug 16 23:17 main.go
-rwxr-xr-x  1 runner docker 13466864 Aug 16 23:34 mypasswords
-rw-r--r--  1 runner docker    72 Aug 16 23:17 setup.sh
drwxr-xr-x  2 runner docker  4096 Aug 16 23:17 store

# Go files count and structure  
find . -name "*.go" | wc -l
11

# All Go file names
find . -name "*.go" -exec basename {} \; | sort | uniq
basic.go
cli.go
database.go
main.go
menu.go
pam.go
passwords.go
root.go
session.go
shell.go
user.go
```

### Key Dependencies from go.mod
```
module mypasswords
go 1.23.6

require (
    github.com/chzyer/readline v1.5.1          # Interactive CLI
    github.com/mattn/go-shellwords v1.0.12     # Shell command parsing
    github.com/msteinert/pam v1.2.0            # PAM authentication
    github.com/spf13/cobra v1.8.1              # CLI framework
    golang.org/x/term v0.33.0                  # Terminal control
    gorm.io/driver/sqlite v1.6.0               # SQLite database driver
    gorm.io/gorm v1.30.1                       # ORM framework
)
```

This knowledge helps reduce unnecessary exploration commands and provides context for common development scenarios.