package main

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/mitchellh/go-ps"
)

func TestProcesses(t *testing.T) {
	p, err := ps.Processes()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if len(p) <= 0 {
		t.Fatal("should have processes")
	}

	found := false
	var executables []string
	for _, p1 := range p {
		executables = append(executables, p1.Executable())
		if p1.Executable() == "go" || p1.Executable() == "go.exe" {
			found = true
			//break
		}
	}

	sort.Strings(executables)

	fmt.Printf("%d processes\nexecutables:\n%s\n", len(executables), strings.Join(executables, "\n"))

	if !found {
		t.Fatal("should have Go")
	}
}
