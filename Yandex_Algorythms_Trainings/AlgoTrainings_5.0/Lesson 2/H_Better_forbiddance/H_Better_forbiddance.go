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

	var aI, aJ, num int

	fmt.Fscan(in, &aI, &aJ)

	table := make([][]int, aI)
	for i := 0; i < aI; i++ {
		table[i] = make([]int, aJ)
		for j := 0; j < aJ; j++ {
			fmt.Fscan(in, &num)
			table[i][j] = num
		}
	}

	fR, fC, _ := findMax(table, -1, -1) // absolut max
	_, banRowC, _ := findMax(table, fR, -1)
	_, _, banRowVal := findMax(table, fR, banRowC)
	banColR, _, _ := findMax(table, -1, fC)
	_, _, banColVal := findMax(table, banColR, fC)
	if banRowVal < banColVal {
		fmt.Fprintln(out, fR+1, banRowC+1)
	} else {
		fmt.Fprintln(out, banColR+1, fC+1)
	}

}

func findMax(f [][]int, bannedRow, bannedCol int) (int, int, int) {
	var ans, r, c int
	for i := 0; i < len(f); i++ {
		if i != bannedRow {
			for j := 0; j < len(f[0]); j++ {
				if j != bannedCol && f[i][j] > ans {
					ans = f[i][j]
					r = i
					c = j
				}
			}
		}
	}
	return r, c, ans
}
