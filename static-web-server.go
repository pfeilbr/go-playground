package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var host = flag.String("host", ":8080", "host to bind to")
var root = flag.String("root", "", "webserver document root folder")
var cert = flag.String("cert", "", "tls certificate")
var key = flag.String("key", "", "tls certificate key")
var useTLS = flag.Bool("tls", false, "enable TLS")

func main() {

	flag.Parse()

	// HTTP
	if !*useTLS {
		fmt.Printf("listening on %s // root=%s\r\n", *host, *root)
		err := http.ListenAndServe(*host, http.FileServer(http.Dir(*root)))
		if err != nil {
			log.Printf("error : %v", err)
		}
	}

	// HTTPS
	if *useTLS {
		fmt.Printf("listening on %s // cert=%s, key=%s, root=%s\r\n", *host, *cert, *key, *root)
		err := http.ListenAndServeTLS(*host, *cert, *key, http.FileServer(http.Dir(*root)))
		if err != nil {
			log.Printf("error : %v", err)
		}
	}
}
