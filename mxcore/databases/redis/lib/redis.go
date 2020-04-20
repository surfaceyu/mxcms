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

func InitRedisPool(host string, port int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   20,
		IdleTimeout: time.Duration(180) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
			if err != nil {
				return nil, errors.Trace(err)
			}
			c.Do("SELECT", 0)
			return c, nil
		},
	}
}