package redis_test

import (
        "fmt"
        "testing"
        "github.com/garyburd/redigo/redis"
)

func TestRedis(t *testing.T) {
        c, err := redis.Dial("tcp", "127.0.0.1:6379")
        if err != nil {
                fmt.Println("Connect to redis error", err)
                return
        }
        defer c.Close()

        _, err = c.Do("SET", "mykey", "langcui")
        if err != nil {
            fmt.Println("redis set failed:", err)
            return
        }

        username, err := redis.String(c.Do("GET", "mykey"))
        if err != nil {
            fmt.Println("redis get failed:", err)
        } else {
            fmt.Printf("Get mykey: %v \n", username)
        }
}
