/**
 * @Description: 连接redis
 * @Author Lee
 * @Date 2023/12/7 15:03
 **/

package database

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	viperInit "github.com/spf13/viper"
	"time"
)

var (
	RedisClient *redis.Pool
)

func ConnectToRedis() {
	host := viperInit.GetString("redis.host")
	port := viperInit.GetString("redis.port")
	//username := viperInit.GetString("redis.username")
	//password := viperInit.GetString("redis.password")
	MaxIdle := viperInit.GetInt("redis.MaxIdle")
	MaxActive := viperInit.GetInt("redis.MaxActive")
	IdleTimeout := viperInit.GetInt("redis.IdleTimeout")

	pool := &redis.Pool{
		MaxIdle:     MaxIdle,                                  // 最大空闲连接数
		MaxActive:   MaxActive,                                // 最大活跃连接数
		IdleTimeout: time.Duration(IdleTimeout) * time.Second, // 空闲连接超时时间
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host+":"+port)
			if err != nil {
				panic(fmt.Errorf("redis连接失败，错误信息：%s", err))
			}
			return c, err
		},
	}
	RedisClient = pool
}
