package stack

import (
	"testing"

	"github.com/gruntpig/algs/utils"
)

func TestStack_Pop(t *testing.T) {
	stack := NewStack()
	input := 3
	input1 := 4
	stack.Pop(input)
	stack.Pop(input1)
	got := stack.Push()
	got1 := stack.Push()
	utils.Equal(t, input1, got)
	utils.Equal(t, input, got1)
}

func TestStack_Push(t *testing.T) {
	stack := NewStack()
	input := 3
	stack.Pop(input)
	stack.Push()

	if !stack.IsEmpty() {
		t.Fatal("struct is not empty")
	}
}

func TestStack_Size(t *testing.T) {
	stack := NewStack()

	stack.Pop(1)
	stack.Pop(2)
	stack.Pop(3)

	want1 := 3
	got1 := stack.Size()
	utils.Equal(t, want1, got1)

	stack.Push()
	want2 := 2
	got2 := stack.Size()
	utils.Equal(t, want2, got2)
}
