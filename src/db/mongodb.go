// Package db init a global mgo session
// use MongoSession() to start your mgo operators
package db

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"../utils"
	"gopkg.in/mgo.v2"
)

// GlobalMgoSession global mgo session
var GlobalMgoSession *mgo.Session

// InitMongo : server ip and port from config file
func InitMongo() {
	var c utils.DBConfig
	url := c.GetMongoIP() + ":" + strconv.Itoa(c.GetMongoPort())
	fmt.Println(url)
	globalMgoSession, err := mgo.DialWithTimeout(url, 10*time.Second)
	if err != nil {
		log.Println(err, url)
		panic(err)
	}
	GlobalMgoSession = globalMgoSession
	GlobalMgoSession.SetMode(mgo.Monotonic, true)
	GlobalMgoSession.SetPoolLimit(300)
}

// CloneSession return a clone of global mgo session
func CloneSession() *mgo.Session {
	return GlobalMgoSession.Clone()
}

// MongoSession return a clone of global mgo session
func MongoSession() *mgo.Session {
	return CloneSession()
}
