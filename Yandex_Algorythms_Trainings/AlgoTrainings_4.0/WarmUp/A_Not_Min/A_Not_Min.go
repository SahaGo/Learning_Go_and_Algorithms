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

type stack struct {
	arr  []int
	max  int
	min  int
	stat bool
}

func main() {
	defer Out.Flush()

	var n, m, l, r int
	fmt.Fscanln(In, &n, &m)

	s := stack{
		arr: make([]int, n),
	}

	for i := 0; i < n; i++ {
		fmt.Fscan(In, &s.arr[i])
	}

	for i := 0; i < m; i++ {
		fmt.Fscan(In, &l, &r)

		s.max = s.arr[l]
		s.min = s.arr[l]
		s.stat = false
		if s.isFound(l, r) {
			fmt.Fprintln(Out, s.max)
		} else {
			fmt.Fprintln(Out, "NOT FOUND")
		}
	}
}

func (s *stack) isFound(l, r int) bool {
	for i := l; i < r; i++ {
		if s.arr[i+1] > s.arr[i] {
			s.max = s.arr[i+1]
			break
		}
		if s.arr[i+1] < s.arr[i] {
			s.min = s.arr[i+1]
		}
	}
	if s.max <= s.min {
		return false
	}
	return true
}
