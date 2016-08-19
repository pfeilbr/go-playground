package main

import (
	"fmt"
	"os"
	"strings"

	"code.google.com/p/go-uuid/uuid"
)

func main() {
	//name := uuid.New()
	fmt.Println("Hello World", uuid.New())
	fmt.Println("Domain(0)", uuid.Domain(5))
	id := uuid.New()
	//ptr := &id
	fmt.Printf("Hello %T", id)
	fmt.Printf("env %v", os.Environ())

	var v string
	for _, v = range os.Environ() {
		components := strings.Split(v, "=")
		key := components[0]
		value := components[1]
		fmt.Printf("key: %v, value: %v\n", key, value)
	}

}
