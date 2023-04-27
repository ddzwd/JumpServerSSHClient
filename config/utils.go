package config

import (
	"os"
	"path/filepath"

	"github.com/ddzwd/JumpServerSSHClient/instance"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Path        string `yaml:"-"`
	Users       []User `yaml:"users"`
	DefaultUser string `yaml:"default_user"`
}

func get_config_path(expandedPath string) string {
	if len(expandedPath) > 0 && expandedPath[0] == '~' {
		home, err := os.UserHomeDir()
		if err != nil {
			instance.Logger.Fatalf("无法获取用户主目录:%v", err)
		}
		expandedPath = filepath.Join(home, expandedPath[1:])
	}

	return expandedPath
}

func ValidateConfigExist(path string) bool {
	// 检查文件是否存在
	_, err := os.Stat(get_config_path(path))
	if err != nil {
		if os.IsNotExist(err) {
			// 文件不存在
			return false
		}
	}
	return true
}

func LoadConfig(path string) Config {
	p := get_config_path(path)
	yamlFile, err := os.ReadFile(p)
	if err != nil {
		instance.Logger.Fatalf("无法打开YAML文件:%v", err)
	}
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		instance.Logger.Fatalf("无法解析YAML文件:%v", err)
	}
	config.Path = p
	return config

}

func (c *Config) Save() {
	instance.Logger.Debugf("configBytes: %v", c)
	configBytes, err := yaml.Marshal(c)
	if err != nil {
		instance.Logger.Fatalf("无法生成YAML数据:%v", err)
	}
	instance.Logger.Debugf("configBytes: %v", string(configBytes))
	err = os.WriteFile(c.Path, configBytes, os.ModePerm)
	if err != nil {
		instance.Logger.Fatalf("无法写入YAML文件:%v", err)
	}
}

func InitConfig(path string) {
	// 初始化文件
	config := Config{Users: []User{}, Path: get_config_path(path)}
	config.Save()
}
