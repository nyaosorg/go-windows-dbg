//go:build ndebug
// +build ndebug

package dbg

// Enable is the switch to output to OutputDebugStringW
const Enabled = false

func _print(v ...interface{}) (int, error) { return 0, nil }

func _printf(format string, v ...interface{}) (int, error) { return 0, nil }

func _println(v ...interface{}) (int, error) { return 0, nil }

func _X[T any](v T) T {
	return v
}
