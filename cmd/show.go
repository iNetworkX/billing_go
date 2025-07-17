package cmd

import (
	"fmt"
	"github.com/liuguangw/billing_go/services/billing"
	"github.com/urfave/cli/v2"
)

// ShowCommand 显示IP连接信息的命令
func ShowCommand() *cli.Command {
	return &cli.Command{
		Name:   "show",
		Usage:  "show IP connections and associated accounts",
		Action: runShowCommand,
	}
}

// runShowCommand 显示IP连接信息
func runShowCommand(c *cli.Context) error {
	//初始化server
	server, err := billing.NewServer()
	if err != nil {
		return err
	}
	fmt.Println("show IP connections ...")
	return server.ShowConnections()
}