package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

// $ go tool pprof http://localhost:6666/debug/pprof/profile
// $ go tool pprof ./net_http_pprof http://localhost:6666/debug/pprof/profile

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6666", nil))
	}()
	fib(10000)
	runtime.Goexit()
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
