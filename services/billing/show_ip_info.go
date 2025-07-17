package billing

import (
	"fmt"

	"github.com/liuguangw/billing_go/common"
)

// ShowIPInfo 显示IP地址信息，包括每个IP的连接数和账号数
func (s *Server) ShowIPInfo() error {
	packet := &common.BillingPacket{
		MsgID:  [2]byte{0, 0},
		OpData: []byte("show_ip_info"),
	}
	response, err := s.sendPacketToServer(packet)
	if err != nil {
		return err
	}
	fmt.Println(string(response.OpData))
	return nil
}