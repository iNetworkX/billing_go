# Show Command Implementation

## Overview
Added a new `show` command to display IP address information with connection and account counts.

## Implementation Details

### 1. Command Definition (`cmd/show.go`)
- Created new CLI command using urfave/cli framework
- Command name: `show`
- Usage: "show IP addresses with connection and account counts"
- Calls `server.ShowIPInfo()` method

### 2. Server Method (`services/billing/show_ip_info.go`)
- Sends packet with OpData "show_ip_info" to command handler
- Follows same pattern as ShowUsers command

### 3. Command Handler Logic (`bhandler/command_handler.go`)
- Added `ShowIPInfo(response *common.BillingPacket)` method
- Updated `GetResponse` to handle "show_ip_info" packet
- Implementation:
  - Collects IP addresses from LoginUsers and OnlineUsers maps
  - Counts connections per IP from IPCounters map
  - Creates formatted table with columns:
    - IP Address (20 chars)
    - Connections (15 chars)
    - Accounts (15 chars)
    - Account List (comma-separated)
  - Shows summary totals at bottom

### 4. Command Registration
- Added `ShowCommand()` to the Commands array in `cmd/app_command.go`

## Usage
```bash
./billing show
```

## Output Format
```
=== IP Address Information ===

IP Address           Connections     Accounts        Account List
--------------------------------------------------------------------------------
192.168.1.100        3               2               user1, user2
192.168.1.101        1               1               user3
--------------------------------------------------------------------------------
Total: 2 IPs, 4 Connections, 3 Unique Sessions
```