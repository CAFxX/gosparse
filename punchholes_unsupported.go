//go:build !linux || !iouring

package gosparse

func punchHoles(holes []Hole) error {
	return punchHolesFallback(holes)
}
