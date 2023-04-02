package config

import (
	"fmt"
)

type User struct {
	UniqueId   string `yaml:"unique_id"`
	UserName   string `yaml:"name"`
	Password   string `yaml:"password"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	SecretKey  string `yaml:"secret_key"`
	RsaKeyPath string `yaml:"rsa_key_path"`
	Alias      string `yaml:"alias"`
}

func AddUser(u User) {
	// 添加用户

	config := LoadConfig()
	config.users = append(config.users, u)
	config.save()
}

func (u *User) Delete() {
	// 删除用户

	config := LoadConfig()
	for i, user := range config.users {
		if user.UniqueId == u.UniqueId {
			config.users = append(config.users[:i], config.users[i+1:]...)
			break
		}
	}
	config.save()
}

func (u *User) Update() {
	// 更新用户
	config := LoadConfig()
	for i, user := range config.users {
		if user.UniqueId == u.UniqueId {
			config.users[i] = *u
			break
		}
	}
	config.save()
}

func ListUsers() {
	// 列出所有用户
	config := LoadConfig()
	for _, user := range config.users {
		fmt.Println(user)
	}
}

func GetUserByNameOrAlias(name string) *User {
	// 根据用户名或别名获取用户
	config := LoadConfig()
	for _, user := range config.users {
		if user.UserName == name || user.Alias == name {
			return &user
		}
	}
	return nil
}
