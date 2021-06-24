package list

import (
	"fmt"
	"strings"
)

// List with head
// head save list len and last ele
type List struct {
	next  *List
	last  *List
	value interface{}
	num   int
}

func New() *List {
	return Init()
}

func Init() *List {
	return &List{}
}

func (head *List) String() string {
	var buf strings.Builder
	buf.WriteString("[")
	for node := head.next; node != nil; node = node.next {
		buf.WriteString(fmt.Sprintf("%v, ", node.value))
	}
	buf.WriteString("]")
	return buf.String()
}

func (head *List) Next() *List {
	if head == nil {
		return nil
	}
	return head.next
}

// insert first, return the insert ele
func (head *List) InsertFirst(value interface{}) *List {
	if head == nil {
		return nil
	}
	head.num++

	oldFirst := head.next
	newFirst := New()
	newFirst.value = value
	newFirst.next = oldFirst

	head.next = newFirst
	return newFirst
}

// insert last, return the insert ele
func (head *List) InsertLast(value interface{}) *List {
	if head == nil {
		return nil
	}
	head.num++

	oldLast := head.last

	newLast := New()
	newLast.value = value

	if oldLast == nil {
		oldLast = newLast
	} else {
		oldLast.next = newLast
	}

	head.last = newLast
	return newLast
}
