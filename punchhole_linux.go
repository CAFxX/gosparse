//go:build linux

package sparse

import (
	"golang.org/x/sys/unix"
)

// punchHole deallocates a range of a file on Linux.
// It uses the fallocate system call with the FALLOC_FL_PUNCH_HOLE flag.
func punchHole(fd int, offset int64, size int64) error {
	// FALLOC_FL_PUNCH_HOLE requires FALLOC_FL_KEEP_SIZE to be set.
	// See fallocate(2) for more details.
	return unix.Fallocate(
		fd,
		unix.FALLOC_FL_PUNCH_HOLE|unix.FALLOC_FL_KEEP_SIZE,
		offset,
		size,
	)
}
