/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package secret

import (
	"fmt"

	"github.com/ddzwd/JumpServerSSHClient/server"
	"github.com/spf13/cobra"
)

// secretCmd represents the secret command
var secretCmd = &cobra.Command{
	Use:   "secret",
	Short: "秘钥管理",
	Long:  `用于生成秘钥,导入秘钥,导出秘钥`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, v := range args {
			fmt.Println(server.NewMFA(v))
		}

	},
	Aliases: []string{"st"},
}

func RegistCMD(parent *cobra.Command) {
	parent.AddCommand(secretCmd)
}

func init() {

}
