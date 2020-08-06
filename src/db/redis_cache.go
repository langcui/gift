package db

import (
	"log"

	"github.com/mediocregopher/radix.v2/pool"
)

// GlobalRedisPool global redis pool
var GlobalRedisPool *pool.Pool

func init() {
	var err error
	GlobalRedisPool, err = pool.New("tcp", "localhost:6379", 10)
	if err != nil {
		log.Println(err)
	}
}

// RedisPool return the global redis pool
func RedisPool() *pool.Pool {
	return GlobalRedisPool
}
