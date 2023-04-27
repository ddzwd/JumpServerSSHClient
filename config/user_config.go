package config

import (
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
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

// 添加用户
func (u *User) AddUser(config *Config) {
	config.Users = append(config.Users, *u)
}

// 设置默认用户
func (u *User) SetDefault(c *Config) {
	c.DefaultUser = u.UniqueId
}

// 删除用户
func (u *User) Delete(config *Config) {
	var users []User
	for _, user := range config.Users {
		if user.UniqueId != u.UniqueId {
			users = append(users, user)
		}
	}
	config.Users = users
	if config.DefaultUser == u.UniqueId {
		config.DefaultUser = ""
	}
}

// 更新用户 并保存
func (u *User) Update(config *Config) {
	for i, user := range config.Users {
		if user.UniqueId == u.UniqueId {
			config.Users[i] = *u
			break
		}
	}
}

// 获取默认用户
func GetDefaultUser(config *Config) *User {
	for _, user := range config.Users {
		if user.UniqueId == config.DefaultUser {
			return &user
		}
	}
	return nil
}

// 列出所有用户
func ListUsers(config *Config) *[]User {
	return &config.Users
}

// 根据用户名或别名获取用户
// FIXME: 此方法存在问题,我们应该允许相同用户名
func GetUserByNameOrAlias(name string, config Config) *User {
	for _, user := range config.Users {
		if user.UserName == name || user.Alias == name {
			return &user
		}
	}
	return nil
}

func GetUserById(id string, config *Config) *User {
	for _, user := range config.Users {
		if user.UniqueId == id {
			return &user
		}
	}
	return nil
}

// 查找用户,通过用户名,别名,主机,端口 列出所有用户
func GrepUser(content string, config *Config) *[]User {
	var users []User
	for _, user := range config.Users {
		if strings.Contains(user.UserName, content) {
			users = append(users, user)
			continue
		}
		if strings.Contains(user.Alias, content) {
			users = append(users, user)
			continue
		}
		if strings.Contains(user.Host, content) {
			users = append(users, user)
			continue
		}
		if strings.Contains(user.Port, content) {
			users = append(users, user)
		}
	}
	return &users
}

// 格式化打印
// | username | password | rsa_key_path | secret_key | server | port |
// |----------|----------|--------------|------------|--------|------|
// |          |          |              |            |        |      |
// |          |          |              |            |        |      |
func PrettyPrint(users *[]User) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Username", "Password", "RsaKeyPath", "SecretKey", "Host", "Port"})
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetRowLine(true)
	// 打印每个用户的行
	for _, u := range *users {
		table.Append([]string{u.UniqueId, u.UserName, u.Password, u.RsaKeyPath, u.SecretKey, u.Host, u.Port})
	}
	// 刷新 tabwriter，确保所有缓冲内容都被写入
	table.Render()
}
