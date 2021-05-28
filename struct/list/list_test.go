package list

import (
	"testing"

	"github.com/gruntpig/algs/utils"
)

func TestList_InsertFirst(t *testing.T) {
	head := New()
	head.InsertFirst("go")
	head.InsertFirst(1)
	head.InsertFirst("java")

	checkListElePoint(t, head, "java", 1, "go")

	head.InsertFirst(5)
	checkListElePoint(t, head, 5, "java", 1, "go")

}

func TestList_InsertLast(t *testing.T) {
	head := New()
	head.InsertLast("go")
	head.InsertLast(1)
	head.InsertLast("java")

	checkListElePoint(t, head, "go", 1, "java")

	head.InsertLast("5")
	checkListElePoint(t, head, "go", 1, "java", 5)
}

func TestList_InsertFirstLast(t *testing.T) {
	head := New()
	head.InsertFirst("go")
	head.InsertFirst(1)
	head.InsertLast("java")

	checkListElePoint(t, head, 1, "go", "java")

	head.InsertLast("c++")
	checkListElePoint(t, head, 1, "go", "java", "c++")
}

func checkListElePoint(t *testing.T, head *List, eles ...interface{}) {
	result := true
	index := 0
	for node := head.next; node != nil; node = node.next {
		if !utils.AssertEqual(node.value, eles[index]) {
			result = false
		}
		index++
	}
	if !result {
		t.Fatalf("not equal, got: %v, but want: %v", head, eles)
	}
}
