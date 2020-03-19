package redis

import (
	"admin-server/config"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	pool      *redis.Pool
	redisHost = config.RedisHost
	redisPass = config.RedisPasswd
)

// 创建 redis 连接池
func newRedisPool() *redis.Pool {
	return &redis.Pool{
		// 建立链接
		Dial: func() (redis.Conn, error) {
			// 打开链接
			c, err := redis.Dial("tcp", redisHost)
			if err != nil {
				fmt.Print(err.Error())
				return nil, err
			}
			// 认证
			if _, err = c.Do("AUTH", redisPass); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		// 存活检测
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("PING")
			return err
		},
		MaxIdle:     50,
		MaxActive:   30,
		IdleTimeout: 300 * time.Second,
		Wait:        false,
	}
}

func init() {
	pool = newRedisPool()
}

func RedisPool() *redis.Pool {
	return pool
}
