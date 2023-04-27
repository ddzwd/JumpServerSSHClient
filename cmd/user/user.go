/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package user

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

func RegistCMD(parent *cobra.Command) {
	parent.AddCommand(userCmd)

}

func init() {
	userCmd.PersistentFlags().StringVarP(&instance.PASSWORD, "password", "p", "", "指定用户密码")
	userCmd.PersistentFlags().StringVar(&instance.HOST, "host", "", "指定用户主机")
	userCmd.PersistentFlags().StringVar(&instance.PORT, "port", "22", "指定用户主机端口")
	userCmd.PersistentFlags().StringVarP(&instance.SECRET_KEY, "secret", "s", "", "MFA 密钥")
	userCmd.PersistentFlags().StringVarP(&instance.RSA_KEY_PATH, "rsa", "r", "", "RSA 密钥路径")
	userCmd.PersistentFlags().StringVar(&instance.ALIAS, "alias", "", "别名")
}
