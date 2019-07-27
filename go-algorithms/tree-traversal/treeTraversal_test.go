package main

import "testing"

func TestRecursivePreOrder(t *testing.T) {

	tree := &Node{
		Value: 5,
		Left:  &Node{Value: 6, Left: &Node{Value: 9}},
		Right: &Node{Value: 7, Left: &Node{Value: 1, Right: &Node{Value: 3}}},
	}

	var actual []int
	walker := func(node *Node) {
		actual = append(actual, node.Value)
	}

	expected := [...]int{5, 6, 9, 7, 1, 3}

	recursivePreOrder(tree, walker)

	var actualArrary [len(expected)]int
	copy(actualArrary[:], actual[:len(expected)])

	if actualArrary != expected {
		t.Errorf("RecursivePreOrder failed, got: %v, want: %v.", actualArrary, expected)
	}
}

func TestRecursiveInOrder(t *testing.T) {

	tree := &Node{
		Value: 5,
		Left:  &Node{Value: 6, Left: &Node{Value: 9}},
		Right: &Node{Value: 7, Left: &Node{Value: 1, Right: &Node{Value: 3}}},
	}

	var actual []int
	walker := func(node *Node) {
		actual = append(actual, node.Value)
	}

	expected := [...]int{9, 6, 5, 1, 3, 7}

	recursiveInOrder(tree, walker)

	var actualArrary [len(expected)]int
	copy(actualArrary[:], actual[:len(expected)])

	if actualArrary != expected {
		t.Errorf("RecursiveInOrder failed, got: %v, want: %v.", actualArrary, expected)
	}
}

func TestRecursivePostOrder(t *testing.T) {

	tree := &Node{
		Value: 5,
		Left:  &Node{Value: 6, Left: &Node{Value: 9}},
		Right: &Node{Value: 7, Left: &Node{Value: 1, Right: &Node{Value: 3}}},
	}

	var actual []int
	walker := func(node *Node) {
		actual = append(actual, node.Value)
	}

	expected := [...]int{9, 6, 3, 1, 7, 5}

	recursivePostOrder(tree, walker)

	var actualArrary [len(expected)]int
	copy(actualArrary[:], actual[:len(expected)])

	if actualArrary != expected {
		t.Errorf("RecursivePostOrder failed, got: %v, want: %v.", actualArrary, expected)
	}
}
