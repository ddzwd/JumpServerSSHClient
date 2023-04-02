/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/ddzwd/JumpServerSSHClient/instance"
	"github.com/spf13/cobra"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "用户相关操作",
	Long:  `支持添加,删除,修改用户`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("user called", USRENAME)
	// },
}

func init() {
	rootCmd.AddCommand(userCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	userCmd.PersistentFlags().StringVarP(&instance.PASSWORD, "password", "p", "", "指定用户密码")
	userCmd.PersistentFlags().StringVar(&instance.HOST, "host", "", "指定用户主机")
	userCmd.PersistentFlags().StringVar(&instance.PORT, "port", "", "指定用户主机端口")
	userCmd.PersistentFlags().StringVarP(&instance.SECRET_KEY, "secret", "s", "", "MFA 密钥")
	userCmd.PersistentFlags().StringVarP(&instance.RSA_KEY_PATH, "ras", "r", "", "RSA 密钥路径")
	userCmd.PersistentFlags().StringVar(&instance.ALIAS, "alas", "", "别名")
}
