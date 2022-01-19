package utils

import "github.com/astaxie/beego/toolbox"

func init() {
	toolbox.AddHealthCheck("redis", &RedisCheck{})
}