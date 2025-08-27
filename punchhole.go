package sparse

// PunchHole deallocates a range of a file.
// The arguments fd, offset, and size are the file descriptor,
// byte offset, and byte size of the range to deallocate.
//
// The range [offset offset+size) is zeroed out, and deallocated
// if possible (whether deallocation occurs depends on the OS,
// filesystem, and file type). For example, on Windows, the range
// is deallocated only if the file is sparse.
//
// This is a cross-platform implementation that uses build tags
// to select the correct underlying system call.
func PunchHole(fd int, offset int64, size int64) error {
	return punchHole(fd, offset, size)
}
