package triangle

import "math"

type sasTriangle struct {
	AbstractTriangle
}

func newSASTriangle(a side, beta angle, c side) (triangle *sasTriangle) {
	triangle = &sasTriangle{}
	triangle.A = a
	triangle.Beta = beta
	triangle.C = c
	triangle.kind = sas
	return
}

func (triangle *sasTriangle) Solve() {
	a := float64(triangle.A)
	c := float64(triangle.C)
	beta := float64(triangle.Beta)

	b := math.Sqrt(math.Pow(a, 2) + math.Pow(c, 2) - (2 * a * c * math.Cos(beta)))
	triangle.B = side(b)

	triangle.Alpha = triangle.sineLawForAngle(triangle.B, triangle.Beta, triangle.A)
	triangle.Gamma = triangle.sineLawForAngle(triangle.B, triangle.Beta, triangle.C)
}
