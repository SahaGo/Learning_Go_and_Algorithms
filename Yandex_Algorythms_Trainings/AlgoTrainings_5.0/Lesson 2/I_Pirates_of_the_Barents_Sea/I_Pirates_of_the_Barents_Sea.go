package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type position struct {
	row int
	col int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	size := 0
	fmt.Fscan(in, &size)

	table := make([][]int, size)
	arr := make([]position, size)
	for i := 0; i < size; i++ {
		table[i] = make([]int, size)
		r, c := 0, 0
		fmt.Fscan(in, &r, &c)
		arr[i] = position{r, c}
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].row == arr[j].row {
			return arr[i].col < arr[j].col
		}
		return arr[i].row < arr[j].row
	})

	res := math.MaxInt

	for col := 1; col <= size; col++ {
		now := 0
		for i := 0; i < size; i++ {
			now += abs(arr[i].row-(i+1)) + abs(arr[i].col-col)
		}
		res = min(res, now)
	}
	fmt.Fprintln(out, res)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
