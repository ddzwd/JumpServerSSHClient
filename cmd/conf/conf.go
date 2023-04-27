/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package conf

import (
	"fmt"

	"github.com/spf13/cobra"
)

// confCmd represents the conf command
var confCmd = &cobra.Command{
	Use:   "conf",
	Short: "配置文件操作",
	Long:  `支持初始化配置文件,合法性校验`,
}

func RegistCMD(parent *cobra.Command) {
	parent.AddCommand(confCmd)
}

func init() {
	usernameFlag := confCmd.Flags().Lookup("username")
	// 隐藏 username 标志
	if usernameFlag != nil {
		fmt.Println("隐藏 username 标志")
		usernameFlag.Hidden = true
	}
}
