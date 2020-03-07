package path

import (
	"image"
)

func squareDistancePointToLine(p, a, z image.Point) float64 {
	// First figure out the perpendicular point q.
	az := image.Point{X: z.X - a.X, Y: z.Y - a.Y}
	ap := image.Point{X: p.X - a.X, Y: p.Y - a.Y}
	t := float64(az.X*ap.X+az.Y*ap.Y) / float64(az.X*az.X+az.Y*az.Y)
	q := image.Point{
		X: a.X + int(t*float64(az.X)),
		Y: a.Y + int(t*float64(az.Y)),
	}

	// Calculate square distance from p to q.
	pq := image.Point{X: q.X - p.X, Y: q.Y - p.Y}
	return float64(pq.X*pq.X + pq.Y*pq.Y)
}

func SimplifyIntPoints(path []image.Point, epsilon float64) []image.Point {
	sqDist := func(p, a, z int) float64 {
		return squareDistancePointToLine(path[p], path[a], path[z])
	}
	indices := Simplify(sqDist, len(path), epsilon*epsilon)
	out := make([]image.Point, 0, len(indices))
	for _, i := range indices {
		out = append(out, path[i])
	}
	return out
}