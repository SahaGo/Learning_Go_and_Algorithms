package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	for i := range arr {
		fmt.Fscan(in, &arr[i])
	}

	sort.Ints(arr)
	arr = append(arr, math.MaxInt)

	var k int
	fmt.Fscan(in, &k)

	for i := 0; i < k; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)

		leftBorderInd := binSearch(0, len(arr), arr, l)
		rightBorderInd := binSearch(0, len(arr), arr, r+1)

		fmt.Fprintln(out, rightBorderInd-leftBorderInd, " ")
	}
}

func binSearch(l, r int, arr []int, check int) int {

	for l < r {
		mid := (l + r) / 2
		if arr[mid] >= check {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
