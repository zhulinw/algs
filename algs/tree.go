package algs

import (
	"fmt"

	"github.com/gruntpig/algs/container/queue"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

// 递归实现，二叉树的所有路径
func BinaryTreePath(root *TreeNode) {
	pathList := make([]string, 0)
	treePath(root, &pathList, "")
	fmt.Println(pathList)
}

// 深度遍历
func treePath(root *TreeNode, pathList *[]string, path string) {
	if root != nil {
		path += fmt.Sprintf("%v", root.Val)
		if root.Left == nil && root.Right == nil {
			*pathList = append(*pathList, path)
		} else {
			path += fmt.Sprint("->")
			treePath(root.Left, pathList, path)
			treePath(root.Right, pathList, path)
		}
	}
}

// 广度遍历
func BinaryTreeBfsQueue(root *TreeNode) {
	queue := queue.New()
	queue.Offer(root)

	for !queue.IsEmpty() {
		treeNode := queue.Poll().(*TreeNode)
		fmt.Println(treeNode.Val)

		if treeNode.Left != nil {
			queue.Offer(treeNode.Left)
		}
		if treeNode.Right != nil {
			queue.Offer(treeNode.Right)
		}
	}
}

// 广度遍历，Slice模拟队列操作
func BinaryTreeBfsSlice(root *TreeNode) {
	nodeQueue := make([]*TreeNode, 0, 0)

	nodeQueue = append(nodeQueue, root)
	for i := 0; i < len(nodeQueue); i++ {
		node:= nodeQueue[i]
		fmt.Println(node.Val)
		if node.Left == nil && node.Right == nil {
			continue
		}

		if node.Left != nil {
			nodeQueue = append(nodeQueue, node.Left)
		}

		if node.Right != nil {
			nodeQueue = append(nodeQueue, node.Right)
		}
	}
}

// 广度遍历，获取二叉树所有的路径
func BinaryTreeBfsPath(root *TreeNode) {
	paths := make([]string, 0, 0)

	nodeQueue := make([]*TreeNode, 0, 0)
	pathQueue := make([]string, 0, 0)

	nodeQueue = append(nodeQueue, root)
	pathQueue = append(pathQueue, fmt.Sprint(root.Val))
	for i := 0; i < len(nodeQueue); i++ {
		node, path := nodeQueue[i], pathQueue[i]
		if node.Left == nil && node.Right == nil {
			paths = append(paths, path)
			continue
		}

		if node.Left != nil {
			nodeQueue = append(nodeQueue, node.Left)
			pathQueue = append(pathQueue, path+fmt.Sprintf("->%v", node.Left.Val))
		}

		if node.Right != nil {
			nodeQueue = append(nodeQueue, node.Right)
			pathQueue = append(pathQueue, path+fmt.Sprintf("->%v", node.Right.Val))
		}
	}
	fmt.Println(paths)
}
