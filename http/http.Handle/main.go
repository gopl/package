package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, s)
}

func main() {
	http.Handle("/", String("Hello, http.Handle!"))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

//$ curl -s http://127.0.0.1:8080
//Hello, http.Handle!
