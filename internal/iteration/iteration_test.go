package iteration

import (
	"testing"
)

func TestIteration(t *testing.T) {
	actual := Repeat("a", 5)
	expected := "aaaaa"

	if actual != expected {
		t.Errorf("Actual %q, not equal %q", actual, expected)
	}
}

func BenchmarkIteration(b *testing.B) {
	for range b.N {
		Repeat("a", 5)
	}
}
