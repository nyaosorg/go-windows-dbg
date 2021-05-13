// +build !ndebug

package dbg

import "fmt"

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
