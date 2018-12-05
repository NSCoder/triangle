package triangle

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"strings"
)

type side float64
type angle float64

func (triangle AbstractTriangle) isEqual(otherTriangle AbstractTriangle) bool {
	return triangle.A.isEqual(otherTriangle.A) &&
		triangle.B.isEqual(otherTriangle.B) &&
		triangle.C.isEqual(otherTriangle.C) &&
		triangle.Alpha.isEqual(otherTriangle.Alpha) &&
		triangle.Beta.isEqual(otherTriangle.Beta) &&
		triangle.Gamma.isEqual(otherTriangle.Gamma)
}

func (a side) isEqual(b side) bool {
	floatA := big.NewFloat(float64(a))
	floatB := big.NewFloat(float64(b))
	return floatA.Cmp(floatB) == 0
}

func (a angle) isEqual(b angle) bool {
	floatA := big.NewFloat(float64(a))
	floatB := big.NewFloat(float64(b))
	return floatA.Cmp(floatB) == 0
}

type kind string
type sideClassification string
type angleClassification string

const (
	sss kind = "Side-Side-Side"
	aas kind = "Angle-Angle-Side"
	asa kind = "Angle-Side-Angle"
	sas kind = "Side-Angle-Side"
	ssa kind = "Side-Side-Angle"
)

const (
	right       sideClassification = "Right"
	equilateral sideClassification = "Equilateral"
	isosceles   sideClassification = "Isosceles"
	scalene     sideClassification = "Scalene"
)

const (
	acute  angleClassification = "Acute"
	obtuse angleClassification = "Obtuse"
)

type AbstractTriangle struct {
	A     side
	Alpha angle
	B     side
	Beta  angle
	C     side
	Gamma angle
	kind  kind
}

type Triangle interface {
	Solve()
	Description() string
}

func (triangle *AbstractTriangle) Description() string {
	var builder strings.Builder
	sideClassification, angleClassification := triangle.classification()
	header := fmt.Sprintf("%v Triangle\n", triangle.kind)
	title := fmt.Sprintf("Classification by sides: %v, by angles: %v\n", sideClassification, angleClassification)
	sides := fmt.Sprintf("Sides - A: %v, B: %v, C: %v\n", triangle.A, triangle.B, triangle.C)
	angles := fmt.Sprintf("Angles - Alpha: %v, Beta: %v, Gamma: %v\n", triangle.Alpha, triangle.Beta, triangle.Gamma)
	perimeter := fmt.Sprintf("Perimeter: %v\n", triangle.perimeter())
	semiPerimeter := fmt.Sprintf("Semiperimeter: %v\n", triangle.semiPerimeter())
	area := fmt.Sprintf("Area: %v\n", triangle.area())
	inradius := fmt.Sprintf("Inradius: %v\n", triangle.inradius())
	circumradius := fmt.Sprintf("Circumradius: %v\n", triangle.circumradius())
	builder.WriteString(header)
	builder.WriteString(title)
	builder.WriteString(sides)
	builder.WriteString(angles)
	builder.WriteString(perimeter)
	builder.WriteString(semiPerimeter)
	builder.WriteString(area)
	builder.WriteString(inradius)
	builder.WriteString(circumradius)
	return builder.String()
}

// NewTriangle returns a Triangle type based on the input received
func NewTriangle(a float64, alpha float64, b float64, beta float64, c float64, gamma float64) (Triangle, error) {
	if a != 0 && b != 0 && c != 0 {
		return newSSSTriangle(side(a), side(b), side(c)), nil
	} else if a != 0 && b != 0 && beta != 0 {
		return newSSATriangle(side(a), side(b), angle(beta)), nil
	} else if a != 0 && beta != 0 && c != 0 {
		return newSASTriangle(side(a), angle(beta), side(c)), nil
	} else if alpha != 0 && b != 0 && gamma != 0 {
		return newASATriangle(angle(alpha), side(b), angle(gamma)), nil
	} else if a != 0 && alpha != 0 && beta != 0 {
		return newAASTriangle(side(a), angle(alpha), angle(beta)), nil
	}
	return nil, errors.New("unable to understand the triangle, please format your input as SSS/SSA/SAS/ASA/AAS triangles")
}

func (triangle *AbstractTriangle) perimeter() float64 {
	return float64(triangle.A) + float64(triangle.B) + float64(triangle.C)
}

func (triangle *AbstractTriangle) semiPerimeter() float64 {
	return triangle.perimeter() / 2
}

func (triangle *AbstractTriangle) area() float64 {
	semiPerimeter := triangle.semiPerimeter()
	return math.Sqrt(semiPerimeter *
		(semiPerimeter - float64(triangle.A)) *
		(semiPerimeter - float64(triangle.B)) *
		(semiPerimeter - float64(triangle.C)))
}

func (triangle *AbstractTriangle) inradius() float64 {
	return triangle.semiPerimeter() / triangle.area()
}

func (triangle *AbstractTriangle) circumradius() float64 {
	return float64(triangle.A) / (2 * math.Sin(float64(triangle.Alpha)))
}

func (triangle *AbstractTriangle) classification() (sideClassification, angleClassification) {
	rightAngle := angle(math.Pi / 2)

	var sideClass sideClassification
	var angleClass angleClassification

	if triangle.Alpha.isEqual(rightAngle) ||
		triangle.Beta.isEqual(rightAngle) ||
		triangle.Gamma.isEqual(rightAngle) {
		sideClass = right
	} else if triangle.A == triangle.B && triangle.B == triangle.C && triangle.C == triangle.A {
		sideClass = equilateral
	} else if triangle.A == triangle.B || triangle.B == triangle.C || triangle.C == triangle.A {
		sideClass = isosceles
	} else {
		sideClass = scalene
	}

	if float64(triangle.Alpha) < float64(rightAngle) ||
		float64(triangle.Beta) < float64(rightAngle) ||
		float64(triangle.Gamma) < float64(rightAngle) {
		angleClass = acute
	} else {
		angleClass = obtuse
	}

	return sideClass, angleClass
}

func radiansToDegrees(radians float64) float64 {
	return ((radians * 180) / math.Pi)
}

func degreesToRadians(degrees float64) float64 {
	return float64((degrees * math.Pi) / 180)
}

func (triangle *AbstractTriangle) sineLawForAngle(baseSide side, baseAngle angle, knownLeg side) angle {
	a := float64(baseSide)
	alpha := float64(baseAngle)
	b := float64(knownLeg)
	return angle(math.Asin(b * math.Sin(alpha) / a))
}

func (triangle *AbstractTriangle) sineLawForSide(baseSide side, baseAngle angle, knownAngle angle) side {
	a := float64(baseSide)
	alpha := float64(baseAngle)
	beta := float64(knownAngle)
	return side(math.Sin(beta) * a / math.Sin(alpha))
}
