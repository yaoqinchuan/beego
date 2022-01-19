package utils

// 使用redis存储
import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/logs"
	"time"
)

var RedisCache cache.Cache // redis cache
var FileCache cache.Cache  // file cache

func init() {
	password, err := Decrypt(beego.AppConfig.String("redis.password"), []byte(EncryptKey))
	if err != nil {
		logs.Error(err)
	}
	cacheRedisConn, _ := json.Marshal(map[string]string{
		"key":      "redisCache",
		"conn":     beego.AppConfig.String("redis.endpoint"),
		"dbNum":    beego.AppConfig.String("redis.db-num"),
		"password": password,
	})
	RedisCache, err = cache.NewCache("redis", string(cacheRedisConn))
	if err != nil {
		logs.Error(err)
	}

	// 配置信息如下所示，配置 CachePath 表示缓存的文件目录，FileSuffix 表示文件后缀，DirectoryLevel 表示目录层级，EmbedExpiry 表示过期设置
	cacheFile, _ := json.Marshal(map[string]string{
		"CachePath":      "./file_cache",
		"FileSuffix":     ".cache",
		"DirectoryLevel": "2",
		"EmbedExpiry":    "120",
	})
	FileCache, err = cache.NewCache("file", string(cacheFile))

	if err != nil {
		logs.Error(err)
	}
	// 用来做redis检查
	err = RedisCache.Put("check", "ok", time.Duration(1000000000))
	if err != nil {
		logs.Error(err)
	}
	value := RedisCache.Get("check")
	if value == nil {
		logs.Error("set value failed.")
	}
}

type RedisCheck struct {
}

func (all *RedisCheck) Check() error {
	value := RedisCache.Get("check")
	if value == nil {
		return errors.New("check redis failed.")
	}
	logs.Info("check redis success")
	return nil
}
