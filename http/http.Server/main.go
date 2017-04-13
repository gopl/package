package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, s)
}

func main() {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        String("custom http.Server"),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
