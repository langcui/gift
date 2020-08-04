package radixv2_test

import "fmt"
import "log"
import "testing"
import "github.com/mediocregopher/radix.v2/redis"

func _TestRadixv2(t *testing.T) {
        fmt.Println("hello radix_v2")
        conn, err := redis.Dial("tcp", "localhost:6379")
        if err != nil {
                log.Fatal(err)
        }
        defer conn.Close()

        resp := conn.Cmd("HMSET", "album:1", "title", "Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
        log.Println(resp)
}


const RedisAuthorTotalGiftWorthKey = "author_gift_worth"
// GetTopN get num of top gift worth
func TestGetTopN(t *testing.T) {
        conn, err := redis.Dial("tcp", "localhost:6379")
        if err != nil {
            log.Fatal(err)
        }
        defer conn.Close()
        
        num := 10

        resp := conn.Cmd("ZREVRANGE", RedisAuthorTotalGiftWorthKey, 0, num, "WITHSCORES")
        log.Println(resp)
        l, _ := resp.List()
        for i, elemStr := range l {
            log.Println(i, elemStr)
        }
}
