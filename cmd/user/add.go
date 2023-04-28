/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package user

import (
	"errors"
	"fmt"

	"JumpServerSSHClient/config"
	"JumpServerSSHClient/instance"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/spf13/cobra"
)

// userAddCmd represents the userAdd command
var userAddCmd = &cobra.Command{
	Use:   "add",
	Short: "添加用户",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().Lookup("input").Value.String() == "false" && instance.USRENAME == "" {
			return errors.New("please add -u from username")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id, err := gonanoid.Generate("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 16)
		if err != nil {
			instance.Logger.Errorf("生成用户唯一ID失败: %v", err)

		}
		if cmd.Flags().Lookup("input").Value.String() == "true" {
			instance.Logger.Debug("使用交互式输入")
			input_map := map[string]*string{
				"用户名":     &instance.USRENAME,
				"密码":      &instance.PASSWORD,
				"主机":      &instance.HOST,
				"端口":      &instance.PORT,
				"RSAPATH": &instance.RSA_KEY_PATH,
				"密钥":      &instance.SECRET_KEY,
				"别名":      &instance.ALIAS,
			}
			sorted_title := []string{"用户名", "密码", "主机", "端口", "RSAPATH", "密钥", "别名"}
			for _, k := range sorted_title {
				v := input_map[k]
				fmt.Printf("请输入%s(default %s):", k, *v)
				fmt.Scanln(v)
			}
		}

		user := config.User{
			UniqueId:   id,
			UserName:   instance.USRENAME,
			Password:   instance.PASSWORD,
			Host:       instance.HOST,
			Port:       instance.PORT,
			SecretKey:  instance.SECRET_KEY,
			RsaKeyPath: instance.RSA_KEY_PATH,
			Alias:      instance.ALIAS,
		}
		c := config.LoadConfig(instance.CONFIG_FILE)
		user.AddUser(&c)
		if cmd.Flags().Lookup("default").Value.String() == "true" {
			user.SetDefault(&c)
		}
		c.Save()

	},
}

func init() {
	userCmd.AddCommand(userAddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userAddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userAddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	userAddCmd.Flags().BoolP("input", "i", false, "使用交互式输入")
	userAddCmd.Flags().BoolP("default", "d", false, "设置为默认用户")
}
