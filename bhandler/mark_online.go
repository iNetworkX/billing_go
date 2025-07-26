package bhandler

import (
	"github.com/liuguangw/billing_go/common"
	"time"
)

// markOnline 标记用户为在线状态
func markOnline(loginUsers, onlineUsers map[string]*common.ClientInfo, ipCounters map[string]int,
	activeConnections map[string]*common.ConnectionInfo, username string, clientInfo *common.ClientInfo) {
	//已经标记为登录了
	if _, userOnline := onlineUsers[username]; userOnline {
		// 用户已经在游戏中，只更新活跃时间，不增加计数
		// 更新活跃连接时间戳
		if conn, exists := activeConnections[username]; exists {
			conn.LastActivity = time.Now()
		}
		return
	}
	//从loginUsers中删除
	if loginInfo, userLogin := loginUsers[username]; userLogin {
		delete(loginUsers, username)
		//补充字段信息
		clientInfo.MacMd5 = loginInfo.MacMd5
		if clientInfo.IP == "" {
			clientInfo.IP = loginInfo.IP
		}
	}
	//写入onlineUsers
	onlineUsers[username] = clientInfo
	//增加IP计数（只有进入游戏才计数）
	if ipCounter, exists := ipCounters[clientInfo.IP]; exists {
		ipCounters[clientInfo.IP] = ipCounter + 1
	} else {
		ipCounters[clientInfo.IP] = 1
	}
	// 更新活跃连接
	if activeConnections != nil {
		activeConnections[username] = &common.ConnectionInfo{
			Username:     username,
			IP:           clientInfo.IP,
			LastActivity: time.Now(),
		}
	}
}

// updateUserActivity updates the LastActivity timestamp for a user in either LoginUsers or OnlineUsers
func updateUserActivity(loginUsers, onlineUsers map[string]*common.ClientInfo,
	activeConnections map[string]*common.ConnectionInfo, username string) {
	// Update activity for active connections
	if activeConnections != nil {
		if conn, exists := activeConnections[username]; exists {
			conn.LastActivity = time.Now()
		}
	}
}
