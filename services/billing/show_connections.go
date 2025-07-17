package billing

import (
	"fmt"
	"github.com/liuguangw/billing_go/common"
)

// ShowConnections 通知服务器显示IP连接信息
func (s *Server) ShowConnections() error {
	packet := &common.BillingPacket{
		MsgID:  [2]byte{0, 0},
		OpData: []byte("show"),
	}
	response, err := s.sendPacketToServer(packet)
	if err != nil {
		return err
	}
	fmt.Println(string(response.OpData))
	return nil
}