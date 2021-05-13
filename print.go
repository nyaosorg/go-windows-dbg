package dbg

import "fmt"

func Print(v ...interface{}) (int, error) {
	s := fmt.Sprint(v...)
	if err := OutputDebugString(s); err != nil {
		return 0, err
	}
	return len(s), nil
}

func Printf(format string, v ...interface{}) (int, error) {
	s := fmt.Sprintf(format, v...)
	if err := OutputDebugString(s); err != nil {
		return 0, err
	}
	return len(s), nil
}

func Println(v ...interface{}) (int, error) {
	s := fmt.Sprintln(v...)
	if err := OutputDebugString(s); err != nil {
		return 0, err
	}
	return len(s), nil
}
