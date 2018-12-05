package triangle

import "math"

type ssaTriangle struct {
	AbstractTriangle
}

func newSSATriangle(a side, b side, beta angle) (triangle *ssaTriangle) {
	triangle = &ssaTriangle{}
	triangle.A = a
	triangle.B = b
	triangle.Beta = beta
	triangle.kind = ssa
	return
}

func (triangle *ssaTriangle) Solve() {
	beta := float64(triangle.Beta)

	triangle.Alpha = triangle.sineLawForAngle(triangle.B, triangle.Beta, triangle.A)
	alpha := float64(triangle.Alpha)

	gamma := math.Pi - alpha - beta
	triangle.Gamma = angle(gamma)
	triangle.C = triangle.sineLawForSide(triangle.B, triangle.Beta, triangle.Gamma)
}
