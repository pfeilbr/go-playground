package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	fsnotify "gopkg.in/fsnotify.v1"
)

// sandbox is a helper to provide a temporary directory
func sandbox(fn func(string)) {
	tmpDir := "./tmp/"
	os.RemoveAll(tmpDir)
	os.MkdirAll(tmpDir, os.ModePerm)
	defer os.RemoveAll(tmpDir)
	fn(tmpDir)
}

func TestWatchWrite(t *testing.T) {

	sandbox(func(tmpDir string) {
		var err error
		var watcher *fsnotify.Watcher

		if watcher, err = fsnotify.NewWatcher(); err != nil {
			t.Fatal(fmt.Errorf("failed to create watcher: %s", err))
		}
		defer watcher.Close()

		result := make(chan interface{})
		timeout := time.Millisecond * 200
		go func() {
			for {
				select {
				case event := <-watcher.Events:
					result <- event.Op
				case err := <-watcher.Errors:
					result <- err
				case <-time.After(timeout):
					result <- fmt.Errorf("event did not occur within timeout of %v", timeout)
				}
			}
		}()

		// create file before adding to watcher
		var f *os.File
		if f, err = os.Create(filepath.Join(tmpDir, "a.txt")); err != nil {
			t.Fatal(err)
		}
		defer f.Close()

		if err = watcher.Add(tmpDir); err != nil {
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
	})

}
