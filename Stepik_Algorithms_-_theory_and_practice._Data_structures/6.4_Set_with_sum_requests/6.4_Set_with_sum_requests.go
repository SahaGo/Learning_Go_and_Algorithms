package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	Scan = bufio.NewScanner(os.Stdin)
	Out  = bufio.NewWriter(os.Stdout)
)

type Node struct {
	key    int
	parent *Node
	left   *Node
	right  *Node
	sum    int
}

type Tree struct {
	root *Node
	prev *int
	sum  int
}

func main() {

	Scan.Scan()
	reqCount, _ := strconv.Atoi(Scan.Text())

	tree := &Tree{}

	for i := 0; i < reqCount; i++ {
		Scan.Scan()
		req := strings.Split(Scan.Text(), " ")

		key, _ := strconv.Atoi(req[1])
		var key2 int
		if len(req) > 2 {
			key2, _ = strconv.Atoi(req[2])
		}

		switch req[0] {
		case "+": // add num
			tree.CheckForInsert(GetInd(key, tree.sum))

		case "-": // delete num
			tree.Remove(GetInd(key, tree.sum))
		case "?": // search num in tree
			if tree.SearchInRoot(GetInd(key, tree.sum)) {
				fmt.Fprintln(Out, "Found")
			} else {
				fmt.Fprintln(Out, "Not found")
			}

		case "s": // sum nums in list
			fmt.Fprintln(Out, tree.Sum(GetInd(key, tree.sum), GetInd(key2, tree.sum)))
		}
	}
	defer Out.Flush()
}

func GetInd(key, sum int) int {
	return (key + sum) % 1000000001
}

func (t *Tree) CheckForInsert(key int) {
	if t.root == nil {
		t.root = &Node{key: key}
		return
	}

	t.Insert(key, t.root)
}

func (t *Tree) Insert(key int, node *Node) {
	if node.key == key { // если ключ есть вершина, ничего не делаем
		return
	}

	if key < node.key { // если ключ меньше ключа корня, идем влево

		if node.left == nil { //  левый сын не существует, создаем левого сына
			node.left = &Node{key: key, parent: node, sum: key}
			t.Splay(node.left) // вытягиваем сына на вершину дерева
			return
		}

		t.Insert(key, node.left) // если есть, то вызовем Add, но в качестве корня передаем существующего левого сына
		return

	} else { // для ключей, больше ключа корня, идем направо

		if node.right == nil { // правый сын не существует, создаем правого сына
			node.right = &Node{key: key, parent: node, sum: key}
			t.Splay(node.right) // вытягиваем сына на вершину дерева
			return
		}

		t.Insert(key, node.right) // если есть, то передаем его рекурсивно как корень
	}
}

func (t *Tree) Splay(u *Node) {
	if u == nil { // проверяем не пустая ли вершина
		return
	}

	// проверяем есть ли родитель у вершины
	if u.parent == nil {
		return
	}

	//проверяем есть ли прародитель у вершины(родитель родителя)
	if u.parent.parent == nil { // прародителя не существует, то

		// если вершина == левый сын, выполняем Правый разворот Zig
		if u.parent.left == u {
			t.Zig(u)
			return
		}
		// если вершина == правый сын, то выполняем Левый разворот Zag
		t.Zag(u) // ес
		return
	}

	// проверяем вершина == левый сын и его родитель == левый сын,
	// выполняем два правых разворота ZigZig
	if u.parent.left == u && u.parent.parent.left == u.parent {
		t.ZigZig(u)
		t.Splay(u)
		return
	}

	// проверяем вершина == левый сын и его родитель == правый сын,
	// выполняем правый, затем левый разворот ZigZag
	if u.parent.left == u && u.parent.parent.right == u.parent {
		t.ZigZag(u)
		t.Splay(u)
		return
	}

	// проверяем вершина == правый сын и его родитель == правый сын,
	// выполняем два левых разворота ZagZag
	if u.parent.right == u && u.parent.parent.right == u.parent {
		t.ZagZag(u)
		t.Splay(u)
		return
	}

	// проверяем вершина == правый сын и его родитель == левый сын,
	// выполняем левый, затем правый разворот ZigZag
	if u.parent.right == u && u.parent.parent.left == u.parent {
		t.ZagZig(u)
		t.Splay(u)
		return
	}

	panic("panic")
}

// LinkingWithParent меняем связи родителя и прародителя
func (t *Tree) LinkingWithParent(currParent *Node, gParent *Node) {
	if gParent.parent == nil { // если у прародителя нет родителя, то

		currParent.parent = nil // затираем родителя у текущего родителя
		t.root = currParent     // делаем родителя - корнем

	} else { // если у прародителя есть родитель, то

		if gParent.parent.left == gParent { //если прародитель == левый сын
			gParent.parent.left = currParent   // делаем родителя левым сыном
			currParent.parent = gParent.parent // а родителя прародителя - прародителем родителя
		} else { // аналогично для правого сына
			gParent.parent.right = currParent
			currParent.parent = gParent.parent
		}
	}
}

// Zig Правый разворот
func (t *Tree) Zig(u *Node) {
	parent := u.parent          // берем родителя вершины Ю
	t.SmallRightTurn(u, parent) // делаем правый разворот
	u.parent = nil              // убираем родителя вершины
	t.root = u                  // теперь вершина - корень
}

// Zag Левый разворот
func (t *Tree) Zag(u *Node) {
	parent := u.parent         // берем родителя вершины Ю
	t.SmallLeftTurn(u, parent) // делаем левый разворот
	u.parent = nil             // убираем родителя вершины
	t.root = u                 // теперь вершина - корень
}

// ZigZig два Правых разворота
func (t *Tree) ZigZig(u *Node) {
	parent := u.parent           // родитель вершины Ю
	grandParent := parent.parent // прародитель вершины Ю

	t.LinkingWithParent(u, grandParent)   // меняем указатели
	t.SmallRightTurn(parent, grandParent) // сначала разворачиваем родителя и прародителя
	t.SmallRightTurn(u, parent)           // затем разворачиваем вершину и родителя
}

// ZagZag два Левых разворота
func (t *Tree) ZagZag(u *Node) {
	parent := u.parent           // родитель вершины Ю
	grandParent := parent.parent // прародитель вершины Ю

	t.LinkingWithParent(u, grandParent)  // меняем указатели
	t.SmallLeftTurn(parent, grandParent) // сначала разворачиваем родителя и прародителя
	t.SmallLeftTurn(u, parent)           // затем разворачиваем вершину и родителя
}

// ZigZag правое, затем левое вращение
func (t *Tree) ZigZag(u *Node) {
	p := u.parent
	gP := p.parent

	t.LinkingWithParent(u, gP)

	gP.right = u.left
	if u.left != nil {
		u.left.parent = gP
	}

	t.UpdateSum(gP)

	p.left = u.right

	if u.right != nil {
		u.right.parent = p
	}

	t.UpdateSum(p)

	u.left = gP
	gP.parent = u

	u.right = p
	p.parent = u

	t.UpdateSum(u)
}

func (t *Tree) ZagZig(u *Node) {
	p := u.parent
	gP := p.parent

	t.LinkingWithParent(u, gP)

	p.right = u.left
	if u.left != nil {
		u.left.parent = p
	}

	t.UpdateSum(p)

	gP.left = u.right

	if u.right != nil {
		u.right.parent = gP
	}

	t.UpdateSum(gP)

	u.left = p
	p.parent = u

	u.right = gP
	gP.parent = u

	t.UpdateSum(u)
}

func (t *Tree) SmallRightTurn(currNode, itsParent *Node) {
	itsParent.left = currNode.right
	if currNode.right != nil {
		currNode.right.parent = itsParent
	}
	t.UpdateSum(itsParent)

	currNode.right = itsParent
	itsParent.parent = currNode
	t.UpdateSum(currNode)
}

func (t *Tree) SmallLeftTurn(currNode, itsParent *Node) {
	itsParent.right = currNode.left
	if currNode.left != nil {
		currNode.left.parent = itsParent
	}
	t.UpdateSum(itsParent)

	currNode.left = itsParent
	itsParent.parent = currNode
	t.UpdateSum(currNode)
}

func (t *Tree) UpdateSum(node *Node) {
	if node == nil {
		return
	}

	sum := node.key
	if node.left != nil {
		sum += node.left.sum
	}

	if node.right != nil {
		sum += node.right.sum
	}

	node.sum = sum
}

func (t *Tree) Remove(key int) {
	if !t.SearchInRoot(key) {
		return
	}

	found := t.root
	t.root = nil
	if found.left != nil {
		found.left.parent = nil
	}
	if found.right != nil {
		found.right.parent = nil
	}

	leftSubTree := &Tree{root: found.left}
	rightSubTree := &Tree{root: found.right}

	t.Merge(leftSubTree, rightSubTree)
}

func (t *Tree) Merge(lTree, rTree *Tree) {
	lTree.SplayMax()

	if lTree.root == nil {
		t.root = rTree.root
		return
	}

	if rTree.root == nil {
		t.root = lTree.root
		return
	}

	lTree.root.right = rTree.root
	rTree.root.parent = lTree.root
	t.UpdateSum(lTree.root)
	t.root = lTree.root
}

func (t *Tree) SplayMax() {
	if t.root == nil {
		return
	}

	node := t.root
	for {
		if node.right == nil {
			t.Splay(node)
			return
		}

		node = node.right
	}
}

func (t *Tree) SearchInRoot(key int) bool {
	if t.root == nil {
		return false
	}

	return t.SearchInTree(key, t.root)
}

func (t *Tree) SearchInTree(key int, node *Node) bool {
	if node.key == key { // ключ найден, вытягиваем его наверх функцией сплей
		t.Splay(node)
		return true
	}

	if key < node.key {
		if node.left == nil {
			t.Splay(node)
			return false
		}
		return t.SearchInTree(key, node.left) // ищем рекурсивно в левом поддереве
	}

	if node.right == nil {
		t.Splay(node)
		return false
	}
	return t.SearchInTree(key, node.right) // ищем рекурсивно в правом поддереве
}

func (t *Tree) Sum(left, right int) int {
	if t.root == nil {
		t.sum = 0
		return 0
	}

	leftTree, rightTree := t.Split(left - 1)

	if rightTree.root == nil {
		t.Merge(leftTree, rightTree)
		t.sum = 0
		return 0
	}

	leftOfRightTree, rightOfRightTree := rightTree.Split(right)

	if leftOfRightTree.root == nil {
		rightTree.Merge(leftOfRightTree, rightOfRightTree)
		t.Merge(leftTree, rightTree)
		t.sum = 0
		return 0
	}

	res := leftOfRightTree.root.sum
	rightTree.Merge(leftOfRightTree, rightOfRightTree)
	t.Merge(leftTree, rightTree)
	t.sum = res
	return res
}

func (t *Tree) Split(key int) (*Tree, *Tree) {
	var left *Node
	var right *Node

	t.SearchInTree(key, t.root)

	if t.root.key <= key {
		left = t.root
		right = t.root.right
		if right != nil {
			left.right = nil
			right.parent = nil
		}
		t.UpdateSum(left)
		t.root = nil
	} else {
		left = t.root.left
		right = t.root
		if left != nil {
			left.parent = nil
			right.left = nil
		}
		t.UpdateSum(right)
		t.root = nil
	}
	return &Tree{root: left}, &Tree{root: right}
}
