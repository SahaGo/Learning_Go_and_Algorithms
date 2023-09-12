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

	var testCount int
	fmt.Fscanln(in, &testCount)

	for test := 0; test < testCount; test++ {
		paluba1 := 0
		paluba2 := 0
		paluba3 := 0
		paluba4 := 0

		sliceships := make([]int, 10)
		scaner := 0

		for i := 0; i < len(sliceships); i++ {

			fmt.Fscan(in, &scaner)

			sliceships[i] = scaner
		}
		for _, v := range sliceships {
			switch v {
			case 1:
				paluba1++
			case 2:
				paluba2++
			case 3:
				paluba3++
			case 4:
				paluba4++
			}
		}
		var mistake bool = false
		switch {
		case paluba1 != 4:
			mistake = true
		case paluba2 != 3:
			mistake = true
		case paluba3 != 2:
			mistake = true
		case paluba4 != 1:
			mistake = true

		}
		if mistake {
			fmt.Fprintln(out, "NO")
		} else {
			fmt.Fprintln(out, "YES")
		}
	}
}
