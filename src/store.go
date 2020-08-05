package main

// 送礼物请求数据,省略了送礼人的id
// 同时也是mongodb里存储的送礼流水数据结构
type Gift struct {
	AuthorID uint
	Worth    uint
	Time     uint
}

// Anchorinfo redis 排行榜数据返回给前端时需要这个数据结构
// 实际存储在redis里的是 authorid为key, 礼物总价值为score的zset
// 用来返回数据给前端
type Anchorinfo struct {
	AuthorID   uint
	TotalWorth uint // 礼物总价值
}
