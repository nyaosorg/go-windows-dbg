//go:build !ndebug && windows
// +build !ndebug,windows

package dbg

import (
	"syscall"
	"unsafe"
)

var kernel32 = syscall.NewLazyDLL("kernel32")
var procOutputDebugString = kernel32.NewProc("OutputDebugStringW")

func outputDebugString(s string) error {
	p, err := syscall.UTF16PtrFromString(s)
	if err != nil {
		return err
	}
	procOutputDebugString.Call(uintptr(unsafe.Pointer(p)))
	return nil
}
