package triangle

import (
	"math"
	"testing"
)

func TestRadiansToDegrees(t *testing.T) {
	radians := math.Pi
	expected := float64(180)
	actual := radiansToDegrees(radians)

	if expected != actual {
		t.Error("Expected 180 degrees, got ", actual)
	}
}

func TestDegreesToRadians(t *testing.T) {
	degrees := float64(180)
	expected := math.Pi
	actual := degreesToRadians(degrees)

	if expected != actual {
		t.Error("Expected pi radians, got ", actual)
	}
}

func TestSolveSSSTriangle(t *testing.T) {
	actual := newSSSTriangle(side(1), side(2), side(2))
	actual.Solve()

	expected := AbstractTriangle{
		A:     side(1),
		Alpha: angle(0.5053605102841574),
		B:     side(2),
		Beta:  angle(1.318116071652818),
		C:     side(2),
		Gamma: angle(1.318116071652818),
	}

	if !actual.isEqual(expected) {
		t.Errorf("Fail at solving triangle, expected %v actual %v", expected, actual)
	}
}

func TestSolveSASTriangle(t *testing.T) {
	alpha := degreesToRadians(30)
	actual := newSASTriangle(side(2), angle(alpha), side(1))
	actual.Solve()

	expected := AbstractTriangle{
		A:     side(2),
		Alpha: angle(0.9388820144198251),
		B:     side(1.2393136749274758),
		Beta:  angle(alpha),
		C:     side(1),
		Gamma: angle(0.41528323882152635),
	}

	if !actual.isEqual(expected) {
		t.Errorf("Fail at solving triangle, expected %v actual %v", expected, actual)
	}
}

func TestSolveSSATriangle(t *testing.T) {
	beta := degreesToRadians(45)
	actual := newSSATriangle(side(1), side(2), angle(beta))
	actual.Solve()

	expected := AbstractTriangle{
		A:     side(1),
		Alpha: angle(0.3613671239067078),
		B:     side(2),
		Beta:  angle(beta),
		C:     side(2.5779354745735183),
		Gamma: angle(1.9948273662856373),
	}

	if !actual.isEqual(expected) {
		t.Errorf("Fail at solving triangle, expected %v actual %v", expected, actual)
	}
}

func TestSolveASATriangle(t *testing.T) {
	alpha := degreesToRadians(30)
	gamma := degreesToRadians(60)
	actual := newASATriangle(angle(alpha), side(2), angle(gamma))
	actual.Solve()

	expected := AbstractTriangle{
		A:     side(0.9999999999999999),
		Alpha: angle(alpha),
		B:     side(2),
		Beta:  angle(1.5707963267948968),
		C:     side(1.7320508075688772),
		Gamma: angle(gamma),
	}

	if !actual.isEqual(expected) {
		t.Errorf("Fail at solving triangle, expected %v actual %v", expected, actual)
	}
}

func TestSolveAASTriangle(t *testing.T) {
	alpha := degreesToRadians(30)
	beta := degreesToRadians(45)
	actual := newAASTriangle(side(2), angle(alpha), angle(beta))
	actual.Solve()

	expected := AbstractTriangle{
		A:     side(2),
		Alpha: angle(alpha),
		B:     side(2.8284271247461903),
		Beta:  angle(beta),
		C:     side(3.8637033051562737),
		Gamma: angle(1.8325957145940461),
	}

	if !actual.isEqual(expected) {
		t.Errorf("Fail at solving triangle, expected %v actual %v", expected, actual)
	}
}
