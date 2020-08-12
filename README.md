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


压测:
1. 使用 go-stress-test:
https://github.com/link1st/go-stress-testing

```
-c 表示并发数
-n 每个并发执行请求的次数，总请求的次数 = 并发数 * 每个并发执行请求的次数
-u 需要压测的地址
cuiyc@hw-sg-mildom-docker1:~/projects$ ./go-stress-testing-linux -c 100 -n 10000 -u http://localhost:8080/gift/top

 开始启动  并发数:100 请求数:10000 请求参数:
request:
 form:http
 url:http://localhost:8080/gift/top
 method:GET
 headers:map[]
 data:
 verify:statusCode
 timeout:30s
 debug:false
 ```

结果: qps 4745    
```
| 耗时 │ 并发数│ 成功数│ 失败数 │ qps │ 最长耗时 │ 最短耗时 │ 平均耗时 │ 错误码 |
| :---- | ----: | ----: | ----: | ----: | ----: | ----: | ----: | :---- |
|  1s │    100│   3952│      0│ 4910.08│   66.03│    1.10│    0.20│200:3952|
|  2s │    100│   7858│      0│ 4785.78│   66.03│    1.10│    0.21│200:7858|
|  3s │    100│  11757│      0│ 4818.58│   66.03│    1.10│    0.21│200:11757|
|  4s │    100│  15674│      0│ 4747.23│   71.86│    1.10│    0.21│200:15674|
|  5s │    100│  19557│      0│ 4748.92│   71.86│    1.10│    0.21│200:19557|
|  6s │    100│  23414│      0│ 4751.52│   71.86│    1.10│    0.21│200:23414|  
|  7s │    100│  27477│      0│ 4776.52│   71.86│    1.10│    0.21│200:27477|
| 255s│    100│ 996262│      0│ 4738.40│  168.58│    0.53│    0.21│200:996262|
| 256s│    100│1000000│      0│ 4745.32│  168.58│    0.53│    0.21│200:1000000|

*************************  结果 stat****************************  
处理协程数量: 100  
请求总数（并发数*请求数 -c * -n）: 1000000 总请求时间: 256.001 秒 successNum: 1000000 failureNum: 0  
*************************  结果 end   ****************************  
```

2. 运行测试代码,同时观察性能
```
./go-stress-testing-linux -c 1000 -n 1000 -u http://localhost:8080/gift/add_test_data
写操作:添加测试数据

cpu 性能:
cuiyc@hw-sg-mildom-docker1:~/projects$ go tool pprof http://localhost:8080/debug/pprof/profile?seconds=60
Fetching profile over HTTP from http://localhost:8080/debug/pprof/profile?seconds=60
Saved profile in /home1/cuiyc/pprof/pprof.gift.samples.cpu.001.pb.gz
File: gift
Type: cpu
Time: Aug 12, 2020 at 7:34pm (CST)
Duration: 1mins, Total samples = 1.09mins (109.10%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top10
Showing nodes accounting for 36620ms, 55.84% of 65580ms total
Dropped 640 nodes (cum <= 327.90ms)
Showing top 10 nodes out of 226
      flat  flat%   sum%        cum   cum%
   10820ms 16.50% 16.50%    11380ms 17.35%  syscall.Syscall
    6930ms 10.57% 27.07%     6930ms 10.57%  math/rand.seedrand (inline)
    5330ms  8.13% 35.19%     5880ms  8.97%  runtime.step
    3180ms  4.85% 40.04%     3180ms  4.85%  runtime.futex
    2560ms  3.90% 43.95%     2560ms  3.90%  runtime.epollwait
    2320ms  3.54% 47.48%     8310ms 12.67%  runtime.pcvalue
    1760ms  2.68% 50.17%     2030ms  3.10%  syscall.Syscall6
    1450ms  2.21% 52.38%     1450ms  2.21%  runtime.usleep
    1370ms  2.09% 54.47%     8310ms 12.67%  math/rand.(*rngSource).Seed
     900ms  1.37% 55.84%      900ms  1.37%  runtime.memmove
     
内存性能:
cuiyc@hw-sg-mildom-docker1:~$ go tool pprof http://localhost:8080/debug/pprof/heap
Fetching profile over HTTP from http://localhost:8080/debug/pprof/heap
Saved profile in /home1/cuiyc/pprof/pprof.gift.alloc_objects.alloc_space.inuse_objects.inuse_space.001.pb.gz
File: gift
Type: inuse_space
Time: Aug 12, 2020 at 7:39pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 9752.91kB, 100% of 9752.91kB total
Showing top 10 nodes out of 18
      flat  flat%   sum%        cum   cum%
 3598.02kB 36.89% 36.89%  3598.02kB 36.89%  bufio.NewReaderSize (inline)
 2570.01kB 26.35% 63.24%  2570.01kB 26.35%  bufio.NewWriterSize (inline)
 2048.75kB 21.01% 84.25%  2048.75kB 21.01%  runtime.malg
  512.08kB  5.25% 89.50%   512.08kB  5.25%  net/http.(*Server).newConn (inline)
  512.03kB  5.25% 94.75%   512.03kB  5.25%  context.WithCancel
  512.02kB  5.25%   100%  7192.08kB 73.74%  net/http.(*conn).serve
         0     0%   100%  3598.02kB 36.89%  bufio.NewReader (inline)
         0     0%   100%   512.08kB  5.25%  main.main
         0     0%   100%   512.08kB  5.25%  net/http.(*Server).ListenAndServe
         0     0%   100%   512.08kB  5.25%  net/http.(*Server).Serve
```
![写数据时的 cpu 占比图](https://github.com/langcui/gift/blob/master/image/profile_add_test_data_cpu.png)

![写数据时的 mem 占比图](https://github.com/langcui/gift/blob/master/image/profile_add_test_data_mem.png)


因为添加测试数据的时候,伪造了一个随机用户id, 系统调用 rand 占比较大, 这个地方可以优化成在最外层设置随机种子.
