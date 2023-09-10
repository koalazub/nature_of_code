package vectors

import (
	"testing"
)

func TestSub(t *testing.T) {
	got := Vector{3., 3., 3.}.Sub(Vector{1., 1., 1.})
	expected := Vector{2., 2., 2.}

	if got != expected {
		t.Fatalf("Expected: %v, got: %v", expected, got)
	}
}

func TestMultiplyByScalar(t *testing.T) {
	got := Vector{3., 3., 3.}.MultiplyByScalar(3.)
	expected := Vector{9., 9., 9.}
	if got != expected {
		t.Fatalf("Expected: %v, got: %v", expected, got)
	}
}

func TestDotProductOfTwoPerpendicularVectors(t *testing.T) {
	got := Vector{1., 0., 0.}.Dot(Vector{0., 1., 0.})
	expected := 0.

	if got != expected {
		t.Fatalf("Expected: %v, Got: %v", expected, got)
	}
}
func TestDotProductOfTwoParallelVectors(t *testing.T) {
	got := Vector{1., 0., 0.}.Dot(Vector{1., 0., 0.})
	expected := 1.

	if got != expected {
		t.Fatalf("Expected: %v, Got: %v", expected, got)
	}
}

// Calculate the length(magnitude) for Vector
func TestDotLengthOfTwoVectors(t *testing.T) {
	got := Vector{6., 6., 6.}.Length()
	expected := 10.392304845413264
	if got != expected {
		t.Fatalf("Got: %v, expected: %v", got, expected)
	}

}

// Cross return perpendicular vector to calculated values i.e. x * y = z
func TestCrossProduct(t *testing.T) {
	got := Vector{3., 1., 1.}.Cross(Vector{1., 1., 3.})
	expected := Vector{2., -8., 2.}
	if got != expected {
		t.Fatalf("Expected: %v, got: %v", expected, got)
	}
}

func TestLength(t *testing.T) {
	got := Vector{1., 0., 0.}.Normalise()
	expected := Vector{1., 0., 0.}

	if got != expected {
		t.Fatalf("Expected: %v, got: %v", expected, got)
	}

	got = Vector{0., 1., 0}.Normalise()
	expected = Vector{0., 1., 0.}

	if got != expected {
		t.Fatalf("Expected: %v, got: %v", expected, got)
	}

	got = Vector{0., 0., 1}.Normalise()
	expected = Vector{0., 0., 1.}

	if got != expected {
		t.Fatalf("Expected: %v, got: %v", expected, got)
	}

}
