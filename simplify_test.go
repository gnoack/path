package path

import (
	"image"
	"testing"
)

func TestDistance(t *testing.T) {
	for _, tc := range []struct {
		p, a, z image.Point
		want    float64
	}{
		{image.Pt(4, 1), image.Pt(0, 0), image.Pt(10, 0), 1},  // horizontal line
		{image.Pt(1, 5), image.Pt(0, 0), image.Pt(0, 10), 1},  // vertical line
		{image.Pt(1, 1), image.Pt(0, 0), image.Pt(10, 10), 0}, // on line
	} {
		got := squareDistancePointToLine(tc.p, tc.a, tc.z)
		if got != tc.want {
			t.Errorf("Square distance from point %v to line %v-%v: got %v, want %v",
				tc.p, tc.a, tc.z, got, tc.want)
		}
	}
}

func sliceEq(a, b []image.Point) bool {
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
		want    []image.Point
	}{
		{
			path: []image.Point{
				image.Pt(0, 0),
				image.Pt(1, 5),
				image.Pt(0, 10),
			},
			epsilon: 1.1,
			want: []image.Point{
				image.Pt(0, 0),
				image.Pt(0, 10),
			},
		},
		{
			path: []image.Point{
				image.Pt(0, 0),
				image.Pt(1, 5),
				image.Pt(0, 10), image.Pt(5, 11), image.Pt(10, 10),
			},
			epsilon: 1.1,
			want: []image.Point{
				image.Pt(0, 0),
				image.Pt(0, 10), image.Pt(10, 10),
			},
		},
		{
			path: []image.Point{
				image.Pt(0, 0),
				image.Pt(1, 5),
				image.Pt(0, 10), image.Pt(5, 11), image.Pt(10, 10),
			},
			epsilon: 1.0,
			want: []image.Point{
				image.Pt(0, 0),
				image.Pt(1, 5),
				image.Pt(0, 10), image.Pt(5, 11), image.Pt(10, 10),
			},
		},
	} {
		got := SimplifyIntPoints(tc.path, tc.epsilon)
		if !sliceEq(got, tc.want) {
			t.Errorf("Simplify %v with ε=%v: got %v, want %v",
				tc.path, tc.epsilon, got, tc.want)
		}
	}
}
