//go:build !windows || ndebug
// +build !windows ndebug

package dbg

func outputDebugString(s string) error { return nil }
