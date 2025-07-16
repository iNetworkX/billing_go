# TODO: IP-based Player Limiting Implementation

## Task 1: Update mark_online.go ✅ COMPLETED
- [x] Change from MAC counting to IP counting when players enter the game
- [x] Update the `markOnline` function signature to accept `ipCounters` instead of `macCounters`
- [x] Increment IP counter when player enters the game
- [x] Update active connection timestamp

**Files modified:**
- `bhandler/mark_online.go` - Updated function signature and logic
- `bhandler/enter_game_handler.go` - Updated markOnline call
- `bhandler/cost_log_handler.go` - Updated markOnline call
- `bhandler/keep_handler.go` - Updated markOnline call
- `bhandler/prize_card_handler.go` - Updated markOnline call
- `bhandler/prize_fetch_handler.go` - Updated markOnline call
- `bhandler/prize_handler.go` - Updated markOnline call
- `bhandler/query_point_handler.go` - Updated markOnline call

---

## Task 2: Update logout_handler.go ✅ COMPLETED
- [x] Decrement IP counters when player logs out
- [x] Remove user from `ActiveConnections` map
- [x] Clean up IP counter if it reaches 0
- [x] Handle both explicit logout and connection cleanup scenarios
- [x] Handle logout for users in both `OnlineUsers` and `LoginUsers` states

**Files modified:**
- `bhandler/logout_handler.go` - Complete implementation with all edge cases handled

---

## Task 2 Summary:
**What was completed:**
- Enhanced IP counter cleanup to prevent memory leaks by removing entries when counter reaches 0
- Added support for handling logout of users in both `OnlineUsers` and `LoginUsers` states
- Ensured proper cleanup of `ActiveConnections` for all user states
- Added bounds checking and proper error handling

**Key changes in logout_handler.go:**
- Lines 41-45: Delete IP counter from map when it reaches 0 (for OnlineUsers)
- Lines 52-75: Added new logic to handle users who logged in but haven't entered game yet
- Lines 66-70: Delete IP counter from map when it reaches 0 (for LoginUsers)

**Note:** Abrupt disconnections (without logout packet) still require Task 3 implementation for automatic cleanup.

---

## Task 3: Implement Connection Health Check
- [ ] Create a new goroutine in server initialization for periodic health checks
- [ ] Check `ActiveConnections` for stale connections (e.g., no activity for 5 minutes)
- [ ] For stale connections:
  - Remove from `LoginUsers` and `OnlineUsers`
  - Decrement IP counter
  - Remove from `ActiveConnections`
  - Log the cleanup action
- [ ] Make timeout configurable

**Files to create/modify:**
- `services/billing/connection_monitor.go` (new file)
- `services/billing/server.go` (add goroutine)
- `common/server_config.go` (add timeout config)

---

## Task 4: Update command_handler.go ⚠️ PARTIALLY COMPLETED
- [x] Change display from MAC counters to IP counters
- [x] Update the "show-users" command output format
- [ ] Display active connections with their last activity time

**Files modified:**
- `bhandler/command_handler.go` - Basic IP counter display done, active connections display pending

---

## Task 5: Initialize New Maps ✅ COMPLETED
- [x] Find where `HandlerResource` is initialized
- [x] Add initialization for `IPCounters` map
- [x] Add initialization for `ActiveConnections` map
- [x] Remove old `MacCounters` initialization

**Files modified:**
- `services/billing/load_handlers.go` - Updated HandlerResource initialization
- `services/load_server_config.go` - Changed PcMaxClientCount to IPMaxClientCount in defaults

---

## Task 6: Update All References ✅ COMPLETED
- [x] Search for all references to `PcMaxClientCount` and update to `IPMaxClientCount`
- [x] Search for all references to `MacCounters` and update to `IPCounters`
- [x] Update any error messages that mention "pc_max_client_count"
- [x] Update configuration loading code

**Files checked and updated:**
- All handler files in `bhandler/` directory that used MacCounters
- `services/load_server_config.go` - Updated default configuration
- `services/billing/load_handlers.go` - Updated LoginHandler initialization

---

## Task 7: Testing and Documentation
- [ ] Test IP limiting with multiple connections from same IP
- [ ] Test connection timeout detection
- [ ] Test proper cleanup on logout
- [ ] Update README or documentation with new configuration
- [ ] Create example configuration with new field

**Testing scenarios:**
1. Connect multiple clients from same IP, verify limit works
2. Disconnect client abruptly, verify timeout cleanup
3. Normal logout, verify counter decrements
4. Mix of login/logout operations, verify counters stay accurate

---

## Configuration Migration Note
Users need to update their configuration files:
```yaml
# Old configuration
pc_max_client_count: 3

# New configuration
ip_max_client_count: 3
connection_timeout: 300  # seconds (optional, for health check)
```