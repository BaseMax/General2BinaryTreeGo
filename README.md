# General-Tree to Binary-Tree Go

This is a simple Go program that converts a general tree to a binary tree. The general tree is represented as a map of nodes, where each node has a key and a list of children. The binary tree is represented as a map of nodes, where each node has a key and two children. The binary tree is constructed by taking the first child of each node in the general tree and making it the left child of the corresponding node in the binary tree. The second child of each node in the general tree is made the right child of the corresponding node in the binary tree. If a node in the general tree has only one child, then the corresponding node in the binary tree will have only a left child.

## Usage

