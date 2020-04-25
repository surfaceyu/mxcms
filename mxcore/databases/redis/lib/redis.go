package lib

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/juju/errors"
	"time"
)

var pool struct{
	redis.Pool
}

func InitRedisPool(host string, port int, auth string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   20,
		IdleTimeout: time.Duration(180) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
			if err != nil {
				return nil, errors.Trace(err)
			}
			if auth != "" {
				if _, err := c.Do("AUTH", auth); err != nil {
					c.Close()
					return nil, err
				}
			}
			c.Do("SELECT", 0)
			return c, nil
		},
	}
}