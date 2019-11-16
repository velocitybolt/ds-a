package main

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func DfsInOrder(treeNode *TreeNode) {
	if treeNode == nil {
		return
	}
	DfsInOrder(treeNode.Left)
	fmt.Println(treeNode.Value)
	DfsInOrder(treeNode.Right)
}

func main() {
	tree := &TreeNode{5,
		&TreeNode{1, nil, nil},
		&TreeNode{2, &TreeNode{7, nil, nil}, nil}}

	DfsInOrder(tree)
}
