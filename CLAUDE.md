# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

billing_go is a billing verification server written in Go, used for game account authentication and billing management. It supports both classic and retro billing types.

## Development Commands

### Build Commands

**Linux:**
```bash
# Build for current platform
make

# Clean build artifacts
make clean

# Build and package 32-bit version
make x32

# Build and package 64-bit version  
make x64

# Build all architectures
make all
```

**Windows:**
```batch
# Double-click to run
build.bat
```

### Code Quality Checks
```bash
# Install golint if not installed
go install golang.org/x/lint/golint@latest

# Run lint check
golint ./...

# Run vet check
go vet ./...

# Format code
go fmt ./...
```

### Running the Service
```bash
# Foreground mode
./billing

# Background daemon mode
./billing up -d

# Stop service
./billing stop

# View version info
./billing version

# Show online users
./billing show-users
```

### Testing

This project does not have unit tests. After making changes:
1. Build the project using `make` (Linux) or `build.bat` (Windows)
2. Manually test by starting the billing service
3. Check `billing.log` for any errors

## Code Architecture

The project follows a clean architecture pattern with the following structure:

- **cmd/** - CLI command implementations using urfave/cli framework
  - `app_command.go` - Main command entry point
  - Individual command files (up.go, stop.go, version.go, etc.)

- **services/** - Core service logic
  - `billing/` - Billing server implementation
  - `handle/` - Connection handling
  - Various service utility functions

- **models/** - Database models and operations
  - Account-related operations (login, registration, points)
  - All database interactions are centralized here

- **common/** - Shared type definitions
  - `billing_packet.go` - Packet definitions
  - `server_config.go` - Server configuration
  - `client_info.go` - Client information

- **bhandler/** - Business handlers
  - Each handler processes specific request types
  - Examples: login_handler.go, register_handler.go, query_point_handler.go

## Key Technical Details

### Database
- Uses MySQL with support for old password authentication
- Database operations are centralized in the `models` package
- Passwords are stored as MD5 hashes (for game compatibility)

### Network Protocol
- TCP server listening on configurable port
- Mixed Big-Endian and Little-Endian byte order
- Fixed packet header structure
- Supports multiple packet types (login, query, conversion, etc.)

### Configuration
- Supports both YAML and JSON formats (YAML preferred)
- Config file must be in the same directory as the executable
- Key configurations: IP, port, database connection, auto-registration, IP whitelist

### Logging
- Uses `go.uber.org/zap` for structured logging
- Log file: `billing.log` in the program directory
- Includes timestamp, level, message, and structured fields

### Encoding
- UTF-8 encoding throughout the codebase
- Supports GBK encoding conversion for game client compatibility

## Code Style Guidelines

- **Package names**: Lowercase, words directly connected (e.g., `bhandler`, `services`)
- **Functions/Methods**: PascalCase (e.g., `NewServer`, `LoadServerConfig`)
- **Variables**: camelCase (e.g., `serverConfig`, `clientData`)
- **Comments**: Use Chinese comments explaining function purpose and important logic
- **Exported functions**: Must have comments in format `// FunctionName 功能说明`
- **Error handling**: Standard Go pattern with Chinese error descriptions

## Security Considerations

- Never hardcode passwords in the code
- Support IP whitelist for connection restrictions
- Configuration file passwords should be quoted if containing special characters
- No sensitive information should be logged or committed

## Recent Changes (IP-based Player Limiting)

### Work Completed:
1. **Changed from MAC to IP-based limiting**:
   - Modified `common/handler_resource.go`: Changed `MacCounters` to `IPCounters` map
   - Added `ConnectionInfo` struct to track active connections with last activity time
   - Added `ActiveConnections` map to `HandlerResource` for connection health monitoring

2. **Updated configuration**:
   - Changed `pc_max_client_count` to `ip_max_client_count` in `server_config.go`
   - Updated `LoginHandler` struct to use `IPMaxClientCount` instead of `PcMaxClientCount`
   - Updated default configuration in `services/load_server_config.go`

3. **Modified login handler**:
   - Updated IP-based limiting logic in `login_handler.go`
   - Added IP counter increment on successful login
   - Added active connection tracking with timestamp

4. **Updated mark_online.go** (Task 1 - COMPLETED):
   - Changed function signature to accept `ipCounters` and `activeConnections` instead of `macCounters`
   - Updated to increment IP counter instead of MAC counter when players enter the game
   - Added logic to update active connection timestamp
   - Updated all 7 files that call `markOnline` function

5. **Initialized new maps** (Task 5 - COMPLETED):
   - Updated `services/billing/load_handlers.go` to initialize `IPCounters` and `ActiveConnections` maps
   - Removed old `MacCounters` initialization

6. **Partial updates to maintain compilation**:
   - Updated `logout_handler.go` to use IP counters (basic implementation, full Task 2 still pending)
   - Updated `command_handler.go` to display IP counters instead of MAC counters (basic implementation, full Task 4 still pending)

### Work Remaining:
1. **Complete logout_handler.go updates** (Task 2):
   - Ensure proper cleanup of IP counters and active connections
   - Handle edge cases for connection cleanup

2. **Implement connection health check** (Task 3):
   - Add periodic goroutine to check for stale connections
   - Remove disconnected clients from IP counters and active connections
   - Handle cases where Windows clients disconnect without sending logout
   - Make timeout configurable

3. **Complete command_handler.go updates** (Task 4):
   - Display active connections with their last activity time
   - Improve formatting of the show-users command output

### Configuration Change Required:
Users need to update their config files:
```yaml
# Old configuration
pc_max_client_count: 3

# New configuration
ip_max_client_count: 3
connection_timeout: 300  # seconds (optional, for health check)
```