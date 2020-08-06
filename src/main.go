package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"./models"
)

func main() {
	http.HandleFunc("/gift/send", send)       // 给主播送礼,同时写入mongodb的流水和redis的收礼排行榜里
	http.HandleFunc("/gift/top", top)         // 主播收礼排行榜, 根据主播收礼价值数从大到小排序,从redis里获取
	http.HandleFunc("/gift/journal", journal) // 查询主播的收礼流水记录，按时间从近到远排序,从mongodb里获取
	http.HandleFunc("/gift/worth", worth)     // 查询主播的礼物总价值,从redis里获取
	http.ListenAndServe(":8080", nil)
}

// 给主播送礼,同时写入mongodb的流水和redis的收礼排行榜里
func send(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("ALLOW", "POST")
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	data, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	var g models.Gift
	if err := json.Unmarshal([]byte(data), &g); err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	g.Time = uint(time.Now().Unix())
	if err := models.SendGift(&g); err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprintf(w, "success")
}

// 主播收礼排行榜, 根据主播收礼价值数从大到小排序,从redis里获取
// socre相同时, redis 默认按key的字典排序
// GET请求,参数num 表示获取前 num 排名, 默认 num = 10
func top(w http.ResponseWriter, r *http.Request) {
	num, err := strconv.Atoi(r.URL.Query().Get("num"))
	if err != nil || num == 0 {
		num = 10
	}

	topN, err := models.GetTopN(num)
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

// 查询主播的收礼流水记录，按时间从近到远排序,从mongodb里获取
// GET请求,参数id表示需要查询的主播id
func journal(w http.ResponseWriter, r *http.Request) {
	authorID := r.URL.Query().Get("id")
	id, err := strconv.Atoi(authorID)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	g, err := models.GetGiftLog(id)
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

// 查询主播的礼物总价值, 从redis里获取
// GET请求,参数id表示需要查询的主播id
func worth(w http.ResponseWriter, r *http.Request) {
	authorID := r.URL.Query().Get("id")
	id, err := strconv.Atoi(authorID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	worth, err := models.GetAuthorWorth(id)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	var a models.Anchorinfo
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
