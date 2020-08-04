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

func send(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("ALLOW", "POST")
		http.Error(w, http.StatusText(405), 405)
		return
	}

	data, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	var g gift
	err := json.Unmarshal([]byte(data), &g)
	b, err := json.Marshal(g)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	fmt.Fprintf(w, "%s", b)
}

func top(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in top")
	fmt.Fprintf(w, "in top")
}

func journal(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in journal")
	fmt.Fprintf(w, "in journal")
}
