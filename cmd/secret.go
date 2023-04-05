/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// secretCmd represents the secret command
var secretCmd = &cobra.Command{
	Use:   "secret",
	Short: "秘钥管理",
	Long:  `用于生成秘钥,导入秘钥,导出秘钥`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	secret_key := args[0]
	// 	fmt.Println("code:", server.MFAHmacSha1(secret_key))
	// },
}

func init() {
	rootCmd.AddCommand(secretCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// secretCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// secretCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
