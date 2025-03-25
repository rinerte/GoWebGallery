package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Any page</h1><p>with any text</p>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "<a href=\"t.me/rinerte\">My telegram </a>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
