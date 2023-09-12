package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var a int

	fmt.Fscan(in, &a)
	arrA := make([]int, a+1)
	for i := len(arrA) - 1; i >= 0; i-- {
		fmt.Fscan(in, &a)
		arrA[i] = a
	}

	var b int
	fmt.Fscan(in, &b)
	arrB := make([]int, b+1)
	for i := len(arrB) - 1; i >= 0; i-- {
		fmt.Fscan(in, &b)
		arrB[i] = b
	}

	res := []int{}
	if len(arrA) < len(arrB) {
		res = append(res, sumUp(arrA, arrB)...)
	} else {
		res = append(res, sumUp(arrB, arrA)...)
	}

	fmt.Fprintln(out, len(res)-1)
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Fprint(out, res[i], " ")
	}
}

func sumUp(min []int, max []int) []int {
	res := max
	for i := 0; i < len(min); i++ {
		res[i] = res[i] + min[i]
	}
	return res
}
