package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

var memprofile = flag.String("memprofile", "", "write memory profile to this file")

func main() {
	flag.Parse()
	var memFile *os.File
	if *memprofile != "" {
		var err error
		memFile, err = os.Create(*memprofile)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("start write heap profile....")
			pprof.WriteHeapProfile(memFile)
			defer memFile.Close()
		}
	}

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go work(&wg, memFile)
	}
	wg.Wait()

	time.Sleep(3 * time.Second)
}

func work(wg *sync.WaitGroup, memFile *os.File) {
	time.Sleep(time.Second)

	var counter int
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Millisecond * 100)
		pprof.WriteHeapProfile(memFile)
		counter++
	}
	wg.Done()
}
