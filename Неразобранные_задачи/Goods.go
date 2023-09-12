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
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var countGoods int

		fmt.Fscan(in, &countGoods)
		sliceGoods := make([]int, countGoods)

		for i := 0; i < countGoods; i++ {
			fmt.Fscan(in, &sliceGoods[i])
		}

		mapGoods := make(map[int]int)

		for _, v := range sliceGoods {
			if _, ok := mapGoods[v]; !ok {
				mapGoods[v] = 1
			} else {
				mapGoods[v]++
			}
		}

		var res int

		for k, v := range mapGoods {
			res += (v - v/3) * k
		}
		fmt.Fprintln(out, res)
	}
}
