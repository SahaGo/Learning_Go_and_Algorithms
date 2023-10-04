package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Stack []int

type MaxStack struct {
	elems Stack
	maxs  Stack
}

type MaxStackMethods interface {
	Push() int
	Max() int
	Pop()
	IsEmpty() bool
	CreateStack() MaxStack
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count int
	fmt.Fscanln(in, &count)

	stack := CreateStack()

	for i := 0; i < count; i++ {

		line, err := in.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		line = strings.Trim(line, "\n")

		buff := strings.Split(line, " ")

		if len(buff) == 2 {
			num, _ := strconv.Atoi(buff[1])
			stack.Push(num)
		} else {
			switch buff[0] {
			case "max":
				m := stack.Max()
				fmt.Fprintln(out, m)
			case "pop":
				stack.Pop()
			}
		}
	}
}

// CreateStack создать Стек
func CreateStack() MaxStack {
	stack := MaxStack{}
	stack.elems = make([]int, 0)
	stack.maxs = make([]int, 0)
	return stack
}

// Push добавляет входящее число в elements и максимум в maxs
func (s *MaxStack) Push(num int) {

	if !s.IsEmpty() {
		s.elems = append(s.elems, num)
		num2 := s.Max()
		if num > num2 {
			s.maxs = append(s.maxs, num)
		} else {
			s.maxs = append(s.maxs, num2)
		}

	} else {
		s.elems = append(s.elems, num)
		s.maxs = append(s.maxs, num)
	}
}

// Max возвращает последний добавленнный элемент из maxs
func (s *MaxStack) Max() int {
	i := s.maxs[len(s.maxs)-1]
	return i
}

// Pop удаляет последний добавленнный элемент из стеков
func (s *MaxStack) Pop() {
	i := len(s.elems) - 1
	s.elems = s.elems[:i]
	s.maxs = s.maxs[:i]
}

// IsEmpty проверка есть ли элементы в стеке
func (s *MaxStack) IsEmpty() bool {
	if len(s.elems) == 0 {
		return true
	}
	return false
}
