package algs

import (
	"strconv"

	"github.com/gruntpig/algs/container/stack"
)

func Evaluate(s string) int {
	ops := stack.NewStack()
	vals := stack.NewStack()

	for _, ele := range s {
		if ele == '(' {
		} else if isOp(ele) {
			ops.Pop(ele)
		} else if ele == ')' {
			val := vals.Push().(int)
			op := ops.Push()
			if op == '+' {
				e := vals.Push().(int)
				val = e + val
			}
			if op == '-' {
				e := vals.Push().(int)
				val = e - val
			}
			if op == '*' {
				e := vals.Push().(int)
				val = e * val
			}
			if op == '/' {
				e := vals.Push().(int)
				val = e / val
			}
			vals.Pop(val)
		} else {
			val, _ := strconv.Atoi(string(ele))
			vals.Pop(val)
		}
	}
	v := vals.Push().(int)
	return v
}

func isOp(v int32) bool {
	switch v {
	case '+', '-', '*', '/':
		return true
	}
	return false
}
