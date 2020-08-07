package models

import (
	"log"
	"strconv"

	"../db"
	"gopkg.in/mgo.v2/bson"
)

// RedisAnchorTotalGiftWorthKey key fot top
const RedisAnchorTotalGiftWorthKey = "anchor_gift_worth"

// SendGift send a Gift to anchor
func SendGift(g *Gift) error {
	if err := UpdateAnchorGiftWorth(g); err != nil {
		return err
	}

	if err := AddGiftLog(g); err != nil {
		return err
	}

	return nil
}

// UpdateAnchorGiftWorth incr anchor's Gift worth to redis
func UpdateAnchorGiftWorth(g *Gift) error {
	conn, err := db.RedisPool().Get()
	if err != nil {
		return err
	}
	defer db.RedisPool().Put(conn)

	if err = conn.Cmd("ZINCRBY", RedisAnchorTotalGiftWorthKey, g.Worth, g.AnchorID).Err; err != nil {
		return err
	}

	return nil
}

// GetAnchorWorth from redis
func GetAnchorWorth(anchorID int) (int, error) {
	conn, err := db.RedisPool().Get()
	if err != nil {
		return 0, err
	}
	defer db.RedisPool().Put(conn)
	worth, err := conn.Cmd("ZSCORE", RedisAnchorTotalGiftWorthKey, anchorID).Int()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return worth, nil
}

// GetTopN get num of top Gift worth from redis
func GetTopN(num int) ([]Anchorinfo, error) {
	conn, err := db.RedisPool().Get()
	if err != nil {
		return nil, err
	}
	defer db.RedisPool().Put(conn)

	data := conn.Cmd("ZREVRANGE", RedisAnchorTotalGiftWorthKey, 0, num-1, "WITHSCORES")
	l, _ := data.List()
	var arr [2]string
	var arrAnchor []Anchorinfo
	for i, elemStr := range l {
		arr[i%2] = elemStr
		if i%2 == 1 {
			anchorID, _ := strconv.Atoi(arr[0])
			worth, _ := strconv.Atoi(arr[1])
			au := Anchorinfo{uint(anchorID), uint(worth)}
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

// GetGiftLog get anchor's Gift log from mongodb
func GetGiftLog(anchorID int) ([]Gift, error) {
	session := db.MongoSession()
	defer session.Close()

	c := session.DB(MongodbGiftDB).C(MongodbJournalCollection)
	var gifts []Gift
	f := bson.M{"anchorid": anchorID}
	// err := c.Find(f).Limit(MongodbMaxPageNum).All(&gifts).sort(bson.M{"time": 1})
	if err := c.Find(f).Sort("-time").Limit(MongodbMaxPageNum).All(&gifts); err != nil {
		log.Println(err)
		return nil, err
	}
	return gifts, nil
}

// AddGiftLog add gift log to mongodb
func AddGiftLog(g *Gift) error {
	session := db.MongoSession()
	defer session.Close()

	if err := session.DB(MongodbGiftDB).C(MongodbJournalCollection).Insert(g); err != nil {
		return err
	}

	return nil
}
