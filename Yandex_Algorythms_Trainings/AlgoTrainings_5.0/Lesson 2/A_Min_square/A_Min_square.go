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

	k := 0
	fmt.Fscan(in, &k)

	var minX, minY int
	fmt.Fscan(in, &minX, &minY)
	maxX, maxY := minX, minY

	for i := 0; i < k-1; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		minX = min(minX, x)
		minY = min(minY, y)
		maxX = max(maxX, x)
		maxY = max(maxY, y)
	}

	fmt.Fprint(out, minX, minY, maxX, maxY)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
