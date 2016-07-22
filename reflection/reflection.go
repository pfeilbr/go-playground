package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

func main() {

	// anonymous function
	fn := func() int { return 1 }
	fmt.Println(reflect.TypeOf(fn))

	// anonymous struct
	people := []struct {
		Name string
		age  int
	}{
		{"Brian", 38},
		{"Tricia", 42},
	}
	fmt.Println(reflect.TypeOf(people))

	// empty struct
	emptyStruct := struct{}{}
	fmt.Println(reflect.TypeOf(emptyStruct))
	fmt.Println(unsafe.Sizeof(emptyStruct)) // 0. takes up no space

	sleepDurationInSeconds := 2 * time.Second
	fmt.Printf("sleeping for %v seconds", sleepDurationInSeconds)
	time.Sleep(sleepDurationInSeconds)
}
