package structs

import (
	"testing"
)

type Shape interface {
	Area() float64
}

func checkArea(t *testing.T, shape Shape, expected float64) {
	t.Helper()
	actual := shape.Area()

	if actual != expected {
		t.Errorf("%#v: Actual %g not equal %g", shape, actual, expected)
	}
}

func TestPerimeter(t *testing.T) {
	t.Run("Calculates the perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		actual := rectangle.Perimeter()
		expected := 40.0

		if actual != expected {
			t.Errorf("Actual %.2f not equal %.2f", actual, expected)
		}
	})
}

func TestArea(t *testing.T) {
	tests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			checkArea(t, test.shape, test.expected)
		})
	}
}
