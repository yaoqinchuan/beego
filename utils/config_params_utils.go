package utils

import (
	"github.com/astaxie/beego"
)

// 用来读取配置的，可以强制读取某个环境的配置
func GetStringConfig(configName string) string {
	return beego.AppConfig.String(configName)
}

// 用来修改配置的
func UpdateStringConfig(key, value string) {
	_ = beego.AppConfig.Set(key, value)
}
