package arrayAndSlice

import (
	"reflect"
	"testing"
)

func assertCorrectMessage(t *testing.T, got, want int, nubs []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, nubs)
	}
}

func assertCorrectInteger(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbs := []int{1, 2, 3, 4, 5}

		got := Sum(numbs)
		want := 15
		assertCorrectMessage(t, got, want, numbs)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbs := []int{1, 2, 3}
		got := Sum(numbs)
		want := 6
		assertCorrectMessage(t, got, want, numbs)
	})
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{})
	}
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	assertCorrectInteger(t, got, want)
}

func TestSumAllTails(t *testing.T) {

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		assertCorrectInteger(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		assertCorrectInteger(t, got, want)
	})

}
