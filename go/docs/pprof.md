#####  启动`pprof` HTTP服务
```go
// 开启pprof
var pprofPort = 6666
if os.Getenv("GOPPROF") == "on" {
  // 启用阻塞分析
  runtime.SetBlockProfileRate(1)
  go func() {
    log.Printf("go pprof start listen on 0.0.0.0:%d", pprofPort)
    _ = http.ListenAndServe(fmt.Sprintf(":%d", pprofPort), nil)
  }()
  // ...
}
```

##### 使用`go tool pprof`采集数据

```bash
# 采集 CPU Profiling 数据，采样时间60s
# 输出 pprof.rtp.samples.cpu.001.pb.gz
go tool pprof http://localhost:6666/debug/pprof/profile?seconds=60

# 采集 Memory Profiling 数据  
# 输出 pprof.rtp.alloc_objects.alloc_space.inuse_objects.inuse_space.001.pb.gz
go tool pprof http://localhost:6666/debug/pprof/heap

# 采集 Block/Mutex Profiling 数据  
# 输出 pprof.rtp.contentions.delay.001.pb.gz
go tool pprof http://localhost:6666/debug/pprof/block
go tool pprof http://localhost:6666/debug/pprof/mutex

# 采集 Goroutine Profiling 数据  
# 输出 pprof.rtp.goroutine.001.pb.gz
go tool pprof http://localhost:6666/debug/pprof/goroutine


# 采集 Threadcreate Profiling 数据 
# 输出 pprof.rtp.threadcreate.001.pb.gz
go tool pprof http://localhost:6666/debug/pprof/threadcreate

```

##### 使用`go tool pprof`分析数据

1. **启动HTTP服务**: `go tool pprof -http=0.0.0.0:8181 /root/pprof/pprof.rtp.samples.cpu.001.pb.gz`

2. **在浏览器中访问**:`http://localhost:8181`

##### 指标分析

1. **CPU Profiling**
2. **Memory Profiling**
3. **Block Profiling**
4. **Goroutine Profiling**
5. **Threadcreate Profiling**