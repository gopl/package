package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	defer func() {
		fmt.Println("defer")
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 100)
		}
	}()

	fmt.Println("Goexit")
	runtime.Goexit()
}
