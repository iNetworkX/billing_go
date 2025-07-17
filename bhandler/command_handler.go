package bhandler

import (
	"bytes"
	"context"
	"strconv"

	"github.com/liuguangw/billing_go/common"
)

// CommandHandler 处理发送过来的命令
type CommandHandler struct {
	Resource *common.HandlerResource
	Cancel   context.CancelFunc //关闭服务器的回调函数
}

// GetType 可以处理的消息类型
func (*CommandHandler) GetType() byte {
	return packetTypeCommand
}

// GetResponse 根据请求获得响应
func (h *CommandHandler) GetResponse(request *common.BillingPacket) *common.BillingPacket {
	response := request.PrepareResponse()
	response.OpData = []byte{0, 0}
	//./billing show_users
	//获取billing中用户列表状态
	if bytes.Equal(request.OpData, []byte("show_users")) {
		h.showUsers(response)
	} else if bytes.Equal(request.OpData, []byte("show")) {
		//./billing show
		//显示IP连接信息
		h.showConnections(response)
	} else {
		//./billing stop
		//关闭billing服务
		//执行cancel后, Server.Run()中的ctx.Done()会达成,主协程会退出
		h.Cancel()
	}
	return response
}

// showUsers 将用户列表状态写入response
func (h *CommandHandler) showUsers(response *common.BillingPacket) {
	content := "login users:"
	if len(h.Resource.LoginUsers) == 0 {
		content += " empty"
	} else {
		for username, clientInfo := range h.Resource.LoginUsers {
			content += "\n\t" + username + ": " + clientInfo.String()
		}
	}
	//
	content += "\n\nonline users:"
	if len(h.Resource.OnlineUsers) == 0 {
		content += " empty"
	} else {
		for username, clientInfo := range h.Resource.OnlineUsers {
			content += "\n\t" + username + ": " + clientInfo.String()
		}
	}
	//
	content += "\n\nIP counters:"
	if len(h.Resource.IPCounters) == 0 {
		content += " empty"
	} else {
		for ip, counterValue := range h.Resource.IPCounters {
			content += "\n\t" + ip + ": " + strconv.Itoa(counterValue)
		}
	}
	//
	content += "\n\nactive connections:"
	if len(h.Resource.ActiveConnections) == 0 {
		content += " empty"
	} else {
		for username, connInfo := range h.Resource.ActiveConnections {
			content += "\n\t" + username + " (" + connInfo.IP + "): last activity " + connInfo.LastActivity.Format("2006-01-02 15:04:05")
		}
	}
	response.OpData = []byte(content)
}

// showConnections 显示IP连接信息和关联账户
func (h *CommandHandler) showConnections(response *common.BillingPacket) {
	content := "IP Connections Summary:\n"
	
	// 创建一个映射来跟踪每个IP的账户
	ipAccounts := make(map[string][]string)
	
	// 从LoginUsers收集IP和账户信息
	for username, clientInfo := range h.Resource.LoginUsers {
		ipAccounts[clientInfo.IP] = append(ipAccounts[clientInfo.IP], username+" (login)")
	}
	
	// 从OnlineUsers收集IP和账户信息
	for username, clientInfo := range h.Resource.OnlineUsers {
		ipAccounts[clientInfo.IP] = append(ipAccounts[clientInfo.IP], username+" (online)")
	}
	
	// 显示IP连接信息
	if len(h.Resource.IPCounters) == 0 {
		content += "No active connections"
	} else {
		for ip, count := range h.Resource.IPCounters {
			content += "\nIP: " + ip + " | Connections: " + strconv.Itoa(count)
			if accounts, exists := ipAccounts[ip]; exists {
				content += " | Accounts: "
				for i, account := range accounts {
					if i > 0 {
						content += ", "
					}
					content += account
				}
			}
		}
	}
	
	response.OpData = []byte(content)
}
