package mgov2_test
import (
    "fmt"
    "log"
    "time"
    "testing"
    "./db/"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Person struct {
    Name string
    Phone string
}

type User struct {
    Id_ bson.ObjectId `bson:"_id"`
    Name string `bson:"name"`
    Age int `bson:"age"`
    JoinedAt time.Time `bson:"joined_at"`
    Interests []string `bson:"interests"`
}

func TestMongodb(t *testing.T) {
    session := mongodb.MongoSession()
    defer session.Close()

    c := session.DB("test1").C("people")
    if err := c.Insert(&User{
            Id_ : bson.NewObjectId(),
            Name: "Jimmy Kuu",
            Age: 33,
            JoinedAt: time.Now(),
            Interests: []string{"Develop", "Movie"}});

    if err != nil {
        panic(err)
    }

    var users []User
    if err = c.Find(nil).Limit(5).All(&users); err != nil {
        panic(err)
    }
    fmt.Println(users)
}


func TestPerson(t *testing.T) {
    session, err := mgo.Dial("")
    if err != nil {
        panic(err)
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
    c := session.DB("test").C("people")
    err = c.Insert(&Person{"Ale", "+55 53 8116 963"},
                   &Person{"Cla", "+55 53 8402 8510"})
    if err != nil {
        log.Fatal(err)
    }

    result := Person{}
    err = c.Find(bson.M{"name":"Ale"}).One(&result)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Phone", result.Phone)

}
