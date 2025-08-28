//go:build solaris

package gosparse

import (
	"syscall"
	"unsafe"
)

func punchHole(fd int, offset int64, size int64) error {
	_, _, errno := syscall.Syscall(
		syscall.SYS_FCNTL,
		uintptr(fd),
		syscall.F_FREESP,
		uintptr(unsafe.Pointer(&syscall.Flock_t{
			Whence: 0,
			Start:  offset,
			Len:    size,
		})),
	)
	if errno != 0 {
		return errno
	}
	return nil
}
