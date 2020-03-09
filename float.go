package path

type floatPath struct {
	pt  func(i int) (x, y float64)
	len int
}

// OfFloatPoints creates an abstract path of float64 2d points.
//
// The path has a length len and float64 2d coordinates for each point
// at index i with 0 <= i < len.
func OfFloatPoints(pt func(i int) (x, y float64), len int) path {
	return &floatPath{pt, len}
}

func (fp *floatPath) length() int { return fp.len }

func (fp *floatPath) squareDistanceToLine(pidx, aidx, zidx int) float64 {
	type Float64Pt struct{ X, Y float64 }

	pt := func(x, y float64) Float64Pt { return Float64Pt{x, y} }
	p := pt(fp.pt(pidx))
	a := pt(fp.pt(aidx))
	z := pt(fp.pt(zidx))

	//                 p
	//                 .
	//                 .
	//                 .90°
	// ---a------------q-------------z---

	// First figure out the perpendicular point q.
	az := Float64Pt{X: z.X - a.X, Y: z.Y - a.Y}
	ap := Float64Pt{X: p.X - a.X, Y: p.Y - a.Y}
	t := (az.X*ap.X + az.Y*ap.Y) / (az.X*az.X + az.Y*az.Y)
	q := Float64Pt{
		X: a.X + t*az.X,
		Y: a.Y + t*az.Y,
	}

	// Calculate square distance from p to q.
	pq := Float64Pt{X: q.X - p.X, Y: q.Y - p.Y}
	return float64(pq.X*pq.X + pq.Y*pq.Y)
}
