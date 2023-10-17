package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Stack struct {
	elements []rune
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var example string
	fmt.Fscanln(in, &example)

	stack := CreateStack()
	exampleRune := []rune(example)
	var res string

	var lastOpenedBracket = -1

	for k, v := range exampleRune {

		if v == '[' || v == '{' || v == '(' {
			if !stack.IsNotEmpty() {
				lastOpenedBracket = k + 1
			}
			stack.Add(v)
			//res = strconv.Itoa(k + 1)
		}
		if v == ']' || v == '}' || v == ')' {
			if !stack.IsNotEmpty() {
				res = strconv.Itoa(k + 1)
				break
			}
			matching := stack.Pull()
			switch {
			case v == ']' && matching == '[':
				stack.Delete()
			case v == '}' && matching == '{':
				stack.Delete()
			case v == ')' && matching == '(':
				stack.Delete()
			default:
				res = strconv.Itoa(k + 1)
				break
			}
		}
	}

	if res == "" {
		if stack.IsNotEmpty() {
			res = strconv.Itoa(lastOpenedBracket)
		} else {
			res = "Success"
		}
	}

	fmt.Fprint(out, res)
}

func CreateStack() Stack {
	stack := Stack{}
	stack.elements = make([]rune, 0)
	return stack
}

func (s *Stack) Add(bracket rune) {
	s.elements = append(s.elements, bracket)
}

func (s *Stack) Pull() rune {
	i := s.elements[len(s.elements)-1]
	return i
}

func (s *Stack) Delete() {
	i := len(s.elements) - 1
	s.elements = s.elements[:i]
}

func (s *Stack) IsNotEmpty() bool {
	if len(s.elements) == 0 {
		return false
	}
	return true
}
