package algs

import (
	"testing"

	"github.com/gruntpig/algs/utils"
)

func TestReverseSlice(t *testing.T) {
	input := []int{1, 3, 6, 7, 9, 13, 18}
	want := []int{18, 13, 9, 7, 6, 3, 1}
	got := ReverseSlice(input)
	utils.Equal(t, want, got)
}
