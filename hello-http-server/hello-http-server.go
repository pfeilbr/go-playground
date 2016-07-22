package main

import (
	"fmt"
	"log"
	"net/http"
)

type hello struct{}

func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h hello
	addrport := ":4000"

	fmt.Printf("Server listening on %v ...\n", addrport)
	if err := http.ListenAndServe(addrport, h); err != nil {
		log.Fatalf("Failed to start server on %v. Error: %v", addrport, err)
	}

}
