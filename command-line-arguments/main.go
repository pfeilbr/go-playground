package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("os.Args: %#v\n", os.Args[1:])
}
