/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/ddzwd/JumpServerSSHClient/cmd"
	"github.com/ddzwd/JumpServerSSHClient/instance"
	"github.com/spf13/cobra"
)

func main() {
	cobra.OnInitialize(func() {
		// 初始化日志配置
		instance.InitLog()
	})
	cmd.Execute()
}
