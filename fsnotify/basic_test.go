package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	fsnotify "gopkg.in/fsnotify.v1"
)

func TestWatchWrite(t *testing.T) {

	tmpDir := "./tmp/"

	os.RemoveAll(tmpDir)
	os.MkdirAll(tmpDir, os.ModePerm)
	defer os.RemoveAll(tmpDir)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		t.Fatal(err)
	}

	if watcher == nil {
		t.Fatal("failed to create watcher.  watcher is nil")
	}
	defer watcher.Close()

	result := make(chan interface{})
	timeout := time.Millisecond * 200
	go func() {
		for {
			select {
			case event := <-watcher.Events:

				if event.Op&fsnotify.Write == fsnotify.Write {
					result <- fsnotify.Write
				} else {
					result <- event.Op
				}
			case err := <-watcher.Errors:
				result <- err
			case <-time.After(timeout):
				result <- fmt.Errorf("event did not occur within timeout of %v", timeout)
			}
		}
	}()

	// create file before adding to watcher
	f, err := os.Create(filepath.Join(tmpDir, "a.txt"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	err = watcher.Add(tmpDir)
	if err != nil {
		t.Fatal(err)
	}

	// perform write
	f.Write([]byte("hello"))

	// wait on result
	res := <-result

	expect := fsnotify.Write
	switch res.(type) {
	case fsnotify.Op:
		if res != fsnotify.Write {
			t.Errorf("expect: %v, result: %v", expect, res)
		}
	case error:
		t.Error(res)
	default:
		t.Fatal("unexpected response")
	}

}
