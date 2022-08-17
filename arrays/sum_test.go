package arrays

import (
	"reflect"
	"testing"
)

func checkSums(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v is given", got, want)
	}
}

func checkSum(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v is given", got, want)
	}
}

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		checkSum(t, got, want)
	})
}

func TestSumAllWithFixedCapacity(t *testing.T) {
	got := SumAllWithFixedCapacity([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	checkSums(t, got, want)
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	checkSums(t, got, want)
}

func BenchmarkSumAllWithFixedCapacity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllWithFixedCapacity([]int{1, 2}, []int{0, 9})
	}
}

// Implementation with append and slice of varying size is ~ 2 times slower
func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2}, []int{0, 9})
	}
}

func TestSumAllTails(t *testing.T) {

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})

}
