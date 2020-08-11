package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"../src/models"
)

// Num send gift num times totally.
var times = flag.Int("n", 10, "send gift num times")

func main() {
	flag.PrintDefaults()
	flag.Parse()
	n := *times

	fail := 0
	success := 0
	i := 0
	begin := time.Now().Unix()
	for ; i < n; i++ {
		if SendGift() {
			success++
		} else {
			fail++
		}
	}

	fmt.Printf("total send[%d], success[%d], fail[%d], cost[%d].\n",
		i, success, fail, time.Now().Unix()-begin)
}

// SendGift send gift to anthor
func SendGift() bool {
	var g models.Gift
	g.AudienceID = uint(RandInt(1000, 2000))
	g.AnchorID = uint(RandInt(1, 100))
	g.Worth = uint(RandInt(1, 10))
	g.Time = uint(time.Now().Unix())

	b, err := json.Marshal(g)
	if err != nil {
		fmt.Println(err, g)
		return false
	}

	resp, err := http.Post("http://localhost:8080/gift/send", "application/json", strings.NewReader(string(b)))
	if err != nil {
		fmt.Println(err, b)
		return false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || "success" != string(body) {
		fmt.Println(err, string(body))
		return false
	}

	return true
}

// RandInt generate a random num between min and max
func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}

	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
