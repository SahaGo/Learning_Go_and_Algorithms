package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Data struct {
	num    int
	maxNum int
}

type Stack struct {
	d []Data
}

// Push добавляет Data на верхушку стека
func (s *Stack) Push(num int) {
	var res Data
	res.num = num

	if s.IsEmpty() {
		res.maxNum = num
	} else {
		prevMax := s.GetMax()

		if prevMax >= num {
			res.maxNum = prevMax
		} else {
			res.maxNum = num
		}
	}
	s.d = append(s.d, res)
}

// Pop удаляет с верха стека и возвращает значение num
func (s *Stack) Pop() int {
	res := math.MinInt
	if s.IsEmpty() {
		return res
	} else {
		i := len(s.d) - 1
		res = s.d[i].num
		s.d = s.d[:i]
	}
	return res
}

// Peek возвращает num с верха стека
func (s *Stack) Peek() int {
	res := s.d[len(s.d)-1].num
	return res
}

// GetMax возвращает maxNum с верха стека
func (s *Stack) GetMax() int {
	if s.IsEmpty() {
		return 0
	}
	return s.d[len(s.d)-1].maxNum
}

// Size возвращает размер стека
func (s *Stack) Size() int {
	size := len(s.d)
	return size
}

// IsEmpty проверка, есть ли элементы в стеке
func (s *Stack) IsEmpty() bool {
	if len(s.d) == 0 || s.d == nil {
		return true
	}
	return false
}

func CreateStack() Stack {
	stack := Stack{}
	stack.d = make([]Data, 0)
	return stack
}

type StackMethods interface {
	Push(num int)
	Pop() int
	Peek() int
	GetMax() int
	Size() int
	IsEmpty() bool
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count int

	fmt.Fscan(in, &count)

	arr := make([]int, count)

	for i := 0; i < count; i++ {
		var scan int
		fmt.Fscan(in, &scan)
		arr[i] = scan
	}

	var window int
	fmt.Fscan(in, &window)

	leftStack := CreateStack()
	rightStack := CreateStack()

	for i := 0; i < window; i++ {
		curr := arr[i]
		leftStack.Push(curr)
	}

	res := make([]int, 0)
	res = append(res, leftStack.GetMax())

	for i := window; i < count; i++ {
		curr := arr[i]

		if rightStack.IsEmpty() {
			reverseOrder(&leftStack, &rightStack)
		}

		rightStack.Pop()
		leftStack.Push(curr)

		res = append(res, maxNum(rightStack.GetMax(), leftStack.GetMax()))
	}

	for _, v := range res {
		fmt.Fprint(out, v, " ")
	}
}

func reverseOrder(left, right *Stack) {
	for !left.IsEmpty() {
		data := left.Pop()
		right.Push(data)
	}
}

func maxNum(r, l int) int {
	if r >= l {
		return r
	}
	return l
}
