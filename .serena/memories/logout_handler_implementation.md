# Logout Handler Implementation Details

## Overview
The logout handler (`bhandler/logout_handler.go`) manages user logout and cleanup of associated resources.

## Implementation Details

### Handler Structure
- Type: `LogoutHandler` 
- Packet Type: `packetTypeLogout` (0xA2)
- Main method: `GetResponse(request *common.BillingPacket) *common.BillingPacket`

### Logout Flow

1. **Parse logout request**
   - Read username length (1 byte)
   - Read username (variable length)

2. **Check OnlineUsers first** (lines 29-51)
   - If user is in game:
     - Remove from OnlineUsers map
     - Decrement IP counter
     - Delete IP counter entry if it reaches 0
     - Remove from ActiveConnections

3. **Check LoginUsers if not in game** (lines 52-76)
   - If user logged in but not in game:
     - Remove from LoginUsers map
     - Decrement IP counter
     - Delete IP counter entry if it reaches 0
     - Remove from ActiveConnections

4. **Log the logout event**
   - Log message: "user [username] logout game"

5. **Send response**
   - Response format: [username_length][username][0x1]

### Key Improvements (Task 2)

1. **Memory leak prevention**: IP counters are deleted when they reach 0
2. **Dual state handling**: Supports users in both LoginUsers and OnlineUsers states
3. **Consistent cleanup**: Same cleanup logic for both user states
4. **Bounds checking**: Ensures IP counter never goes negative

### Edge Cases Handled

1. **User not found**: No error, just logs logout (graceful handling)
2. **IP counter underflow**: Prevented with bounds check (lines 37-39, 62-64)
3. **Nil ActiveConnections**: Checked before deletion (lines 48, 73)
4. **Empty IP string**: Checked before processing (lines 31, 56)

### Remaining Limitation
Abrupt disconnections (without logout packet) are not handled by this handler. Task 3's connection health check will address this.