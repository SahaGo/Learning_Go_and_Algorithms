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
	var n int
	var m int

	fmt.Fscanf(in, "%d %d", &n, &m)
	arrA := make([][]int, n)

	for i := range arrA {
		arrA[i] = make([]int, m)
	}

	var buff int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &buff)
			arrA[i][j] = buff
		}
	}

	arrB := make([][]int, n)
	for i := range arrB {
		arrB[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &buff)
			arrB[i][j] = buff
		}
	}

	arrC := make([][]int, n)
	for i := range arrC {
		arrC[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			arrC[i][j] = arrA[i][j] + arrB[i][j]
			fmt.Fprint(out, arrC[i][j], " ")
		}
		fmt.Fprintln(out, "")
	}
}
