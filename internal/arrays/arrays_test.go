package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Sums integers together", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5}

		actual := Sum(numbers)
		expected := 15

		if actual != expected {
			t.Errorf("Actual: %q, not equal %q", actual, expected)
		}
	})

	t.Run("Works with collections of different sizes", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		actual := Sum(numbers)
		expected := 6

		if actual != expected {
			t.Errorf("Actual: %q, not equal %q", actual, expected)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Reduces an array of arrays of integers into a single array", func(t *testing.T) {
		actual := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Actual: %q, not equal %q", actual, expected)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("Sums all the tail elements of all given int arrays", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Actual: %q, not equal %q", actual, expected)
		}
	})

	t.Run("Safely sums empty slices", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{0, 2, 3})
		expected := []int{0, 5}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Actual: %q, not equal %q", actual, expected)
		}
	})
}
