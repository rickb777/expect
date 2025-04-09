// See https://magefile.org/

//go:build mage

// Build steps for the expect API:
package main

import (
	"github.com/magefile/mage/sh"
)

var Default = Build

func Build() {
	sh.RunV("go", "mod", "download")
	sh.RunV("go", "mod", "tidy")
	sh.RunV("go", "test", "./...")
	sh.RunV("gofmt", "-l", "-w", "-s", ".")
	sh.RunV("go", "vet", "./...")
}
