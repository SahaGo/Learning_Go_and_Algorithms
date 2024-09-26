package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	w, n, m := 0, 0, 0

	fmt.Fscan(in, &w, &n, &m)

	firstPart := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &firstPart[i])
	}

	secondPart := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &secondPart[i])
	}

	l, r := 0, w
	for l < r {
		mid := (l + r) / 2

		if check(firstPart, secondPart, w, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	answer := min(max(getLen(firstPart, l), getLen(secondPart, w-l)), max(getLen(firstPart, l-1), getLen(secondPart, w-l+l)))
	fmt.Fprint(out, answer)
}

func check(firstPart []int, secondPart []int, w int, left int) bool {
	right := w - left
	return getLen(firstPart, left) < getLen(secondPart, right)
}

func getLen(arr []int, size int) int {
	lines := 0
	nowSize := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > size {
			return math.MaxInt
		}
		if arr[i] <= nowSize {
			nowSize -= arr[i] + 1
		} else {
			nowSize = size - arr[i] - 1
			lines += 1
		}
	}
	return lines
}
