package main

import (
	"fmt"
	"net/http"
)

func main() {
	// rooted paths
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "/favicon.ico")
	})

	// rooted subtrees
	http.HandleFunc("/images/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "/images/")
	})

	// 分别注册 /thumbnails 和 /thumbnails/
	http.HandleFunc("/thumbnails/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "/thumbnails/")
	})
	http.HandleFunc("/thumbnails", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "/thumbnails")
	})

	http.ListenAndServe(":8080", nil)
}

/*
https://golang.org/pkg/net/http/#ServeMux

~$ curl -s http://127.0.0.1:8080
404 page not found
~$ curl -s http://127.0.0.1:8080/
404 page not found


~$ curl -s http://127.0.0.1:8080/favicon.ico
/favicon.ico


~$ # ServeMux 会将 /images 重定向到 /images/
~$ curl -s http://127.0.0.1:8080/images/
/images/
~$ curl -sv http://127.0.0.1:8080/images
<a href="/images/">Moved Permanently</a>.

...
< HTTP/1.1 301 Moved Permanently
< Location: /images/
< Date: Wed, 12 Apr 2017 07:47:37 GMT
< Content-Length: 43
< Content-Type: text/html; charset=utf-8
<
{ [43 bytes data]
* Connection #0 to host 127.0.0.1 left intact
~$ curl -sL http://127.0.0.1:8080/images
/images/


~$ # 存在 /thumbnails，ServeMux 不会重定向 /thumbnails 到 /thumbnails/
~$ curl -s http://127.0.0.1:8080/thumbnails/
/thumbnails/
~$ curl -s http://127.0.0.1:8080/thumbnails
/thumbnails

*/
