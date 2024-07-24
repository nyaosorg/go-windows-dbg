package dbg

func Print(v ...interface{}) (int, error) {
	return _print(v...)
}

func Printf(format string, v ...interface{}) (int, error) {
	return _printf(format, v...)
}

func Println(v ...interface{}) (int, error) {
	return _println(v...)
}

func OutputDebugString(s string) error {
	return outputDebugString(s)
}

func X[T any](v T) T {
	return _X(v)
}
