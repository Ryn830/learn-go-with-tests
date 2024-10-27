package hello

import (
	"testing"
)

func TestHelloDefaul(t *testing.T) {

	assertEqual := func(t *testing.T, actual, expected string) {
		t.Helper()
		if actual != expected {
			t.Errorf("Actual: %q, Expected: %q", actual, expected)
		}
	}
	t.Run("should default to Hello, World", func(t *testing.T) {
		actual := Hello("")
		expected := "Hello, World"

		assertEqual(t, actual, expected)
	})

	t.Run("should use name when available", func(t *testing.T) {
		actual := Hello("Ryan")
		expected := "Hello, Ryan"

		assertEqual(t, actual, expected)
	})
}
