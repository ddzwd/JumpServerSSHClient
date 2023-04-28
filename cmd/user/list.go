/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package user

import (
	"JumpServerSSHClient/config"
	"JumpServerSSHClient/instance"

	"github.com/spf13/cobra"
)

// userListCmd represents the userList command
var userListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "列出所有用户",
	Run: func(cmd *cobra.Command, args []string) {
		var content string = ""
		if len(args) > 0 {
			content = args[0]
		}
		c := config.LoadConfig(instance.CONFIG_FILE)
		users := config.GrepUser(content, &c)
		config.PrettyPrint(users)
	},
}

func init() {
	userCmd.AddCommand(userListCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// userListCmd.PersistentFlags().MarkHidden("password")

	// 隐藏不需要的flag
	userListCmd.SetHelpFunc(func(cmd *cobra.Command, strings []string) {
		// Hide flag for this command
		if err := cmd.Flags().MarkHidden("password"); err != nil {
			instance.Logger.Errorln(err)
		}
		if err := cmd.Flags().MarkHidden("host"); err != nil {
			instance.Logger.Errorln(err)
		}
		if err := cmd.Flags().MarkHidden("rsa"); err != nil {
			instance.Logger.Errorln(err)
		}
		if err := cmd.Flags().MarkHidden("secret"); err != nil {
			instance.Logger.Errorln(err)
		}
		if err := cmd.Flags().MarkHidden("alias"); err != nil {
			instance.Logger.Errorln(err)
		}
		if err := cmd.Flags().MarkHidden("port"); err != nil {
			instance.Logger.Errorln(err)
		}
		// Call parent help func
		cmd.Parent().HelpFunc()(cmd, strings)
	})
}
