/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package user

import (
	"fmt"

	"JumpServerSSHClient/config"
	"JumpServerSSHClient/instance"

	"github.com/spf13/cobra"
)

// userDeleteCmd represents the userDelete command
var userDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "删除用户数据",
	// Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		conf := config.LoadConfig(instance.CONFIG_FILE)
		fmt.Println(args)
		for _, unique_id := range args {
			instance.Logger.Infof("ready to delete user %s", unique_id)
			user := config.GetUserById(unique_id, &conf)
			if user == nil {
				instance.Logger.Errorf("cont find record by id %s", unique_id)
			} else {
				user.Delete(&conf)
				instance.Logger.Infof("delete user %s success", unique_id)
			}
		}
		conf.Save()
	},
}

func init() {
	userCmd.AddCommand(userDeleteCmd)
	userDeleteCmd.SetHelpFunc(func(cmd *cobra.Command, strings []string) {
		// Hide flag for this command
		if err := cmd.Flags().MarkHidden("username"); err != nil {
			instance.Logger.Errorln(err)
		}
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
