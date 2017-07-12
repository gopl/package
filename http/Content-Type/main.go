package main

import (
	"encoding/json"
	"net/http"
)

type user struct {
	Name  string
	Email string
}

func handler(w http.ResponseWriter, r *http.Request) {
	u := user{"custa", "custa@126.com"}
	//w.WriteHeader(http.StatusOK) // 导致 Content-Type 设置不生效
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(&u)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
