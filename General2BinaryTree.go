package main

type BinaryNode struct {
	data  interface{}
	left  *BinaryNode
	right *BinaryNode
}

type Node struct {
	data  interface{}
	child []*Node
}

type GeneralTree struct {
	root *Node
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *GeneralTree) General2BinaryTree() *BinaryTree {
	return &BinaryTree{root: general2BinaryTree(t.root)}
}

func printBinaryTree(root *BinaryNode) {
	if root == nil {
		return
	}
	printBinaryTree(root.left)
	printBinaryTree(root.right)
}

func general2BinaryTree(root *Node) *BinaryNode {
	if root == nil {
		return nil
	}
	binaryRoot := &BinaryNode{data: root.data}
	if len(root.child) > 0 {
		binaryRoot.left = general2BinaryTree(root.child[0])
	}
	if len(root.child) > 1 {
		binaryRoot.right = general2BinaryTree(root.child[1])
	}
	return binaryRoot
}
