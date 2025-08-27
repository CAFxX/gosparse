//go:build windows

package sparse

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// punchHole deallocates a range of a file on Windows.
// It uses the DeviceIoControl system call with the FSCTL_SET_ZERO_DATA control code.
func punchHole(fd int, offset int64, size int64) error {
	zeroInfo := struct { // FILE_ZERO_DATA_INFORMATION
		FileOffset      int64
		BeyondFinalZero int64
	}{
		FileOffset:      offset,
		BeyondFinalZero: offset + size,
	}
	var bytesReturned uint32 // unused
	return syscall.DeviceIoControl(
		syscall.Handle(fd),
		windows.FSCTL_SET_ZERO_DATA,
		(*byte)(unsafe.Pointer(&zeroInfo)),
		uint32(unsafe.Sizeof(zeroInfo)),
		nil,
		0,
		&bytesReturned, // must be non-nil if lpOverlapped is nil
		nil,
	)
}
