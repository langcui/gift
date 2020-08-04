package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/send", send)
	http.HandleFunc("/top", top)
	http.HandleFunc("/journal", journal)
	http.ListenAndServe(":8080", nil)
}

func send(reps http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		fmt.Fprintf(reps, "use POST only.")
		return
	}

	data, _ := ioutil.ReadAll(req.Body)
	req.Body.Close()
	var g gift
	err := json.Unmarshal([]byte(data), &g)
	fmt.Println(err)
	fmt.Printf("post data:%s, gift.AuthorID: %d.\n", data, g.AuthorID)
	fmt.Fprintf(reps, "in send, gift:AuthorID:%d, Worth:%d, Time:%d",
		g.AuthorID, g.Worth, g.Time)
}

func top(reps http.ResponseWriter, req *http.Request) {
	fmt.Println("in top")
	fmt.Fprintf(reps, "in top")
}

func journal(reps http.ResponseWriter, req *http.Request) {
	fmt.Println("in journal")
	fmt.Fprintf(reps, "in journal")
}
