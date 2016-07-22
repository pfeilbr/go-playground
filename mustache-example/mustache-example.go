package main

import (
	"fmt"
	"github.com/hoisie/mustache"
)

func main() {
	data := mustache.Render("hello {{c}}", map[string]string{"c": "world"})
	fmt.Printf("data = %s", data)
}
