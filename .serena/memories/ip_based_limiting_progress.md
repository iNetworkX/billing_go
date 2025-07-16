# IP-based Player Limiting Implementation Progress

## Overview
The billing_go project has been migrated from MAC-based to IP-based player limiting. This change improves security and simplifies client connections.

## Completed Tasks

### Task 1: Update mark_online.go ✅
- Changed from MAC counting to IP counting when players enter the game
- Updated `markOnline` function signature to accept `ipCounters` and `activeConnections`
- Updated all 7 handler files that call `markOnline`

### Task 2: Update logout_handler.go ✅ 
- **Enhanced IP counter cleanup**: Deletes IP counter entries when they reach 0 to prevent memory leaks
- **Extended logout handling**: Now handles both `OnlineUsers` (in-game) and `LoginUsers` (logged in but not in game)
- **Complete cleanup flow**:
  - Removes user from appropriate state map
  - Decrements IP counter for the user's IP  
  - Removes IP counter entry if it reaches 0
  - Removes user from ActiveConnections map
- **Code location**: `bhandler/logout_handler.go` lines 28-76

### Task 3: Implement Connection Health Check ✅
- Created periodic health check goroutine in `services/billing/health_check.go`
- Automatically cleans up stale connections based on configurable timeout
- Default timeout: 300 seconds (5 minutes)
- Health check runs every timeout/2 seconds
- Can be disabled by setting `connection_timeout: 0`

### Task 4: Update command_handler.go ✅
- **Enhanced show-users command**: Now displays four sections:
  1. Login users (authenticated but not in game)
  2. Online users (currently in game) 
  3. IP counters (connections per IP)
  4. Active connections (all connections with last activity time)
- **Active connection display**: Shows username, IP, and last activity timestamp
- **Code location**: `bhandler/command_handler.go` lines 75-86

### Task 5: Initialize New Maps ✅
- Initialized `IPCounters` and `ActiveConnections` maps in `services/billing/load_handlers.go`
- Removed old `MacCounters` initialization

### Task 6: Update All References ✅
- Changed all `PcMaxClientCount` references to `IPMaxClientCount`
- Changed all `MacCounters` references to `IPCounters`
- Updated configuration loading in `services/load_server_config.go`

## Key Data Structures

### HandlerResource (common/handler_resource.go)
```go
type HandlerResource struct {
    Db                *sql.DB
    Logger            *zap.Logger
    LoginUsers        map[string]*ClientInfo     // 已登录,还未进入游戏的用户
    OnlineUsers       map[string]*ClientInfo     // 已进入游戏的用户
    IPCounters        map[string]int             // IP地址计数器
    ActiveConnections map[string]*ConnectionInfo // 活跃连接映射
}
```

### ConnectionInfo (common/handler_resource.go)
```go
type ConnectionInfo struct {
    Username     string
    IP           string
    LastActivity time.Time
}
```

## Configuration Change
Users must update their config files:
- Old: `pc_max_client_count: 3`
- New: `ip_max_client_count: 3`
- Optional: `connection_timeout: 300` (for health check, in seconds)

## All Tasks Completed
The IP-based player limiting implementation is now fully complete with automatic health monitoring and comprehensive connection tracking.