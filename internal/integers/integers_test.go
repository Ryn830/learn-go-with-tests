package integers

import (
	"testing"
)

func TestIntegers(t *testing.T) {

	assertEqual := func(t *testing.T, actual, expected int) {
		t.Helper()
		if actual != expected {
			t.Errorf("Actual: %q, Expected: %q", actual, expected)
		}
	}
	t.Run("should add numbers", func(t *testing.T) {
		actual := Add(2, 2)
		expected := 4

		assertEqual(t, actual, expected)
	})
}
