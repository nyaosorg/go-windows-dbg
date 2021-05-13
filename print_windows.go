// +build !ndebug,windows

package dbg

import (
	"syscall"
	"unsafe"
)

// Enable is the switch to output to OutputDebugStringW
const Enabled = true

var kernel32 = syscall.NewLazyDLL("kernel32")
var outputDebugString = kernel32.NewProc("OutputDebugStringW")

func OutputDebugString(s string) error {
	p, err := syscall.UTF16PtrFromString(s)
	if err != nil {
		return err
	}
	outputDebugString.Call(uintptr(unsafe.Pointer(p)))
	return nil
}
