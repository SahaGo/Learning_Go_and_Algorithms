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
	t []Node
}

func CreateTree(count int) Tree {
	tree := Tree{}
	tree.t = make([]Node, count)
	return tree
}

func main() {
	defer Out.Flush()

	var nodeCount int
	fmt.Fscan(In, &nodeCount)

	if nodeCount == 0 {
		fmt.Fprint(Out, "CORRECT")
		return
	}

	tree := CreateTree(nodeCount)

	var k, l, r int
	for i := 0; i < nodeCount; i++ {

		fmt.Fscan(In, &k, &l, &r)
		node := Node{key: k, left: l, right: r}
		tree.t[i] = node
	}

	if tree.IsValid(0, nodeCount) {
		fmt.Fprint(Out, "CORRECT")
		return
	}
	fmt.Fprint(Out, "INCORRECT")
}

func (t *Tree) IsValid(node, capa int) bool {
	list := make([]int, 0, capa)
	t.Inorder(node, &list)
	return t.IsSorted(list)
}

func (t *Tree) Inorder(node int, list *[]int) {
	if node != -1 {
		t.Inorder(t.t[node].left, list)
		*list = append(*list, t.t[node].key)
		t.Inorder(t.t[node].right, list)
		return
	}
	return
}

func (t *Tree) IsSorted(list []int) bool {

	for i := 1; i < len(list); i++ {
		if list[i] <= list[i-1] {
			return false
		}
	}
	return true
}
