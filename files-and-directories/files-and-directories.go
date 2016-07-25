package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
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

func glob() {
	if matches, err := filepath.Glob("./**/*"); err != nil {
		panic(err)
	} else {
		fmt.Printf("matches = \n%v\n", strings.Join(matches, "\n"))
	}
}

func main() {
	directoryContents()
	glob()
}
