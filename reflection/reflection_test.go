package main

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestAnonFunc(t *testing.T) {
	expect := reflect.Func
	fn := func() int { return 1 }
	val := reflect.ValueOf(fn)
	result := val.Kind()

	if result != expect {
		t.Errorf("expect: %v, result: %v", expect, result)
	}
}

func TestSliceAnonStruct(t *testing.T) {
	expect := reflect.Slice
	people := []struct {
		Name string
		age  int
	}{
		{"Brian", 38},
		{"Tricia", 42},
	}
	result := reflect.ValueOf(people).Kind()
	if result != expect {
		t.Errorf("expect: %v, result: %v", expect, result)
	}
}

func TestEmptyStruct(t *testing.T) {
	expect := reflect.Struct
	// empty struct
	emptyStruct := struct{}{}

	result := reflect.ValueOf(emptyStruct).Kind()
	if result != expect {
		t.Errorf("expect: %v, result: %v", expect, result)
	}
}

func TestEmptyStructSize(t *testing.T) {
	expect := uint(0)         // takes up 0 bytes
	emptyStruct := struct{}{} // empty struct
	result := uint(unsafe.Sizeof(emptyStruct))
	if result != expect {
		t.Errorf("expect: %v, result: %v", expect, result)
	}
}
