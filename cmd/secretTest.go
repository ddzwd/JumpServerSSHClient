/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/ddzwd/JumpServerSSHClient/server"
	"github.com/spf13/cobra"
)

// secretTestCmd represents the secretTest command
var secretTestCmd = &cobra.Command{
	Use:   "secretTest",
	Short: "测试MFA秘钥生成算法",
	Run: func(cmd *cobra.Command, args []string) {
		secret_key := args[0]
		fmt.Println("MFA code:", server.NewMFA(secret_key))
	},
	Aliases: []string{"st"},
}

func init() {
	secretCmd.AddCommand(secretTestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// secretTestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// secretTestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
