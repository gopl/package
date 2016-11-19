package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

// $ go build
// $ ./fabonacci
// $ go tool pprof ./fabonacci cpu.pprof
// (pprof) web

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {

		return x
	}
	return fib(x-1) + fib(x-2)
}
