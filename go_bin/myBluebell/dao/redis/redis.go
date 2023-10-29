package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var rdb *redis.Client

func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", viper.GetString("redis.host")),
		Password: "",                              // 密码
		DB:       viper.GetInt("redis.db"),        // 数据库
		PoolSize: viper.GetInt("redis.pool_size"), // 连接池大小
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	} else {
		println("redis connection successfully")
	}
	return nil
}
func Close() {
	_ = rdb.Close()
}
