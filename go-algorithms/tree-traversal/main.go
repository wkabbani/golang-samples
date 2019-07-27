// Package main implements simple tree traversal functions
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Tree Traversal!!!")

	root := &Node{
		Value: 5,
		Left:  &Node{Value: 6, Left: &Node{Value: 9}},
		Right: &Node{Value: 7, Left: &Node{Value: 1, Right: &Node{Value: 3}}},
	}

	var result []string
	recursivePreOrder(root, func(node *Node) { result = append(result, fmt.Sprint(node)) })
	fmt.Printf("PreOrder Tree: %v\n", result)

	result = nil
	recursiveInOrder(root, func(node *Node) { result = append(result, fmt.Sprint(node)) })
	fmt.Printf("InOrder Tree: %v\n", result)

	result = nil
	recursivePostOrder(root, func(node *Node) { result = append(result, fmt.Sprint(node)) })
	fmt.Printf("PostOrder Tree: %v\n", result)

}
