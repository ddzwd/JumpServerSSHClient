/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package conf

import (
	"github.com/ddzwd/JumpServerSSHClient/config"
	"github.com/ddzwd/JumpServerSSHClient/instance"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化配置文件",
	Long:  `在指定路径创建一个空的配置文件,指定配置文件存在时,将忽略此次操作`,
	Run: func(cmd *cobra.Command, args []string) {

		instance.Logger.Infof("正在初始化配置文件: %s", instance.CONFIG_FILE)
		if config.ValidateConfigExist(instance.CONFIG_FILE) {
			is_force, _ := cmd.Flags().GetBool("force")
			if is_force {
				instance.Logger.Warnln("强制初始化配置文件")
			} else {
				instance.Logger.Fatalln("配置文件已存在,请勿重复初始化")
			}
		}
		config.InitConfig(instance.CONFIG_FILE)
	},
}

func init() {
	confCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	initCmd.Flags().BoolP("force", "f", false, "强制初始化配置文件")

	initCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		cmd.Flags().MarkHidden("username")
		cmd.Parent().HelpFunc()(cmd, args)
	})

}
