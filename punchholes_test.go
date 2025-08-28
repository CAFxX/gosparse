package gosparse

import (
	"reflect"
	"slices"
	"testing"
)

func TestOptimizeHoles(t *testing.T) {
	cases := map[string]struct {
		in  []Hole
		out []Hole
	}{
		"different fds": {
			[]Hole{{0, 4, 4}, {1, 4, 4}},
			[]Hole{{0, 4, 4}, {1, 4, 4}},
		},
		"identical": {
			[]Hole{{0, 4, 4}, {0, 4, 4}},
			[]Hole{{0, 4, 4}},
		},
		"contiguous": {
			[]Hole{{0, 4, 4}, {0, 8, 4}},
			[]Hole{{0, 4, 8}},
		},
		"contiguous/ooo": {
			[]Hole{{0, 8, 4}, {0, 4, 4}},
			[]Hole{{0, 4, 8}},
		},
		"ooo": {
			[]Hole{{0, 10, 4}, {0, 4, 4}},
			[]Hole{{0, 4, 4}, {0, 10, 4}},
		},
		"single": {
			[]Hole{{0, 4, 4}},
			[]Hole{{0, 4, 4}},
		},
		"empty": {
			[]Hole{},
			[]Hole{},
		},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			out := OptimizeHoles(slices.Clone(c.in))
			if !reflect.DeepEqual(out, c.out) {
				t.Errorf("mismatch\nwant: %+v\ngot:  %+v", c.out, out)
			}
		})
	}
}
