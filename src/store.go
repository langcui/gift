package main

// 送礼物请求数据,省略了送礼人的id
// 同时也是mongodb里存储的送礼流水数据结构
type gift struct {
	AuthorID uint
	Worth    uint
	Time     uint
}

// redis 排行榜数据返回给前端时需要这个数据结构
// 实际存储在redis里的是 authorid为key, 礼物总价值为score的zset
type anchor struct {
	AuthorID   uint
	TotalWorth uint // 礼物总价值
}
