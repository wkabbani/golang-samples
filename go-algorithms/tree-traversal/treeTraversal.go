package main

import (
	"fmt"
	"strings"
)

// Node represents a node in a tree data structure
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Walker represents the type of the function that acts on the node when visited
type Walker func(*Node)

func (node *Node) String() string {
	var str strings.Builder
	fmt.Fprintf(&str, "%d ", node.Value)
	return str.String()
}

// Left -> Node -> Right
func recursiveInOrder(node *Node, fn Walker) {
	if node == nil {
		return
	}
	recursiveInOrder(node.Left, fn)
	fn(node)
	recursiveInOrder(node.Right, fn)
}

// Node -> Left -> Right
func recursivePreOrder(node *Node, fn Walker) {
	if node == nil {
		return
	}
	fn(node)
	recursivePreOrder(node.Left, fn)
	recursivePreOrder(node.Right, fn)
}

// Left -> Right -> Node
func recursivePostOrder(node *Node, fn Walker) {
	if node == nil {
		return
	}
	recursivePostOrder(node.Left, fn)
	recursivePostOrder(node.Right, fn)
	fn(node)
}
