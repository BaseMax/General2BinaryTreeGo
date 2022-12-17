package main

func main() {
	// Create a general tree
	generalTree := &GeneralTree{
		root: &Node{
			data: 1,
			child: []*Node{
				&Node{
					data: 2,
					child: []*Node{
						&Node{
							data: 4,
							child: []*Node{
								&Node{data: 8},
								&Node{data: 9},
							},
						},
						&Node{data: 5},
					},
				},
				&Node{
					data: 3,
					child: []*Node{
						&Node{data: 6},
						&Node{data: 7},
					},
				},
			},
		},
	}

	// Convert general tree to binary tree
	binaryTree := generalTree.General2BinaryTree()

	// Print the binary tree
	printBinaryTree(binaryTree.root, 0)
}
