package algs

import (
	"fmt"
	"testing"
)

// 			5
//        /	  \
//       4     8
//		/ 	 /   \
//     11    13	  4
// 	  /	 \		   \
// 	 7	  2			1
func treeNode1() *TreeNode {
	tl11 := NewTreeNode(11)
	tl11.Left = NewTreeNode(7)
	tl11.Right = NewTreeNode(2)

	tl4 := NewTreeNode(4)
	tl4.Left = tl11

	tr4 := NewTreeNode(4)
	tr4.Right = NewTreeNode(1)
	tr8 := NewTreeNode(8)
	tr8.Left = NewTreeNode(13)
	tr8.Right = tr4

	tr5 := NewTreeNode(5)
	tr5.Left = tl4
	tr5.Right = tr8

	return tr5
}

func treeNode2() *TreeNode {
	tn5 := NewTreeNode(5)
	tn2 := NewTreeNode(2)
	tn2.Right = tn5

	tn1 := NewTreeNode(1)
	tn1.Left = tn2
	tn1.Right = NewTreeNode(3)
	return tn1
}

func TestBinaryTreePath(t *testing.T) {
	root := treeNode2()
	BinaryTreePath(root)

	root = treeNode1()
	BinaryTreePath(root)
}

func TestBinaryTreeDfs(t *testing.T) {
	root := treeNode2()
	BinaryTreeBfsQueue(root)

	fmt.Println()
	root = treeNode1()
	BinaryTreeBfsQueue(root)
}

func TestBinaryTreeBfs(t *testing.T) {
	root := treeNode2()
	BinaryTreeBfsSlice(root)

	fmt.Println()
	root = treeNode1()
	BinaryTreeBfsSlice(root)
}

func TestBinaryTreeDfsSlice(t *testing.T) {
	root := treeNode2()
	BinaryTreeBfsPath(root)

	fmt.Println()
	root = treeNode1()
	BinaryTreeBfsPath(root)
}
