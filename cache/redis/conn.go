package conn

import (
	"flag"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	pool          *redis.Pool
	redisServer   = flag.String("redis-srv", ":6379", "")
	redisPassword = flag.String("redis-pwd", "123456", "")
)

func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     10,
		MaxActive:   10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			// open connection
			c, err := redis.Dial("tcp", *redisServer)
			if err != nil {
				return nil, err
			}
			// auth password
			if _, err := c.Do("AUTH", *redisPassword); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		// check health status
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func init() {
	pool = newRedisPool()
}

func RedisConnPool() *redis.Pool {
	return pool
}
