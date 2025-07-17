# Available Billing Commands

## Command List

### 1. `billing` or `billing up`
Run the billing server in foreground mode

### 2. `billing up -d`
Run the billing server in background daemon mode

### 3. `billing stop`
Stop the running billing server

### 4. `billing version`
Display version information

### 5. `billing show-users` or `billing show_users`
Show current user status including:
- Login users (authenticated but not in game)
- Online users (currently in game)
- IP counters (connections per IP)
- Active connections (with last activity timestamps)

### 6. `billing show`
**NEW**: Display comprehensive IP address information in table format:
- IP Address
- Number of connections per IP
- Number of unique accounts per IP
- List of account names per IP
- Summary totals

## Global Options
- `--log-path value`: Specify billing log file path
- `--help, -h`: Show help information

## Example Usage
```bash
# Start server in foreground
./billing

# Start server in background
./billing up -d

# Check IP statistics
./billing show

# Check user status
./billing show-users

# Stop server
./billing stop
```