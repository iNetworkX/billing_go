# billing_go

A billing verification server written in Go for game account authentication and billing management.

## System Requirements

- Linux (Kernel version 2.6.32 or higher)

## Bug Reports

If you encounter issues using this program, please submit an Issue describing the problem and attach relevant log files.

**billing** logs are located in the same directory as the billing program, with the filename `billing.log`. This file is automatically created on first run.

You'll also need **Login** server logs, as only the Login server connects to the billing server.

You can modify the run script `run.sh` to write Login server logs to a file for easier debugging:

```sh
# Method for recording Login server logs

# Original command
./Login >/dev/null 2>&1 &

# Change /dev/null to /home/login.log
# After modification, logs will be saved to /home/login.log
./Login >/home/login.log 2>&1 &
```

## Getting the Program

There are two ways to obtain this program:

### 1. Use Pre-compiled Versions

[Click here](https://github.com/liuguangw/billing_go/releases) to view pre-compiled versions. Note that the latest compiled version may not always be available.

### 2. Manual Compilation

Requirements:
- Internet connection
- Git installed
- make installed
- Go 1.23 or higher installed

**Compilation using make:**

```bash
# Build
make

# Clean
make clean

# Build and package 32-bit
make x32

# Build and package 64-bit
make x64

# Build and package all architectures
make all
```

## File Description

```
billing       - Billing server executable
config.yaml   - Configuration file
```

## Configuration File

The configuration file must be in the same directory as the program. Supports two formats: `yaml` or `json`. File names should be `config.yaml` or `config.json`. If both files exist, `yaml` format takes priority.

### YAML Configuration Example

```yaml
# Lines starting with # are comments
# Strings don't need quotes unless they contain # characters
# If database password contains # or spaces, use quotes
#
# Billing server IP, default 127.0.0.1
ip: 127.0.0.1
#
# Billing server listening port (choose an unused port)
port: 12680
#
# MySQL server IP or hostname
db_host: localhost
#
# MySQL server port
db_port: 3306
#
# MySQL username
db_user: root
#
# MySQL password
db_password: 'root'
#
# Account database name (usually 'web')
db_name: web
#
# Only set to true when old MySQL versions report old_password error
allow_old_password: false
#
# Whether to guide users to register when login account doesn't exist
auto_reg: true
#
# Allowed server connection IPs. Empty means allow any IP
# When not empty, only specified IPs are allowed to connect
#allow_ips:
#  - 1.1.1.1
#  - 127.0.0.1
#
# Point correction. Configure only when displayed points are ±1 from actual
# If displayed points are 1 less, set to 1 (temporary solution, fix client script instead)
# If displayed points are 1 more, set to -1 (usually doesn't happen)
point_fix: 0
# Total player count limit, 0 means unlimited
max_client_count: 500
#
# Maximum users per IP address, 0 means unlimited
ip_max_client_count: 3
# Billing type: 0=classic, 1=retro
bill_type: 0
```

### JSON Configuration Example

```json
{
  "ip": "127.0.0.1",
  "port": 12680,
  "db_host": "localhost",
  "db_port": 3306,
  "db_user": "root",
  "db_password": "root",
  "db_name": "web",
  "allow_old_password": false,
  "auto_reg": true,
  "allow_ips": [],
  "point_fix": 0,
  "max_client_count": 500,
  "ip_max_client_count": 3,
  "bill_type": 0
}
```

> If billing and game server are on the same machine, use 127.0.0.1 for billing IP to avoid external network routing.
>
> The configuration file included with this project contains default values. If your configuration values match the defaults, you can omit those fields.

Place `billing` and the configuration file in the same directory.

Modify the game server configuration file `....../tlbb/Server/Config/ServerInfo.ini` billing section:

```ini
#........
[Billing]
Number=1
# Billing server IP
IP0=127.0.0.1
# Billing server listening port
Port0=12680
#.........
```

Finally, start the game server and billing service.

## Starting and Stopping

### Starting

**Foreground mode:**

```bash
# Navigate to billing directory, e.g., /home
cd /home
# Add execute permission
chmod +x ./billing
# Start billing
./billing
```

**Daemon mode:**

```bash
# Navigate to billing directory, e.g., /home
cd /home
# Add execute permission
chmod +x ./billing
# Start billing as daemon
./billing up -d
```

**Systemd service:**
See reference file [billing.service](billing.service)

### Stopping

Stop billing command:

```bash
# Use stop command
./billing stop

# Or use kill command
kill -SIGTERM $(pgrep -f billing)
```

For foreground mode, use Ctrl + C to stop the server.

## Development Commands

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

### Additional Commands
```bash
# View version info
./billing version

# Show online users
./billing show-users
```

## Testing

This project does not have unit tests. After making changes:
1. Build the project using `make`
2. Manually test by starting the billing service
3. Check `billing.log` for any errors