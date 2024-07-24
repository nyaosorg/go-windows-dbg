//go:build !ndebug
// +build !ndebug

package dbg

import (
	"fmt"
	"runtime"
)

// Enable is the switch to output to OutputDebugStringW
const Enabled = true

func _print(v ...interface{}) (int, error) {
	s := fmt.Sprint(v...)
	if err := OutputDebugString(s); err != nil {
		return 0, err
	}
	return len(s), nil
}

func _printf(format string, v ...interface{}) (int, error) {
	s := fmt.Sprintf(format, v...)
	if err := OutputDebugString(s); err != nil {
		return 0, err
	}
	return len(s), nil
}

func _println(v ...interface{}) (int, error) {
	s := fmt.Sprintln(v...)
	if err := OutputDebugString(s); err != nil {
		return 0, err
	}
	return len(s), nil
}

func _X[T any](v T) T {
	_, f, l, _ := runtime.Caller(1)
	Printf("[%s:%d] %#v\n", f, l, v)
	return v
}
