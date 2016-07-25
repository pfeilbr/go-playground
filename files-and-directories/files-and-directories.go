package main

import (
	"fmt"
	"io/ioutil"
)

func directoryContents() {
	list, err := ioutil.ReadDir("../../")
	if err != nil {
		panic(err)
	}

	for _, fi := range list {
		fmt.Printf("Name = %v, Size = %v bytes\n", fi.Name(), fi.Size())
	}
}

func main() {
	directoryContents()
}
