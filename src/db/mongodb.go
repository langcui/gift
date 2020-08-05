// Package db init a global mgo session
// use CloneSession() to start your mgo operators
package db

import (
	"time"

	"gopkg.in/mgo.v2"
)

// GlobalMgoSession global mgo session
var GlobalMgoSession *mgo.Session

func init() {
	globalMgoSession, err := mgo.DialWithTimeout("", 10*time.Second)
	if err != nil {
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
