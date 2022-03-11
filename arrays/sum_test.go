package arrays

import (
	"testing"
)

func TestSum(t *testing.T) {
	assertCorrectSum := func(t testing.TB, numbers []int, got, expected int) {
		t.Helper()
		if got != expected {
			t.Errorf("got '%d' but expected '%d' for sum of %v", got, expected, numbers)
		}
	}

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6
		assertCorrectSum(t, numbers, got, expected)
	})
}
