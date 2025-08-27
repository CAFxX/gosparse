package sparse

import (
	"cmp"
	"errors"
	"fmt"
	"slices"
)

type Hole struct {
	Fd     int   // File descriptor
	Offset int64 // Byte offset of the hole
	Size   int64 // Size of the hole in bytes
}

// PunchHoles performs multiple punch hole operations.
//
// All holes are processed, and any errors encountered
// are collected and returned as a single error at the end.
func PunchHoles(holes []Hole) error {
	if len(holes) == 0 {
		return nil // No holes to punch
	} else if len(holes) == 1 {
		return punchHole(holes[0].Fd, holes[0].Offset, holes[0].Size)
	}
	return punchHoles(punchHolesOpt(holes))
}

type punchHoleError struct {
	hole Hole
	err  error
}

func (e punchHoleError) Error() string {
	return fmt.Sprintf("punch hole fd: %d, offset: %d, size: %d, error: %s", e.hole.Fd, e.hole.Offset, e.hole.Size, e.err)
}

func (e punchHoleError) Unwrap() error {
	return e.err
}

func punchHolesFallback(holes []Hole) error {
	var errs []error
	for _, hole := range holes {
		if err := PunchHole(hole.Fd, hole.Offset, hole.Size); err != nil {
			errs = append(errs, punchHoleError{hole: hole, err: err})
		}
	}
	return errors.Join(errs...)
}

func punchHolesOpt(holes []Hole) []Hole {
	if len(holes) <= 1 {
		return holes
	}
	slices.SortFunc(holes, func(a, b Hole) int {
		return cmp.Or(cmp.Compare(a.Fd, b.Fd), cmp.Compare(a.Offset, b.Offset))
	})
	var i, j int
	for i, j = 0, 1; j < len(holes); j++ {
		if holes[i].Fd == holes[j].Fd && holes[i].Offset+holes[i].Size >= holes[j].Offset {
			// Merge overlapping or contiguous holes
			holes[i].Size = max(holes[i].Size, holes[j].Offset+holes[j].Size-holes[i].Offset)
		} else {
			i++
			holes[i] = holes[j]
		}
	}
	holes = holes[:i+1]
	return holes
}
