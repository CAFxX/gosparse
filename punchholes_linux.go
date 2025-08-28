//go:build linux

package sparse

import (
	"errors"
	"sync/atomic"
	"syscall"

	"github.com/iceber/iouring-go"
	"golang.org/x/sys/unix"
)

var iouringUnsupported atomic.Bool

func punchHoles(holes []Hole) error {
	const minHolesForIouring = 4
	if iouringUnsupported.Load() || len(holes) < minHolesForIouring {
		return punchHolesFallback(holes)
	}

	iour, err := iouring.New(uint(len(holes)))
	if err != nil {
		if errors.Is(err, syscall.ENOSYS) {
			iouringUnsupported.Store(true)
		}
		return punchHolesFallback(holes)
	}
	defer iour.Close()

	rs := make([]iouring.PrepRequest, 0, len(holes))
	for idx, hole := range holes {
		r := iouring.Fallocate(
			hole.Fd,
			unix.FALLOC_FL_PUNCH_HOLE|unix.FALLOC_FL_KEEP_SIZE,
			hole.Offset,
			hole.Size,
		).WithInfo(&holes[idx])
		rs = append(rs, r)
	}

	req, err := iour.SubmitRequests(rs, nil)
	if err != nil {
		return err
	}
	<-req.Done()

	var nerr int
	for _, r := range req.Requests() {
		if r.Err() != nil {
			nerr++
		}
	}
	errs := make([]error, 0, nerr)
	for _, r := range req.Requests() {
		if err := r.Err(); err != nil {
			holePtr := r.GetRequestInfo().(*Hole)
			errs = append(errs, punchHoleError{hole: *holePtr, err: err})
		}
	}

	return errors.Join(errs...)
}
