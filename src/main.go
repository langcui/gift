package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/send", send)
	http.HandleFunc("/top", top)
	http.HandleFunc("/gift_log", giftLog)
	http.HandleFunc("/get_worth", getWorth)
	http.ListenAndServe(":8080", nil)
}

func send(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("ALLOW", "POST")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	data, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	var g Gift
	err := json.Unmarshal([]byte(data), &g)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	g.Time = uint(time.Second)
	err = SendGift(&g)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprintf(w, "success")
}

func top(w http.ResponseWriter, r *http.Request) {
	num, err := strconv.Atoi(r.URL.Query().Get("num"))
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	topN, err := GetTopN(num)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(topN)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", b)
}

func giftLog(w http.ResponseWriter, r *http.Request) {
	authorID := r.URL.Query().Get("id")
	id, err := strconv.Atoi(authorID)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	g, err := GetGiftLog(id)
	if err != nil {
		log.Println(g, err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(g)
	if err != nil {
		log.Println(err, g)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", b)
}

func getWorth(w http.ResponseWriter, r *http.Request) {
	authorID := r.URL.Query().Get("id")
	id, err := strconv.Atoi(authorID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	worth, err := GetAuthorWorth(id)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	var a Anchorinfo
	a.AuthorID = uint(id)
	a.TotalWorth = uint(worth)
	b, err := json.Marshal(a)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s", b)
}
