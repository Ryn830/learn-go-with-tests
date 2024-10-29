package mocking

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	t.Run("Countdown", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeperSpy := &SpySleeper{}

		Countdown(buffer, sleeperSpy)

		actual := buffer.String()
		expected := `3
2
1
Go!`

		if actual != expected {
			t.Errorf("Actual %q not equal %q", actual, expected)
		}

		if sleeperSpy.Calls != 4 {
			t.Errorf("Expected sleer.Sleep to be called 4 times, called %d", sleeperSpy.Calls)
		}
	})
}
