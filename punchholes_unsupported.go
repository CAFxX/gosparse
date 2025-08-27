//go:build !linux

package sparse

func punchHoles(holes []Hole) error {
	return punchHolesFallback(holes)
}
