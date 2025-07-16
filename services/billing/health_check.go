package billing

import (
	"time"

	"github.com/liuguangw/billing_go/common"
	"go.uber.org/zap"
)

// runHealthCheck 运行定期健康检查，清理断开的连接
func (s *Server) runHealthCheck(resource *common.HandlerResource) {
	// 如果ConnectionTimeout为0，表示不启用健康检查
	if s.config.ConnectionTimeout <= 0 {
		s.logger.Info("Connection health check disabled (connection_timeout=0)")
		return
	}

	ticker := time.NewTicker(time.Duration(s.config.ConnectionTimeout) * time.Second / 2)
	defer ticker.Stop()

	s.logger.Info("Connection health check started",
		zap.Int("timeout_seconds", s.config.ConnectionTimeout),
		zap.Int("check_interval_seconds", s.config.ConnectionTimeout/2))

	for {
		select {
		case <-ticker.C:
			// 执行健康检查
			s.performHealthCheck(resource)
		case <-time.After(time.Hour * 24 * 365): // 实际上永远不会触发，用于防止goroutine泄漏
			return
		}

		// 检查服务器是否还在运行
		if !s.running {
			s.logger.Info("Health check stopped: server is no longer running")
			return
		}
	}
}

// performHealthCheck 执行实际的健康检查
func (s *Server) performHealthCheck(resource *common.HandlerResource) {
	now := time.Now()
	timeout := time.Duration(s.config.ConnectionTimeout) * time.Second
	var staleConnections []string

	// 检查所有活跃连接
	for username, connInfo := range resource.ActiveConnections {
		if now.Sub(connInfo.LastActivity) > timeout {
			staleConnections = append(staleConnections, username)
		}
	}

	// 清理过期的连接
	for _, username := range staleConnections {
		connInfo, exists := resource.ActiveConnections[username]
		if !exists {
			continue
		}

		// 如果IP为空，使用占位符避免空字符串日志
		logIP := connInfo.IP
		if logIP == "" {
			logIP = "unknown"
		}
		
		s.logger.Warn("Cleaning stale connection",
			zap.String("username", username),
			zap.String("ip", logIP),
			zap.Duration("inactive_duration", now.Sub(connInfo.LastActivity)))

		// 从在线用户中移除
		if _, ok := resource.OnlineUsers[username]; ok {
			delete(resource.OnlineUsers, username)
			s.logger.Debug("Removed from OnlineUsers", zap.String("username", username))
		}

		// 从登录用户中移除
		if _, ok := resource.LoginUsers[username]; ok {
			delete(resource.LoginUsers, username)
			s.logger.Debug("Removed from LoginUsers", zap.String("username", username))
		}

		// 减少IP计数器
		if resource.IPCounters[connInfo.IP] > 0 {
			resource.IPCounters[connInfo.IP]--
			s.logger.Debug("Decremented IP counter",
				zap.String("ip", connInfo.IP),
				zap.Int("new_count", resource.IPCounters[connInfo.IP]))

			// 如果计数器为0，删除该条目以防止内存泄漏
			if resource.IPCounters[connInfo.IP] == 0 {
				delete(resource.IPCounters, connInfo.IP)
				s.logger.Debug("Removed IP counter entry", zap.String("ip", connInfo.IP))
			}
		}

		// 从活跃连接中移除
		delete(resource.ActiveConnections, username)
		s.logger.Info("Cleaned stale connection",
			zap.String("username", username),
			zap.String("ip", logIP))
	}

	if len(staleConnections) > 0 {
		s.logger.Info("Health check completed",
			zap.Int("stale_connections_cleaned", len(staleConnections)),
			zap.Int("active_connections", len(resource.ActiveConnections)),
			zap.Int("online_users", len(resource.OnlineUsers)),
			zap.Int("login_users", len(resource.LoginUsers)))
	}
}
