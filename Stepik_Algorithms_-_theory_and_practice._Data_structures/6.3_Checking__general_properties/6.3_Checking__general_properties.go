package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	In  = bufio.NewReader(os.Stdin)
	Out = bufio.NewWriter(os.Stdout)
)

type Node struct {
	key   int
	left  int
	right int
}

type Tree struct {
	t       []*Node
	prevKey *int
	correct bool
}

func CreateTree(count int) Tree {
	tree := Tree{}
	tree.t = make([]*Node, count)
	tree.correct = true
	return tree
}

func main() {
	defer Out.Flush()

	var nodeCount int
	fmt.Fscan(In, &nodeCount)

	if nodeCount == 0 || nodeCount == 1 {
		fmt.Fprint(Out, "CORRECT")
		return
	}

	tree := CreateTree(nodeCount)

	var k, l, r int
	for i := 0; i < nodeCount; i++ {

		fmt.Fscan(In, &k, &l, &r)
		node := Node{key: k, left: l, right: r}
		tree.t[i] = &node
	}

	tree.PreOrder(tree.GetNode(0))

	if tree.correct {
		fmt.Fprint(Out, "CORRECT")
		return
	}
	fmt.Fprint(Out, "INCORRECT")

}

func (t *Tree) GetNode(index int) *Node {
	if index == -1 || index >= len(t.t) {
		return nil
	}

	return t.t[index]
}

func (t *Tree) Max(node *Node) int {
	right := t.GetNode(node.right)
	if right == nil {
		return node.key
	}
	return t.Max(right)
}

func (t *Tree) Min(node *Node) int {
	left := t.GetNode(node.left)
	if left == nil {
		return node.key
	}
	return t.Min(left)
}

func (t *Tree) PreOrder(node *Node) {
	if node == nil {
		return
	}

	leftChild := t.GetNode(node.left)
	rightChild := t.GetNode(node.right)

	if leftChild != nil {

		if t.Max(leftChild) >= node.key {
			t.correct = false
			return
		}
		t.PreOrder(leftChild)
	}

	if rightChild != nil {

		if t.Min(rightChild) < node.key {

			t.correct = false
			return
		}

		t.PreOrder(rightChild)
	}
}
