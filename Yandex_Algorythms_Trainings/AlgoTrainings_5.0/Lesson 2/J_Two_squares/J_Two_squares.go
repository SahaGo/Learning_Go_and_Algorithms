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

	var m, n int
	fmt.Fscan(in, &m, &n)

	table := make([][]rune, m)
	buff := ""
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &buff)
		table[i] = []rune(buff)
	}

	for i := 0; i < m; i++ {
		if inspect(table, 0, 0, i, n, 'a') && inspect(table, i, 0, m, n, 'b') {
			printAnswer(table, out)
			return
		}
	}

	for j := 0; j < n; j++ {
		if inspect(table, 0, 0, m, j, 'a') && inspect(table, 0, j, m, n, 'b') {
			printAnswer(table, out)
			return
		}
	}
	fmt.Fprint(out, "NO")

}

func printAnswer(table [][]rune, out *bufio.Writer) {
	fmt.Fprint(out, "YES\n")
	for _, strings := range table {
		fmt.Fprintf(out, "%s\n", string(strings))
	}
}

func inspect(table [][]rune, fI, fJ, tI, tJ int, fill rune) bool {
	minI, minJ := 201, 201
	maxI, maxJ := -1, -1
	count := 0
	for i := fI; i < tI; i++ {
		for j := fJ; j < tJ; j++ {
			if table[i][j] != '.' {
				minI = min(minI, i)
				maxI = max(maxI, i)
				minJ = min(minJ, j)
				maxJ = max(maxJ, j)
				table[i][j] = fill
				count += 1
			}
		}
	}
	return count > 0 && (maxI-minI+1)*(maxJ-minJ+1) == count
}
