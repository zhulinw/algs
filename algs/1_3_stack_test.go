package algs

import (
	"testing"

	"github.com/gruntpig/algs/utils"
)

func TestEvaluate(t *testing.T) {
	e := "(1+((2+3)*(4*5)))"
	want := 101
	got := Evaluate(e)
	utils.Equal(t, want, got)
}
