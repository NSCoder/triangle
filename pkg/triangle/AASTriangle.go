package triangle

import "math"

type aasTriangle struct {
	AbstractTriangle
}

func newAASTriangle(a side, alpha angle, beta angle) (triangle *aasTriangle) {
	triangle = &aasTriangle{}
	triangle.A = a
	triangle.Alpha = alpha
	triangle.Beta = beta
	triangle.kind = aas
	return
}

func (triangle *aasTriangle) Solve() {
	alpha := float64(triangle.Alpha)
	beta := float64(triangle.Beta)

	triangle.B = triangle.sineLawForSide(triangle.A, triangle.Alpha, triangle.Beta)

	gamma := math.Pi - alpha - beta
	triangle.Gamma = angle(gamma)

	triangle.C = triangle.sineLawForSide(triangle.A, triangle.Alpha, triangle.Gamma)
}
