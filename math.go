package raytracer

import "math"

func solveQuadratic(a, b, c float64) (bool, float64, float64) {
	var q, x0, x1, discr float64
	discr = b*b - 4*a*c
	if discr < 0 {
		return false, 0, 0
	} else if discr == 0 {
		x0 = -0.5 * b / a
		x1 = x0
	} else {
		if b > 0 {
			q = -0.5 * (b + math.Sqrt(discr))
		} else {
			q = -0.5 * (b - math.Sqrt(discr))
		}
		x0 = q / a
		x1 = c / q
	}
	if x0 > x1 {
		tmp := x0
		x0 = x1
		x1 = tmp
	}
	return true, x0, x1
}
