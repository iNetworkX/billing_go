package bhandler

import (
	"github.com/liuguangw/billing_go/common"
	"github.com/liuguangw/billing_go/services"
)

// LogoutHandler 退出游戏
type LogoutHandler struct {
	Resource *common.HandlerResource
}

// GetType 可以处理的消息类型
func (*LogoutHandler) GetType() byte {
	return packetTypeLogout
}

// GetResponse 根据请求获得响应
func (h *LogoutHandler) GetResponse(request *common.BillingPacket) *common.BillingPacket {
	response := request.PrepareResponse()
	packetReader := services.NewPacketDataReader(request.OpData)
	//用户名
	usernameLength := packetReader.ReadByteValue()
	tmpLength := int(usernameLength)
	username := packetReader.ReadBytes(tmpLength)
	//更新在线状态
	usernameStr := string(username)
	if clientInfo, userOnline := h.Resource.OnlineUsers[usernameStr]; userOnline {
		delete(h.Resource.OnlineUsers, usernameStr)
		ip := clientInfo.IP
		if ip != "" {
			ipCounter := 0
			if value, valueExists := h.Resource.IPCounters[ip]; valueExists {
				ipCounter = value
			}
			ipCounter--
			if ipCounter < 0 {
				ipCounter = 0
			}
			h.Resource.IPCounters[ip] = ipCounter
		}
		// 从活跃连接中删除
		if h.Resource.ActiveConnections != nil {
			delete(h.Resource.ActiveConnections, usernameStr)
		}
	}
	//
	h.Resource.Logger.Info("user [" + string(username) + "] logout game")
	//Packets::BLRetBillingEnd
	opData := make([]byte, 0, usernameLength+2)
	opData = append(opData, usernameLength)
	opData = append(opData, username...)
	opData = append(opData, 0x1)
	response.OpData = opData
	return response
}
