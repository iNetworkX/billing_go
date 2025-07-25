package common

import (
	"database/sql"
	"go.uber.org/zap"
	"time"
)

// HandlerResource handler所需的资源
type HandlerResource struct {
	Db                *sql.DB                    //数据库连接
	Logger            *zap.Logger                //日志对象
	LoginUsers        map[string]*ClientInfo     //已登录,还未进入游戏的用户
	OnlineUsers       map[string]*ClientInfo     //已进入游戏的用户
	IPCounters        map[string]int             //已连接用户的IP地址计数器（包括登录和游戏中状态）
	ActiveConnections map[string]*ConnectionInfo //活跃连接映射，key为用户名
	Config            *ServerConfig              //服务器配置
}

// ConnectionInfo 连接信息，用于跟踪活跃连接
type ConnectionInfo struct {
	Username     string    //用户名
	IP           string    //IP地址
	LastActivity time.Time //最后活动时间
}
