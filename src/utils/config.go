package utils

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("db")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./conf")
	viper.AddConfigPath("../conf")
	viper.AddConfigPath("../src/conf")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("read config failed:", err)
	}
}

// DBConfig for viper.Unmarshal
type DBConfig struct {
	Redis RedisConfig
	Mongo MongoConfig
}

// RedisConfig redis's ip and port
type RedisConfig struct {
	IP   string
	Port int
}

// MongoConfig mongodb's ip and port
type MongoConfig struct {
	IP   string
	Port int
}

// GetDBConfig Unmarshal to DBConfig
func (c *DBConfig) GetDBConfig() {
	viper.Unmarshal(c)
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
	ip = viper.GetString("mongo.ip")
	return
}

// GetMongoPort return mongodb's port
func (c *DBConfig) GetMongoPort() (port int) {
	port = viper.GetInt("mongo.port")
	return
}
