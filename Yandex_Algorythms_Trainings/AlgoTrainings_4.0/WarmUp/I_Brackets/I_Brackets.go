package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	data []rune
}

func (s *Stack) Push(e rune) {
	s.data = append(s.data, e)
}

func (s *Stack) Top() rune {
	return s.data[len(s.data)-1]
}

func (s *Stack) Pop() {
	s.data = s.data[:len(s.data)-1]
}

func (s *Stack) IsEmpty() bool {
	if len(s.data) == 0 {
		return true
	}
	return false
}

func (s Stack) String() string {
	return string(s.data)
}

var openBracketsMap = map[rune]rune{
	'[': ']',
	'{': '}',
	'(': ')',
}

var closeBracketsMap = map[rune]rune{
	']': '[',
	'}': '{',
	')': '(',
}

var StatusSuccess string = "yes"

func CheckBrackets(inputData string) string {
	stack := Stack{data: make([]rune, 0)}
	stackPositions := Stack{data: make([]rune, 0)}
	var errorIndex int
	for i, r := range inputData {
		errorIndex = i + 1
		if _, ok := openBracketsMap[r]; ok {
			stack.Push(r)
			stackPositions.Push(rune(errorIndex))
		} else if openBracket, ok := closeBracketsMap[r]; ok {
			if stack.IsEmpty() {
				return fmt.Sprint("no")
			} else {
				if openBracket == stack.Top() {
					stack.Pop()
					stackPositions.Pop()
				} else {
					return fmt.Sprint("no")
				}
			}
		}
	}

	if !stack.IsEmpty() {

		for !stack.IsEmpty() {
			stack.Pop()

			stackPositions.Pop()
		}
		return fmt.Sprint("no")
	}

	return fmt.Sprint(StatusSuccess)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	inputData, _ := reader.ReadString('\n')

	fmt.Println(CheckBrackets(inputData))
}
