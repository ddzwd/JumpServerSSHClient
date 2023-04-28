/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"JumpServerSSHClient/cmd/conf"
	"JumpServerSSHClient/cmd/secret"
	"JumpServerSSHClient/cmd/ssh"
	"JumpServerSSHClient/cmd/user"
	"JumpServerSSHClient/instance"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "JumpServerSSHClient",
	Short: "JumpServer 简化连接工具",
	Long:  `支持自动登录，支持多用户，支持多服务器，友好的代码提示`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.JumpServerSSHClient.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.PersistentFlags().StringVarP(&instance.USRENAME, "username", "u", "", "指定用户")
	RootCmd.PersistentFlags().StringVarP(&instance.CONFIG_FILE, "conf", "c", instance.CONFIG_FILE, "配置文件路径")

	RootCmd.PersistentFlags().StringVarP(&instance.LogLevel, "level", "l", instance.LogLevel, fmt.Sprintf("日志等级: %s", "debug, info, warn, error, fatal, panic"))

	cobra.OnInitialize(func() {
		instance.InitLog()
	})

	user.RegistCMD(RootCmd)
	secret.RegistCMD(RootCmd)
	conf.RegistCMD(RootCmd)
	ssh.RegistCMD(RootCmd)

}
