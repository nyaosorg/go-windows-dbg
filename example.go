// +build run

package main

import (
	"github.com/zetamatta/go-windows-dbg"
)

func main() {
	dbg.Print("output", 1, "text")
	dbg.Printf("output<%d>text", 1)
	dbg.Println("output", 1, "text")
}
