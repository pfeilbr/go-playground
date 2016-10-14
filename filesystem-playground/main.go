package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	paths := make(chan string)
	done := make(chan struct{})
	// visit each file and directory recursively for the given path
	producer := func() {
		filepath.Walk("/etc/", func(path string, info os.FileInfo, err error) error {
			paths <- path
			return nil
		})
		close(paths)
	}
	go producer()

	consumer := func() {
		for {
			select {
			case path, more := <-paths:
				if more {
					fmt.Println(path)
					os.Stdout.Sync() // flush output to ensure live progress
				} else {
					done <- struct{}{}
				}
			}
		}
	}
	go consumer()

	<-done
}
