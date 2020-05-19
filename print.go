package dbg

import "fmt"

func Print(v ...interface{}) {
	OutputDebugString(fmt.Sprint(v...))
}

func Printf(format string, v ...interface{}) {
	OutputDebugString(fmt.Sprintf(format, v...))
}

func Println(v ...interface{}) {
	OutputDebugString(fmt.Sprintln(v...))
}
