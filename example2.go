//go:build run
// +build run

package main

import (
	"github.com/nyaosorg/go-windows-dbg"
)

func main() {
	println(dbg.X(1 + 2))
}
