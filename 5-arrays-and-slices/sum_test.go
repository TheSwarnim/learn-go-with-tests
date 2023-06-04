package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	// numbers := [5]int{1, 3, 5, 7, 9}

	// got := Sum(numbers)
	// want := 25

	// if got != want {
	// 	t.Errorf("got %d want %d given, %v", got, want, numbers)
	// }

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 3, 5, 7, 9}

		got := Sum(numbers)
		want := 25

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{2, 4, 6}

		got := Sum(numbers)
		want := 12

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v and want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{5, 6, 3})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 5})
		want := []int{0, 5}
		checkSums(t, got, want)
	})
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{2, 4, 6, 8, 10})
	}
}
