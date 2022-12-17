# General-Tree to Binary-Tree Go

This is a simple Go program that converts a general tree to a binary tree.

### General Tree

The general tree is represented as a map of nodes, where each node has a key and a list of children. The binary tree is represented as a map of nodes, where each node has a key and two children.

### Binary Tree

The binary tree is constructed by taking the first child of each node in the general tree and making it the left child of the corresponding node in the binary tree. The second child of each node in the general tree is made the right child of the corresponding node in the binary tree. If a node in the general tree has only one child, then the corresponding node in the binary tree will have only a left child.

## Usage

```go
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
```

## License

This project is licensed under the GPL-3.0 License - see the [LICENSE](LICENSE) file for details.

Copyright (c) 2022, Max Base
