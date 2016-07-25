package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	tickerDuration := time.Second * 1
	t := time.NewTicker(tickerDuration)
	sleepDuration := time.Millisecond * 500

	go func() {
		for {
			select {
			case interval := <-t.C: // receives every `tickerDuration`
				fmt.Printf("%v\n", interval)
			default:
				fmt.Println("sleeping", sleepDuration)
				time.Sleep(sleepDuration) // block for a max of `sleepDuration`
			}
		}

	}()

	// press any key to stop program
	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b) // block until a key is pressed
		t.Stop()
		fmt.Println("exiting")
		return // exit for block and end program
	}

}
