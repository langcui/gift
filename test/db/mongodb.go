package mongodb

import (
    "gopkg.in/mgo.v2"
    "time"
)



//global
var GlobalMgoSession *mgo.Session


func init() {
    globalMgoSession, err := mgo.DialWithTimeout("", 10 * time.Second)
    if err != nil {
        panic(err)
    }
    GlobalMgoSession=globalMgoSession
    GlobalMgoSession.SetMode(mgo.Monotonic, true)
    //default is 4096
    GlobalMgoSession.SetPoolLimit(300)
}

func CloneSession() *mgo.Session {
    return GlobalMgoSession.Clone()
}
