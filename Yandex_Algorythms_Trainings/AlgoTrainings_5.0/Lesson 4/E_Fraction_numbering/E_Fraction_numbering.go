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

	var n int64
	fmt.Fscan(in, &n)

	var l int64 = 0
	var r int64 = n + 1

	for l < r {
		mid := (l + r) / 2
		if sumProg(mid) >= n {
			r = mid
		} else {
			l = mid + 1
		}
	}
	diag := l
	posInDiag := n - sumProg(diag-1)

	if diag%2 == 1 {
		numerator := posInDiag
		denominator := diag - posInDiag + 1

		fmt.Fprintf(out, "%d/%d", numerator, denominator)
	} else {
		numerator := diag - posInDiag + 1
		denominator := posInDiag

		fmt.Fprintf(out, "%d/%d", numerator, denominator)
	}
}

func sumProg(n int64) int64 {
	// защитимся от переполнения
	var odd int64
	var even int64
	if n%2 == 0 {
		odd = n
		even = n + 1
	} else {
		odd = n + 1
		even = n
	}
	return odd / 2 * even
}
