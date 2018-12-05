package triangle

import "math"

type asaTriangle struct {
	AbstractTriangle
}

func newASATriangle(alpha angle, b side, gamma angle) (triangle *asaTriangle) {
	triangle = &asaTriangle{}
	triangle.Alpha = alpha
	triangle.B = b
	triangle.Gamma = gamma
	triangle.kind = asa
	return
}

func (triangle *asaTriangle) Solve() {
	alpha := float64(triangle.Alpha)
	gamma := float64(triangle.Gamma)
	beta := math.Pi - alpha - gamma
	triangle.Beta = angle(beta)

	a := triangle.sineLawForSide(triangle.B, triangle.Beta, triangle.Alpha)
	triangle.A = side(a)

	c := triangle.sineLawForSide(triangle.B, triangle.Beta, triangle.Gamma)
	triangle.C = side(c)
}
