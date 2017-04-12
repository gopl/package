package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Hello, http.HandleFunc!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

//$ curl -s http://127.0.0.1:8080
//Hello, http.HandleFunc!
