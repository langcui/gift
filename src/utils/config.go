package utils

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("db")
	viper.SetConfigType("toml")
	viper.AddConfigPath("../conf/")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("read config failed:", err)
	}
}

// DBConfig for calling redisip...
type DBConfig struct {
	RedisIP   string
	RedisPort int
	MongoIP   string
	MongoPort int
}

// GetRedisIP return redis's ip
func (c *DBConfig) GetRedisIP() (ip string) {
	ip = viper.GetString("redis.ip")
	return
}

// GetRedisPort return redis's port
func (c *DBConfig) GetRedisPort() (port int) {
	port = viper.GetInt("redis.port")
	return
}

// GetMongoIP return mongodb's ip
func (c *DBConfig) GetMongoIP() (ip string) {
	ip = viper.GetString("mongodb.ip")
	return
}

// GetMongoPort return mongodb's port
func (c *DBConfig) GetMongoPort() (port int) {
	port = viper.GetInt("mongodb.port")
	return
}
