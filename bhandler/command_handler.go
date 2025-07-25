package bhandler

import (
	"bytes"
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"

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
	} else if bytes.Equal(request.OpData, []byte("show_ip_info")) {
		h.ShowIPInfo(response)
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

// ShowIPInfo 显示IP地址信息，包括每个IP的连接数和账号数
func (h *CommandHandler) ShowIPInfo(response *common.BillingPacket) {
	var content strings.Builder
	content.WriteString("\n=== IP Address Information ===\n\n")

	// 添加调试信息
	content.WriteString(fmt.Sprintf("Debug: LoginUsers count: %d\n", len(h.Resource.LoginUsers)))
	content.WriteString(fmt.Sprintf("Debug: OnlineUsers count: %d\n", len(h.Resource.OnlineUsers)))
	content.WriteString(fmt.Sprintf("Debug: IPCounters count: %d\n", len(h.Resource.IPCounters)))
	content.WriteString("\n")

	// 创建一个map来存储每个IP的账号列表
	ipAccounts := make(map[string]map[string]bool)

	// 从LoginUsers收集IP和账号信息
	for username, clientInfo := range h.Resource.LoginUsers {
		if clientInfo == nil {
			continue
		}
		if _, exists := ipAccounts[clientInfo.IP]; !exists {
			ipAccounts[clientInfo.IP] = make(map[string]bool)
		}
		ipAccounts[clientInfo.IP][username] = true
	}

	// 从OnlineUsers收集IP和账号信息
	for username, clientInfo := range h.Resource.OnlineUsers {
		if clientInfo == nil {
			continue
		}
		if _, exists := ipAccounts[clientInfo.IP]; !exists {
			ipAccounts[clientInfo.IP] = make(map[string]bool)
		}
		ipAccounts[clientInfo.IP][username] = true
	}

	// 获取所有IP地址并排序
	var ips []string
	for ip := range ipAccounts {
		ips = append(ips, ip)
	}
	// 也包括只在IPCounters中的IP（可能没有活跃账号）
	for ip := range h.Resource.IPCounters {
		if _, exists := ipAccounts[ip]; !exists {
			ips = append(ips, ip)
			ipAccounts[ip] = make(map[string]bool)
		}
	}
	sort.Strings(ips)

	// 打印表头
	content.WriteString(fmt.Sprintf("%-20s %-15s %-15s %s\n", "IP Address", "Game Connections", "Accounts", "Account List"))
	content.WriteString(strings.Repeat("-", 80) + "\n")

	// 打印每个IP的信息
	for _, ip := range ips {
		connections := h.Resource.IPCounters[ip]
		accounts := ipAccounts[ip]
		// 账号数等于连接数（每个连接算作一个账号会话）
		accountCount := connections
		if accountCount == 0 {
			accountCount = len(accounts)
		}

		// 获取账号列表
		var accountList []string
		for account := range accounts {
			accountList = append(accountList, account)
		}
		sort.Strings(accountList)
		accountListStr := strings.Join(accountList, ", ")
		if accountListStr == "" {
			accountListStr = "-"
		}

		content.WriteString(fmt.Sprintf("%-20s %-15d %-15d %s\n", ip, connections, accountCount, accountListStr))
	}

	// 打印总计
	content.WriteString(strings.Repeat("-", 80) + "\n")
	totalIPs := len(ips)
	totalConnections := 0
	for _, count := range h.Resource.IPCounters {
		totalConnections += count
	}
	totalAccounts := len(h.Resource.LoginUsers) + len(h.Resource.OnlineUsers)
	content.WriteString(fmt.Sprintf("Total: %d IPs, %d Game Connections, %d Unique Sessions\n", totalIPs, totalConnections, totalAccounts))

	response.OpData = []byte(content.String())
}
