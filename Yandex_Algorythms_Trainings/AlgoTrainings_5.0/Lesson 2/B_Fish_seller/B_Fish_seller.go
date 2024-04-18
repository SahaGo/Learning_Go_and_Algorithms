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

	var daysCount, freshness int
	fmt.Fscan(in, &daysCount, &freshness)

	days := make([]int, daysCount)
	buff := 0
	for i := 0; i < daysCount; i++ {
		//d := 0
		fmt.Fscan(in, &buff)
		days[i] = buff
	}

	maxPrice := 0
	for i := 0; i < daysCount; i++ {
		for j := 0; j <= freshness; j++ {
			ind := i + j
			if ind >= daysCount {
				break
			}
			m := days[ind] - days[i]
			maxPrice = max(maxPrice, m)
		}

	}
	fmt.Fprintln(out, maxPrice)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
