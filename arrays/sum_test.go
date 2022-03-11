package arrays

import (
	"reflect"
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

func TestSumAll(t *testing.T) {
	t.Run("sum of 2 slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v but expected %v", got, expected)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v but expected %v", got, expected)
		}
	}

	t.Run("sum of tails of 2 slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 9}, []int{2, 3, 4})
		expected := []int{5, 18, 7}
		checkSums(t, got, expected)
	})

	t.Run("sum of tails of an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3})
		expected := []int{0, 5}
		checkSums(t, got, expected)
	})
}
