//go:build darwin || ios

package gosparse

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

// punchHole deallocates a range of a file on Darwin (macOS) and iOS.
// It uses the fcntl system call with the F_PUNCHHOLE command.
func punchHole(fd int, offset int64, size int64) error {
	_, _, errno := syscall.Syscall(
		syscall.SYS_FCNTL,
		uintptr(fd),
		unix.F_PUNCHHOLE,
		uintptr(unsafe.Pointer(&struct { // fpunchhole_t
			Flags  uint32 // No flags are currently defined.
			_      uint32 // Reserved padding.
			Offset uint64 // Offset in bytes.
			Length uint64 // Length in bytes.
		}{
			Offset: uint64(offset),
			Length: uint64(size),
		})),
	)

	if errno != 0 {
		return errno
	}
	return nil
}
