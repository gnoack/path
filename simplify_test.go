package path

import (
	"testing"
)

func TestDistance(t *testing.T) {
	for _, tc := range []struct {
		p, a, z IntPt
		want    float64
	}{
		{IntPt{4, 1}, IntPt{0, 0}, IntPt{10, 0}, 1},  // horizontal line
		{IntPt{1, 5}, IntPt{0, 0}, IntPt{0, 10}, 1},  // vertical line
		{IntPt{1, 1}, IntPt{0, 0}, IntPt{10, 10}, 0}, // on line
	} {
		got := squareDistanceIntPtToLine(tc.p, tc.a, tc.z)
		if got != tc.want {
			t.Errorf("Square distance from point %v to line %v-%v: got %v, want %v",
				tc.p, tc.a, tc.z, got, tc.want)
		}
	}
}

func sliceEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, ae := range a {
		if ae != b[i] {
			return false
		}
	}
	return true
}

func TestSimplify(t *testing.T) {
	for _, tc := range []struct {
		path    []IntPt
		epsilon float64
		want    []int
	}{
		{
			path: []IntPt{
				IntPt{0, 0},
				IntPt{1, 5},
				IntPt{0, 10},
			},
			epsilon: 1.1,
			want:    []int{0, 2},
		},
		{
			path: []IntPt{
				IntPt{0, 0},
				IntPt{1, 5},
				IntPt{0, 10}, IntPt{5, 11}, IntPt{10, 10},
			},
			epsilon: 1.1,
			want:    []int{0, 2, 4},
		},
		{
			path: []IntPt{
				IntPt{0, 0},
				IntPt{1, 5},
				IntPt{0, 10}, IntPt{5, 11}, IntPt{10, 10},
			},
			epsilon: 1.0,
			want:    []int{0, 1, 2, 3, 4},
		},
	} {
		got := SimplifyIntPoints(func(i int) IntPt { return tc.path[i] }, len(tc.path), tc.epsilon)
		if !sliceEq(got, tc.want) {
			t.Errorf("Simplify %v with ε=%v: got %v, want %v",
				tc.path, tc.epsilon, got, tc.want)
		}
	}
}
