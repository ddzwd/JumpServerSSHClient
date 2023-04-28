/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package user

import (
	"JumpServerSSHClient/config"
	"JumpServerSSHClient/instance"

	"github.com/spf13/cobra"
)

// userModifyCmd represents the userModify command
var userModifyCmd = &cobra.Command{
	Use:     "modify",
	Aliases: []string{"md"},
	Short:   "修改用户信息",
	Args:    cobra.RangeArgs(1, 1),
	Run: func(cmd *cobra.Command, args []string) {
		unique_id := args[0]
		// 获取指定用户
		conf := config.LoadConfig(instance.CONFIG_FILE)
		user := config.GetUserById(unique_id, &conf)
		if user == nil {
			instance.Logger.Error("用户不存在")
		}
		// 修改用户信息
		if cmd.Flags().Lookup("username").Changed {
			instance.Logger.Debugf("change username %s -> %s", user.UserName, instance.USRENAME)
			user.UserName = instance.USRENAME
			instance.Logger.Debugf("change username success")

		}
		if cmd.Flags().Lookup("password").Changed {
			instance.Logger.Debugf("change password %s -> %s", user.Password, instance.PASSWORD)
			user.Password = instance.PASSWORD
			instance.Logger.Debug("change password success")
		}
		if cmd.Flags().Lookup("host").Changed {
			instance.Logger.Debugf("change host %s -> %s", user.Host, instance.HOST)
			user.Host = instance.HOST
			instance.Logger.Debug("change host success")
		}
		if cmd.Flags().Lookup("port").Changed {
			instance.Logger.Debugf("change port %s -> %s", user.Port, instance.PORT)
			user.Port = instance.PORT
			instance.Logger.Debug("change port success")
		}
		if cmd.Flags().Lookup("secret").Changed {
			instance.Logger.Debugf("change secret %s -> %s", user.SecretKey, instance.SECRET_KEY)
			user.SecretKey = instance.SECRET_KEY
			instance.Logger.Debug("change secret success")
		}
		if cmd.Flags().Lookup("rsa").Changed {
			instance.Logger.Debugf("change rsa %s -> %s", user.RsaKeyPath, instance.RSA_KEY_PATH)
			user.RsaKeyPath = instance.RSA_KEY_PATH
			instance.Logger.Debug("change rsa success")
		}
		if cmd.Flags().Lookup("alias").Changed {
			instance.Logger.Debugf("change alias %s -> %s", user.Alias, instance.ALIAS)
			user.Alias = instance.ALIAS
			instance.Logger.Debug("change alias success")
		}
		if cmd.Flags().Lookup("default").Changed {
			instance.Logger.Debugf("change default user %s -> %s", conf.DefaultUser, unique_id)
			conf.DefaultUser = unique_id
			instance.Logger.Debug("change default success")
		}
		user.Update(&conf)
		conf.Save()
	},
}

func init() {
	userCmd.AddCommand(userModifyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userModifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userModifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	userModifyCmd.Flags().BoolP("default", "d", false, "设置为默认用户")
}
