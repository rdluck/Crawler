package main

import (
	"runtime"
	"net/http"
	"Crawler/redis"
	"Crawler/analysis"
)

func main() {
	//利用cpu多核处理http请求
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.HandleFunc("/",redis.RedisServer)
	http.HandleFunc("/a",analysis.Analysis)
	http.ListenAndServe(":9527",nil)
}
