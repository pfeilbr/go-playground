package main

import (
	"fmt"
	"testing"
)

func TestSliceTypeAssertion(t *testing.T) {

	fn := func(v interface{}) {
		switch typ := v.(type) {
		default:
			fmt.Printf("typ: %v\n", typ) // output: typ: [milk eggs cheese]
		}

	}

	items := []string{"milk", "eggs", "cheese"}
	fn(items)

}
