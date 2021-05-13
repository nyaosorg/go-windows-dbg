// +build ndebug

package dbg

func _print(v ...interface{}) (int, error) { return 0, nil }

func _printf(format string, v ...interface{}) (int, error) { return 0, nil }

func _println(v ...interface{}) (int, error) { return 0, nil }
