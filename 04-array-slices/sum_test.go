package array_slices

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := SumSlices(numbers)

		want := 15

		if got != want {
			t.Errorf("got %d but expected %d", got, want)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := SumSlices(numbers)

		want := 6

		if got != want {
			t.Errorf("want %d but got %d", want, got)
		}
	})
}
