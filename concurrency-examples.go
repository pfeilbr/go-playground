package main

import (
	"fmt"
	"time"
)

func ready(w string, sec int, c chan string) {
	time.Sleep(time.Duration(sec) * time.Second)
	c <- (w + " is ready!")
}

func main() {
	c := make(chan string)
	fns := []func(){func() { go ready("Tea", 2, c) }, func() { go ready("Coffee", 1, c) }}

	for _, fn := range fns {
		fn()
	}

	counter := 0

L:
	for {
		select {
		case msg := <-c:
			fmt.Println(msg)
			counter++
			if counter == len(fns) {
				break L
			}
		}
	}

}
