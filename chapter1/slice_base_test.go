package chapter1

import (
	"github.com/gruntpig/algs/utils"
	"testing"
)

func TestReverseSlice(t *testing.T) {
	input := []int{1, 3, 6, 7, 9, 13, 18}
	want := []int{18, 13, 9, 7, 6, 3, 1}
	got := ReverseSlice(input)
	utils.Equal(t, want, got)
}

func TestBinarySearchOdd(t *testing.T) {
	// test data
	s1 := []int{1, 3, 6, 7, 9, 13, 18}
	i1 := 3
	i2 := 13
	i3 := 7
	// test utils
	got1 := BinarySearch(s1, i1)
	utils.Equal(t, i1, got1)
	got2 := BinarySearch(s1, i2)
	utils.Equal(t, i2, got2)
	got3 := BinarySearch(s1, i3)
	utils.Equal(t, i3, got3)
}

func TestBinarySearchEven(t *testing.T) {
	// test data
	s2 := []int{2, 5, 9, 22, 26, 34}
	i1 := 2
	i2 := 9
	i3 := 22
	// test utils
	got1 := BinarySearch(s2, i1)
	utils.Equal(t, i1, got1)
	got2 := BinarySearch(s2, i2)
	utils.Equal(t, i2, got2)
	got3 := BinarySearch(s2, i3)
	utils.Equal(t, i3, got3)
}

func TestRecurBinarySearch(t *testing.T) {
	// test data
	s2 := []int{2, 5, 9, 22, 26, 34}
	i1 := 2
	i2 := 9
	// test utils
	got1 := RecurBinarySearch(s2, 0, len(s2)-1, i1)
	utils.Equal(t, i1, got1)

	got2 := RecurBinarySearch(s2, 0, len(s2)-1, i2)
	utils.Equal(t, i2, got2)
}