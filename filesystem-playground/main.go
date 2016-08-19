package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// visit each file and directory recursively for the given path
	filepath.Walk("/tmp/", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})

	fmt.Println("done")
}
