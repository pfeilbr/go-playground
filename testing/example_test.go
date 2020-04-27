package exampletest

import "testing"

func TestExample(t *testing.T) {
	if false {
		t.Error("no pass")
	}
}

func TestExample2(t *testing.T) {
	if false {
		t.Error("no pass")
	}
}

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("Add(1, 2) != 2, result = %d", result)
	}
}

func TestSub(t *testing.T) {
	expect := 1
	result := Sub(2, 1)
	if result != expect {
		t.Errorf("expect %v. got %v", expect, result)
	}
}

func TestFoo(t *testing.T) {
	if false {
		t.Error("Encountered a failure")
	}
}
