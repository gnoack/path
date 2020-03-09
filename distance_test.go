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
		points := []image.Point{tc.p, tc.a, tc.z}
		p := OfIntPoints(func(i int) (x, y int) { return points[i].X, points[i].Y }, len(points))
		got := p.squareDistanceToLine(0, 1, 2)
		if got != tc.want {
			t.Errorf("Square distance from point %v to line %v-%v: got %v, want %v",
				tc.p, tc.a, tc.z, got, tc.want)
		}
	}
}
