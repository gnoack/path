// Package path simplifies paths of 2-dimensional points.
package path

// A distance function which calculates a distance metric between the
// point x[p] and the line through x[a] and x[z]. p, a and z are
// indices for the given points in the input to Simplify.
type DistanceFunc func(p, a, z int) float64

// Simplify finds representative points in the given path and returns
// their indices. The path is given through its length and the
// DistanceFunc calculating distances on its points.
func Simplify(d DistanceFunc, length int, epsilon float64) []int {
	if length <= 0 {
		return []int{}
	}
	result := rdp(d, 0, length-1, epsilon)
	return append(result, length-1)
}

// Returns the all simplified indices within [a, z).
// Note: This excludes z, which is convenient for concatenating results.
//
// This is an implementation of the Ramer-Douglas-Peucker algorithm.
func rdp(d DistanceFunc, a, z int, epsilon float64) []int {
	if z-a < 0 {
		return []int{}
	}
	if z-a == 0 {
		return []int{}
	}
	if z-a == 1 {
		return []int{a}
	}
	// z-a >= 2 (at least one in the middle)

	maxdist := 0.0
	maxp := a + 1
	for p := a + 1; p < z; p++ {
		dist := d(p, a, z)
		if dist >= maxdist {
			maxdist = dist
			maxp = p
		}
	}

	if maxdist < epsilon {
		// There is no point further away than epsilon
		// from the straight line a-z, so we just skip
		// all middle points.
		return []int{a}
	}

	subpath1 := rdp(d, a, maxp, epsilon)
	subpath2 := rdp(d, maxp, z, epsilon)

	return append(subpath1, subpath2...)
}
