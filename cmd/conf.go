/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

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

func init() {

	rootCmd.AddCommand(confCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// confCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// confCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// 将 rootCmd 的标志继承到 confCmd
	// confCmd.InheritedFlags = rootCmd.PersistentFlags()

	// 获取 username 标志
	usernameFlag := confCmd.Flags().Lookup("username")
	// 隐藏 username 标志
	if usernameFlag != nil {
		fmt.Println("隐藏 username 标志")
		usernameFlag.Hidden = true
	}

}
