package dbg

import (
	"syscall"
	"unsafe"
)

var kernel32 = syscall.NewLazyDLL("kernel32")
var outputDebugString = kernel32.NewProc("OutputDebugStringW")

func OutputDebugString(s string) {
	p, err := syscall.UTF16PtrFromString(s)
	if err == nil {
		outputDebugString.Call(uintptr(unsafe.Pointer(p)))
	}
}
