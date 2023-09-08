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
