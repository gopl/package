package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"
	"strconv"
	"time"
)

/*
参考文档：
1. Executing Commands In Go
http://www.darrencoxall.com/golang/executing-commands-in-go/

*/

func execute(printer func(...interface{}), name string, args ...string) (*exec.Cmd, error) {
	cmd := exec.Command(name, args...)
	stdout, _ := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	// 一边执行，一遍打印输出内容
	go func() {
		rd := bufio.NewReader(stdout)
		for {
			line, err := rd.ReadString('\n')
			if err != nil {
				if io.EOF == err {
					break
				}
				printer(err.Error())
				printer(debug.Stack())
				return
			}
			printer(line)
		}
	}()

	return cmd, nil
}

func executeWithTimeout(seconds float64, printer func(...interface{}), name string, args ...string) error {
	cmd, err := execute(printer, name, args...)
	if nil != err {
		return err
	}
	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case err := <-done:
		return err
	case <-time.After(time.Second * time.Duration(seconds)):
		cmd.Process.Kill()
		return errors.New("timeout")
	}
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile("execute.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)

	seconds, _ := strconv.Atoi(os.Args[1])

	err = executeWithTimeout(float64(seconds), logger.Print, cwd+string(filepath.Separator)+os.Args[2], os.Args[3:]...)
	if nil != err {
		log.Fatalln("Execute failed: " + err.Error())
	}
}
