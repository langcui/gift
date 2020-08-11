package models

// Gift 送礼物请求数据,省略了送礼人的id
// 同时也是mongodb里存储的送礼流水数据结构
type Gift struct {
	AudienceID uint `json:"audience_id"` // 观众id
	AnchorID   uint `json:"auchor_id"`   // 收到礼物的主播id
	Worth      uint `json:"worth"`       // 礼物价值, 这里可以改造成礼物id, 然后用 礼物id 再查询礼物价值
	Time       uint `json:"time"`        // 送礼时间, 后台自动生成
}

// Anchorinfo redis 排行榜数据返回给前端时需要这个数据结构
// 实际存储在redis里的是 anchorid为key, 礼物总价值为 score 的zset
// score相同时, redis默认按字典排序;
// score相同时, 如果希望按时间排序, 可以将 uint(score) 按如下方式改造
//  a. 按时间倒序: 改成 float(score . timestamp)
//  b. 按时间正序: 改成 float(score . (maxtimestamp - timestamp))
type Anchorinfo struct {
	AnchorID   uint `json:"anchor_id"`   // 主播id
	TotalWorth uint `json:"total_worth"` // 礼物总价值
}

// Response for http response body
type Response struct {
	Code    uint32      `json:"code"`              // 0 表示正常，其他值表示出错
	Message string      `json:"message,omitempty"` // Code为非0时, Message表示出错信息
	Data    interface{} `json:"data,omitempty"`    // 返回的json数据
}
