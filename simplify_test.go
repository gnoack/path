package path_test

import (
	"image"
	"testing"

	"github.com/gnoack/path"
)

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
		path    []image.Point
		epsilon float64
		want    []int
	}{
		{
			path: []image.Point{
				image.Pt(0, 0),
				image.Pt(1, 5),
				image.Pt(0, 10),
			},
			epsilon: 1.1,
			want:    []int{0, 2},
		},
		{
			path: []image.Point{
				image.Pt(0, 0),
				image.Pt(1, 5),
				image.Pt(0, 10), image.Pt(5, 11), image.Pt(10, 10),
			},
			epsilon: 1.1,
			want:    []int{0, 2, 4},
		},
		{
			path: []image.Point{
				image.Pt(0, 0),
				image.Pt(1, 5),
				image.Pt(0, 10), image.Pt(5, 11), image.Pt(10, 10),
			},
			epsilon: 1.0,
			want:    []int{0, 1, 2, 3, 4},
		},
	} {
		p := path.OfIntPoints(func(i int) (x, y int) { return tc.path[i].X, tc.path[i].Y }, len(tc.path))
		got := path.Simplify(p, tc.epsilon)
		if !sliceEq(got, tc.want) {
			t.Errorf("Simplify %v with ε=%v: got %v, want %v",
				tc.path, tc.epsilon, got, tc.want)
		}
	}
}

func TestSimplifyFloatPoints(t *testing.T) {
	type Float64Pt struct{ X, Y float64 }
	pt := func(x, y float64) Float64Pt { return Float64Pt{x, y} }

	for _, tc := range []struct {
		path    []Float64Pt
		epsilon float64
		want    []int
	}{
		{
			path: []Float64Pt{
				pt(0, 0),
				pt(1, 5),
				pt(0, 10),
			},
			epsilon: 1.1,
			want:    []int{0, 2},
		},
		{
			path: []Float64Pt{
				pt(0, 0),
				pt(1, 5),
				pt(0, 10), pt(5, 11), pt(10, 10),
			},
			epsilon: 1.1,
			want:    []int{0, 2, 4},
		},
		{
			path: []Float64Pt{
				pt(0, 0),
				pt(1, 5),
				pt(0, 10), pt(5, 11), pt(10, 10),
			},
			epsilon: 1.0,
			want:    []int{0, 1, 2, 3, 4},
		},
	} {
		p := path.OfFloatPoints(func(i int) (x, y float64) { return tc.path[i].X, tc.path[i].Y }, len(tc.path))
		got := path.Simplify(p, tc.epsilon)
		if !sliceEq(got, tc.want) {
			t.Errorf("Simplify %v with ε=%v: got %v, want %v",
				tc.path, tc.epsilon, got, tc.want)
		}
	}
}
