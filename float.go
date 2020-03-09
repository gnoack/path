package path

type Float64Pt struct{ X, Y int }

func squareDistanceFloat64PtToLine(p, a, z Float64Pt) float64 {
	// First figure out the perpendicular point q.
	az := Float64Pt{X: z.X - a.X, Y: z.Y - a.Y}
	ap := Float64Pt{X: p.X - a.X, Y: p.Y - a.Y}
	t := float64(az.X*ap.X+az.Y*ap.Y) / float64(az.X*az.X+az.Y*az.Y)
	q := Float64Pt{
		X: a.X + int(t*float64(az.X)),
		Y: a.Y + int(t*float64(az.Y)),
	}

	// Calculate square distance from p to q.
	pq := Float64Pt{X: q.X - p.X, Y: q.Y - p.Y}
	return float64(pq.X*pq.X + pq.Y*pq.Y)
}

func SimplifyFloat64Points(ptAt func(i int) Float64Pt, length int, epsilon float64) []int {
	// We use square distance as a metric and square epsilon too,
	// which saves us from calculating the square root for the
	// distance.
	sqDist := func(p, a, z int) float64 {
		return squareDistanceFloat64PtToLine(ptAt(p), ptAt(a), ptAt(z))
	}
	return Simplify(sqDist, length, epsilon*epsilon)
}
