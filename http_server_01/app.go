package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Path -> [%s]\n", r.URL.Path)
	fmt.Fprintf(w, "Proto -> [%s]\n", r.Proto)
	fmt.Fprintf(w, "Host -> [%s]\n", r.Host)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%s] -> %s\n", k, v)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
