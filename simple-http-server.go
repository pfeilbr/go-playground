package main

import (
	"fmt"
	"net/http"
	//"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("zname", "pfeil")
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
