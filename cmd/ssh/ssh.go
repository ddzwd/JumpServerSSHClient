/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package ssh

import (
	"github.com/ddzwd/JumpServerSSHClient/config"
	"github.com/ddzwd/JumpServerSSHClient/instance"
	"github.com/ddzwd/JumpServerSSHClient/terminal"
	"github.com/spf13/cobra"
)

// sshCmd represents the ssh command
var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "ssh 连接",
	Run: func(cmd *cobra.Command, args []string) {
		var unique_id string
		if len(args) > 0 {
			unique_id = args[0]
		}
		conf := config.LoadConfig(instance.CONFIG_FILE)
		if unique_id == "" {
			unique_id = conf.DefaultUser
		}
		if unique_id == "" {
			instance.Logger.Error("cant find default user")
		}

		user := config.GetUserById(unique_id, &conf)
		if user == nil {
			instance.Logger.Errorf("cant find user by id %s", unique_id)
		}
		terminal.StartTerminal(user)
	},
}

func RegistCMD(parent *cobra.Command) {
	parent.AddCommand(sshCmd)
}

func init() {

}
