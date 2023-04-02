package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ddzwd/JumpServerSSHClient/instance"
	"gopkg.in/yaml.v2"
)

type Config struct {
	users []User `yaml:"users"`
}

func get_config_path() string {
	// 获取配置文件路径
	expandedPath := instance.CONFIG_FILE
	fmt.Println(expandedPath)
	// if err != nil {
	// 	log.Fatalf("无法获取配置文件路径:%v", err)
	// }
	if len(expandedPath) > 0 && expandedPath[0] == '~' {
		home, err := os.UserHomeDir()
		if err != nil {
			instance.Logger.Fatalf("无法获取用户主目录:%v", err)
		}
		expandedPath = filepath.Join(home, expandedPath[1:])
	}

	return expandedPath
}

func ValidateConfigExist() bool {
	// 检查文件是否存在
	_, err := os.Stat(get_config_path())
	if err != nil {
		if os.IsNotExist(err) {
			// 文件不存在
			return false
		}
	}
	return true
}

func LoadConfig() Config {
	yamlFile, err := ioutil.ReadFile(get_config_path())
	if err != nil {
		instance.Logger.Fatalf("无法打开YAML文件:%v", err)
	}
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		instance.Logger.Fatalf("无法解析YAML文件:%v", err)
	}
	return config

}

func (c *Config) save() {
	configBytes, err := yaml.Marshal(c)
	if err != nil {
		instance.Logger.Fatalf("无法生成YAML数据:%v", err)
	}
	err = ioutil.WriteFile(get_config_path(), configBytes, os.ModePerm)
	if err != nil {
		instance.Logger.Fatalf("无法写入YAML文件:%v", err)
	}
}

func InitConfig() {
	// 初始化文件
	config := Config{users: []User{}}
	config.users = append(config.users, User{
		UniqueId:   "test",
		UserName:   "test",
		Password:   "",
		Host:       "",
		Port:       "",
		SecretKey:  "",
		RsaKeyPath: "",
		Alias:      "",
	})
	config.save()
}
