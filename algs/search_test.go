package algs

import (
	"math"
	"testing"

	"github.com/gruntpig/algs/utils"
)

func TestBinarySearch(t *testing.T) {
	methods := []func(s []int, ele int) int{
		BinarySearch, RecurBinarySearch}

	for _, method := range methods {
		// input data
		s1 := []int{1, 3, 6, 7, 9, 13, 18}
		i1 := 3
		i2 := 13
		i3 := 7
		i4 := 5
		// test func
		got1 := method(s1, i1)
		utils.Equal(t, i1, got1)
		got2 := method(s1, i2)
		utils.Equal(t, i2, got2)
		got3 := method(s1, i3)
		utils.Equal(t, i3, got3)
		got4 := method(s1, i4)
		utils.Equal(t, math.MaxInt32, got4)

		// input data
		s1 = []int{2, 5, 9, 22, 26, 34}
		i1 = 2
		i2 = 9
		i3 = 34
		i4 = 100
		// test func
		got1 = method(s1, i1)
		utils.Equal(t, i1, got1)
		got2 = method(s1, i2)
		utils.Equal(t, i2, got2)
		got3 = method(s1, i3)
		utils.Equal(t, i3, got3)
		got4 = method(s1, i4)
		utils.Equal(t, math.MaxInt32, got4)
	}
}
