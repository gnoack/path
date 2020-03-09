package path

type intPath struct {
	pt  func(i int) (x, y int)
	len int
}

// OfIntPoints creates an abstract path of integer 2d points.
//
// The path has a length len and integer 2d coordinates for each point
// at index i with 0 <= i < len.
func OfIntPoints(pt func(i int) (x, y int), len int) path {
	return &intPath{pt, len}
}

func (ip *intPath) length() int { return ip.len }

func (ip *intPath) squareDistanceToLine(pidx, aidx, zidx int) float64 {
	type IntPt struct{ X, Y int }

	pt := func(x, y int) IntPt { return IntPt{x, y} }
	p := pt(ip.pt(pidx))
	a := pt(ip.pt(aidx))
	z := pt(ip.pt(zidx))

	//                 p
	//                 .
	//                 .
	//                 .90°
	// ---a------------q-------------z---

	// First figure out the perpendicular point q.
	az := IntPt{X: z.X - a.X, Y: z.Y - a.Y}
	ap := IntPt{X: p.X - a.X, Y: p.Y - a.Y}
	t := float64(az.X*ap.X+az.Y*ap.Y) / float64(az.X*az.X+az.Y*az.Y)
	q := IntPt{
		X: a.X + int(t*float64(az.X)),
		Y: a.Y + int(t*float64(az.Y)),
	}

	// Calculate square distance from p to q.
	pq := IntPt{X: q.X - p.X, Y: q.Y - p.Y}
	return float64(pq.X*pq.X + pq.Y*pq.Y)
}
