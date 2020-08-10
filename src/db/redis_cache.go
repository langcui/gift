package db

import (
	"fmt"
	"log"
	"strconv"

	"../utils"
	"github.com/mediocregopher/radix.v2/pool"
)

// GlobalRedisPool global redis pool
var GlobalRedisPool *pool.Pool

// InitRedis : redis ip and port from config file
func InitRedis() {
	var c utils.DBConfig
	url := c.GetRedisIP() + ":" + strconv.Itoa(c.GetRedisPort())
	fmt.Println(url)

	var err error
	GlobalRedisPool, err = pool.New("tcp", url, 10)
	if err != nil {
		log.Println(err, url)
		panic(err)
	}
}

// RedisPool return the global redis pool
func RedisPool() *pool.Pool {
	return GlobalRedisPool
}
