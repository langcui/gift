实现一个活动收礼排行榜

功能描述：
1. 提供一个送礼接口，用户B给主播A送礼时调用；
2. 提供一个主播收礼排行榜查询接口，根据主播收礼价值数从大到小排序；
3. 提供一个收礼流水查询接口，查询主播的收礼流水记录，按时间从近到远排序；
 

要求：
1. 使用golang原生HTTP框架； 
2. 送礼接口要求使用HTTP POST方法，请求数据使用Json格式；
3. 排行榜查询接口和收礼流水查询接口要求使用HTTP GET方法，响应数据使用Json格式；
4. 要求使用redis做缓存，使用mongo做数据存储；
5. 排行榜接口并发要求超过1000qps；
6. 排行榜要求实时统计；
7. 要求最多只能10个协程访问mongo数据库；