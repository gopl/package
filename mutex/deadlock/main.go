package main

import (
	"fmt"
	"sync"
	"time"
)

var mux sync.Mutex

func foo() {
	mux.Lock()
	defer mux.Unlock()
	bar()
}

func bar() {
	mux.Lock()
	defer mux.Unlock()
	fmt.Println("in bar")
}

func main() {
	go foo()
	time.Sleep(3 * time.Second)
}
