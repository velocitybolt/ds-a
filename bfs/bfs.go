package main

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func bfsInOrder(treeNode *TreeNode) {
	visited := make(map[*TreeNode]struct{})
	currentNode := treeNode
	if val, ok := visited
}

func main() {
	treeNode := &TreeNode{3, // root of the tree
		&TreeNode{2, &TreeNode{5, nil, nil}, nil}}
	bfsInOrder(treeNode)
}
