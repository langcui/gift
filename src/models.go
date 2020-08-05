package main

import (
	"log"
	"strconv"

	"./db"
	"github.com/mediocregopher/radix.v2/pool"
	"gopkg.in/mgo.v2/bson"
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

// SendGift send a Gift to author
func SendGift(g *Gift) error {
	err := UpdateAuthorGiftWorthRedis(g)
	if err != nil {
		return err
	}

	err = AddGiftLogMongo(g)
	if err != nil {
		return err
	}

	return nil
}

// UpdateAuthorGiftWorthRedis incr author's Gift worth
func UpdateAuthorGiftWorthRedis(g *Gift) error {
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

// GetTopN get num of top Gift worth
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

// MongodbGiftDB db gift
const MongodbGiftDB = "GiftDB"

// MongodbJournalCollection GiftDB.JournalCollection
const MongodbJournalCollection = "JournalCollection"

// MongodbMaxPageNum max items num each page
const MongodbMaxPageNum = 10

// GetGiftLog get author's Gift log
func GetGiftLog(authorID int) ([]Gift, error) {
	session := db.CloneSession()
	defer session.Close()

	c := session.DB(MongodbGiftDB).C(MongodbJournalCollection)
	var gifts []Gift
	f := bson.M{"authorid": authorID}
	// err := c.Find(f).Limit(MongodbMaxPageNum).All(&gifts).sort(bson.M{"time": 1})
	err := c.Find(f).Sort("-time").Limit(MongodbMaxPageNum).All(&gifts)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return gifts, nil
}

// AddGiftLogMongo add gift log to mongodb
func AddGiftLogMongo(g *Gift) error {
	session := db.CloneSession()
	defer session.Close()

	err := session.DB(MongodbGiftDB).C(MongodbJournalCollection).Insert(g)
	if err != nil {
		return err
	}

	return nil
}
