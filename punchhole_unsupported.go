//go:build !linux && !darwin && !freebsd && !windows && !solaris && !illumos && !android && !ios

package sparse

import "errors"

// punchHole is a fallback for unsupported platforms.
func punchHole(_ int, _, _ int64) error {
	return errors.New("punch hole operation not supported on this platform")
}
