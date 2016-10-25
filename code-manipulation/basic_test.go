package main

import (
	"fmt"
	"go/build"
	"testing"
)

func TestBuildContext(t *testing.T) {

	ctxt := &build.Default
	fmt.Printf("ctxt.SrcDirs(): %v\n", ctxt.SrcDirs())

	pkg, _ := ctxt.Import("github.com/pfeilbr/mytool", ".", build.ImportComment)
	fmt.Printf("pkg: %#v\n", pkg)

}
