package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("Greet name given", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Chris")

		actual := buffer.String()
		expected := "Hello, Chris"

		if actual != expected {
			t.Errorf("Actual %q doesn't equal %q", actual, expected)
		}
	})
}
