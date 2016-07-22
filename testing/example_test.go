package exampletest

import "testing"

func TestExample(t *testing.T) {
	if false {
		t.Error("no pass")
	}
}

func TestExample2(t *testing.T) {
	if true {
		t.Error("no pass")
	}
}

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 2 {
		t.Errorf("Add(1, 2) != 2, result = %d", result)
	}
}
