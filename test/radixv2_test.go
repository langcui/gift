package radixv2_test

import "fmt"
import "log"
import "testing"
import "github.com/mediocregopher/radix.v2/redis"

func TestRadixv2(t *testing.T) {
    fmt.Println("hello radix_v2")
    conn, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    resp := conn.Cmd("HMSET", "album:1", "title", "Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
}
