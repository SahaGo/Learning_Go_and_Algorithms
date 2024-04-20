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

	table := make([][]int, 10)

	for i := 0; i < 10; i++ {
		table[i] = make([]int, 10)
	}

	var count, x, y int
	fmt.Fscan(in, &count)

	for i := 0; i < count; i++ {
		fmt.Fscan(in, &x, &y)
		table[x][y] = 1
	}
	var res int

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {

			if table[i][j] != table[i][j-1] {
				res += 1
			}
			if table[i][j] != table[i-1][j] {
				res += 1
			}
		}

	}
	fmt.Fprintln(out, res)
}
