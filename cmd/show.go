package cmd

import (
	"fmt"

	"github.com/liuguangw/billing_go/services/billing"
	"github.com/urfave/cli/v2"
)

// ShowCommand 返回 show 命令
func ShowCommand() *cli.Command {
	return &cli.Command{
		Name:   "show",
		Usage:  "show IP addresses with connection and account counts",
		Action: runShowCommand,
	}
}

// runShowCommand 执行 show 命令
func runShowCommand(c *cli.Context) error {
	// 初始化server
	server, err := billing.NewServer()
	if err != nil {
		return err
	}
	fmt.Println("Show IP addresses information...")
	return server.ShowIPInfo()
}
