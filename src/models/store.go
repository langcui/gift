package models

// Gift 送礼物请求数据,省略了送礼人的id
// 同时也是mongodb里存储的送礼流水数据结构
type Gift struct {
	AudienceID uint // 观众id
	AuthorID   uint // 收到礼物的主播id
	Worth      uint // 礼物价值, 这里可以改造成礼物id, 然后用 礼物id 再查询礼物价值
	Time       uint // 送礼时间, 后台自动生成
}

// Anchorinfo redis 排行榜数据返回给前端时需要这个数据结构
// 实际存储在redis里的是 authorid为key, 礼物总价值为 score 的zset
// score相同时, redis默认按字典排序;
// score相同时, 如果希望按时间排序, 可以将 uint(score) 按如下方式改造
//  a. 按时间倒序: 改成 float(score . timestamp)
//  b. 按时间正序: 改成 float(score . (maxtimestamp - timestamp))
type Anchorinfo struct {
	AuthorID   uint // 主播id
	TotalWorth uint // 礼物总价值
}
