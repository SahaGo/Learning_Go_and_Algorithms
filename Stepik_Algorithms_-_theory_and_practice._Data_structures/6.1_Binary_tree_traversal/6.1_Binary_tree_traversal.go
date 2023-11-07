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

	tree := CreateTree(nodeCount)

	var k, l, r int
	for i := 0; i < nodeCount; i++ {

		fmt.Fscan(In, &k, &l, &r)
		node := Node{key: k, left: l, right: r}
		tree.t[i] = node
	}

	tree.InOrder(0)
	fmt.Fprint(Out, "\n")
	tree.PreOrder(0)
	fmt.Fprint(Out, "\n")
	tree.PostOrder(0)
}

func (t *Tree) InOrder(node int) {
	if node == -1 {
		return
	}
	t.InOrder(t.t[node].left)
	fmt.Fprint(Out, t.t[node].key, " ")
	t.InOrder(t.t[node].right)
}

func (t *Tree) PreOrder(node int) {
	if node == -1 {
		return
	}
	fmt.Fprint(Out, t.t[node].key, " ")
	t.PreOrder(t.t[node].left)
	t.PreOrder(t.t[node].right)
}

func (t *Tree) PostOrder(node int) {
	if node == -1 {
		return
	}
	t.PostOrder(t.t[node].left)
	t.PostOrder(t.t[node].right)
	fmt.Fprint(Out, t.t[node].key, " ")
}
