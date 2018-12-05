package triangle

import "math"

type sssTriangle struct {
	AbstractTriangle
}

func newSSSTriangle(a side, b side, c side) (triangle *sssTriangle) {
	triangle = &sssTriangle{}
	triangle.A = a
	triangle.B = b
	triangle.C = c
	triangle.kind = sss
	return
}

func (triangle *sssTriangle) Solve() {
	a := float64(triangle.A)
	b := float64(triangle.B)
	c := float64(triangle.C)

	triangle.Alpha = angle(math.Acos((math.Pow(b, 2) + math.Pow(c, 2) - math.Pow(a, 2)) / (2 * b * c)))
	triangle.Beta = angle(math.Acos((math.Pow(c, 2) + math.Pow(a, 2) - math.Pow(b, 2)) / (2 * c * a)))
	triangle.Gamma = angle(math.Acos((math.Pow(a, 2) + math.Pow(b, 2) - math.Pow(c, 2)) / (2 * a * b)))
}
