package main

import (
	"log"
	"os"
	"os/exec"
)

// go build
// ./exit ./exit.sh
// ./exit ./exit.sh 1
// ./exit ./exit.sh 2
// ./exit ls -a

func main() {
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	log.Println(cmd)

	err := cmd.Start()
	if nil != err {
		log.Fatalln(err)
	}

	err = cmd.Wait()
	if nil != err {
		if "exit status 2" == err.Error() {
			log.Println("Catch 'exit status 2'")
		} else {
			log.Fatalln(err)
		}
	}
}
