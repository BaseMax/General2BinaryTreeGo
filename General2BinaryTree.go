/**
 *
 * @Name: General Tree To Binary Tree
 * @Author: Max Base
 * @Date: 2022/12/17
 * @Repository: https://github.com/BaseMax/General2BinaryTree
 * @License: GPL-3
 *
 */

package main

import "fmt"

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

func printBinaryTree(root *BinaryNode, level int) {
	if root == nil {
		return
	}
	for i := 0; i < level; i++ {
		fmt.Print(" ")
	}
	fmt.Println(root.data)
	printBinaryTree(root.left, level+1)
	printBinaryTree(root.right, level+1)
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
