package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	go func() {
		fmt.Println("before")
		runtime.Goexit()
		fmt.Println("after")
	}()

	fmt.Println(runtime.GOROOT(), runtime.GOMAXPROCS(0), runtime.NumCPU(), runtime.NumCgoCall())
	<-time.After(time.Second * 2)
}
