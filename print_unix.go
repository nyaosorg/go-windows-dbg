// +build !windows ndebug

package dbg

const Enabled = false

func OutputDebugString(s string) error { return nil }
