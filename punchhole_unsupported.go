//go:build !linux && !darwin && !windows && !solaris && !illumos && !android && !ios

package gosparse

import "errors"

// punchHole is a fallback for unsupported platforms.
func punchHole(_ int, _, _ int64) error {
	return errors.New("punch hole operation not supported on this platform")
}
