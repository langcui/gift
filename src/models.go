package main

import (
	"log"
	"strconv"

	"github.com/mediocregopher/radix.v2/pool"
)

var rdb *pool.Pool

func init() {
	var err error
	rdb, err = pool.New("tcp", "localhost:6379", 10)
	if err != nil {
		log.Println(err)
	}
}

// RedisAuthorTotalGiftWorthKey key fot top
const RedisAuthorTotalGiftWorthKey = "author_gift_worth"

// SendGift send a gift to author
func SendGift(g *gift) error {
	err := UpdateAuthorGiftWorthRedis(g)
	if err != nil {
		return err
	}

	return nil
}

// UpdateAuthorGiftWorthRedis incr author's gift worth
func UpdateAuthorGiftWorthRedis(g *gift) error {
	conn, err := rdb.Get()
	if err != nil {
		return err
	}
	defer rdb.Put(conn)

	err = conn.Cmd("ZINCRBY", RedisAuthorTotalGiftWorthKey, g.Worth, g.AuthorID).Err
	if err != nil {
		return err
	}

	return nil
}

// GetAuthorWorth from redis
func GetAuthorWorth(authorID int) (int, error) {
	conn, err := rdb.Get()
	if err != nil {
		return 0, err
	}
	defer rdb.Put(conn)
	worth, err := conn.Cmd("ZSCORE", RedisAuthorTotalGiftWorthKey, authorID).Int()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return worth, nil
}

// GetTopN get num of top gift worth
func GetTopN(num int) ([]Anchorinfo, error) {
	conn, err := rdb.Get()
	if err != nil {
		return nil, err
	}
	defer rdb.Put(conn)

	data := conn.Cmd("ZREVRANGE", RedisAuthorTotalGiftWorthKey, 0, num, "WITHSCORES")
	l, _ := data.List()
	var arr [2]string
	var arrAnchor []Anchorinfo
	for i, elemStr := range l {
		arr[i%2] = elemStr
		if i%2 == 1 {
			authorID, _ := strconv.Atoi(arr[0])
			worth, _ := strconv.Atoi(arr[1])
			au := Anchorinfo{uint(authorID), uint(worth)}
			arrAnchor = append(arrAnchor, au)
		}
	}

	return arrAnchor, nil
}
