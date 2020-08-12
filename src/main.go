package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"./models"
	"./utils"

	_ "net/http/pprof"
)

func main() {
	models.InitDB()
	http.HandleFunc("/gift/send", send)                 // 给主播送礼,同时写入mongodb的流水和redis的收礼排行榜里
	http.HandleFunc("/gift/top", top)                   // 主播收礼排行榜, 根据主播收礼价值数从大到小排序,从redis里获取
	http.HandleFunc("/gift/journal", journal)           // 查询主播的收礼流水记录，按时间从近到远排序,从mongodb里获取
	http.HandleFunc("/gift/worth", worth)               // 查询主播的礼物总价值,从redis里获取
	http.HandleFunc("/gift/config", config)             // 获取配置文件, 目前只有db的配置文件
	http.HandleFunc("/gift/add_test_data", addTestData) // 添加测试数据, 用于压测

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
		log.Println(err, data)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	g.Time = uint(time.Now().Unix())
	if err := models.SendGift(&g); err != nil {
		log.Println(err, g)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	resp := models.Response{Code: 0, Message: "success", Data: nil}
	RespJSON(w, resp)
}

// RespJSON return json to client
func RespJSON(w http.ResponseWriter, r models.Response) {
	b, err := json.Marshal(r)
	if err != nil {
		log.Println(err, r)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
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
		log.Println(err, num)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	resp := models.Response{Code: 0, Message: "success", Data: topN}
	RespJSON(w, resp)
}

// 查询主播的收礼流水记录，按时间从近到远排序,从mongodb里获取
// GET请求,参数id表示需要查询的主播id
func journal(w http.ResponseWriter, r *http.Request) {
	anchorID := r.URL.Query().Get("id")
	id, err := strconv.Atoi(anchorID)
	if err != nil {
		log.Println(err, anchorID)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	g, err := models.GetGiftLog(id)
	if err != nil {
		log.Println(err, id)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	resp := models.Response{Code: 0, Message: "success", Data: g}
	RespJSON(w, resp)
}

// 查询主播的礼物总价值, 从redis里获取
// GET请求,参数id表示需要查询的主播id
func worth(w http.ResponseWriter, r *http.Request) {
	anchorID := r.URL.Query().Get("id")
	id, err := strconv.Atoi(anchorID)
	if err != nil {
		log.Println(err, anchorID)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	worth, err := models.GetAnchorWorth(id)
	if err != nil {
		log.Println(err, id)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	var a models.Anchorinfo
	a.AnchorID = uint(id)
	a.TotalWorth = uint(worth)

	resp := models.Response{Code: 0, Message: "success", Data: a}
	RespJSON(w, resp)
}

func config(w http.ResponseWriter, r *http.Request) {
	var c utils.DBConfig
	c.GetDBConfig()

	resp := models.Response{Code: 0, Message: "success", Data: c}
	RespJSON(w, resp)
}

// addTestData 添加一个随机测试数据, 用于压测等场景
func addTestData(w http.ResponseWriter, r *http.Request) {

	// RandInt generate a random num between min and max
	fRandInt := func(min, max int) int {
		if min >= max || min == 0 || max == 0 {
			return max
		}

		rand.Seed(time.Now().UnixNano())
		return rand.Intn(max-min) + min
	}

	var g models.Gift
	g.AudienceID = uint(fRandInt(1000, 2000))
	g.AnchorID = uint(fRandInt(1, 100))
	g.Worth = uint(fRandInt(1, 10))
	g.Time = uint(time.Now().Unix())

	if err := models.SendGift(&g); err != nil {
		log.Println(err, g)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	resp := models.Response{Code: 0, Message: "success", Data: nil}
	RespJSON(w, resp)
}
