package instance

import (
	"github.com/sirupsen/logrus"
)

// 配置文件
var CONFIG_FILE string = "~/.jumpclient.yaml"

// 用户信息相关数据
var USRENAME string
var PASSWORD string
var HOST string
var PORT string
var SECRET_KEY string
var RSA_KEY_PATH string
var ALIAS string

// 日志
var Logger = logrus.New()

var LogLevel string = "warn"

var LogLevelMap = map[string]logrus.Level{
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
	"fatal": logrus.FatalLevel,
	"panic": logrus.PanicLevel,
}

func InitLog() {
	// 将LogLevel转换为logrus.Level
	level, ok := LogLevelMap[LogLevel]
	if !ok {
		Logger.Fatalf("Invalid log level: %s", LogLevel)
	}
	Logger.SetLevel(level)
}
