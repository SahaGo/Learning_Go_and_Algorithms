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

	var ropesLeft int
	fmt.Fscan(in, &ropesLeft)

	ropes := make([]int, ropesLeft)
	buff := 0
	for i := 0; i < ropesLeft; i++ {
		fmt.Fscan(in, &buff)
		ropes[i] = buff
	}

	maxNum := 0
	sum := 0

	for i := 0; i < ropesLeft; i++ {
		sum += ropes[i]
		if ropes[i] >= maxNum {
			maxNum = ropes[i]
		}
	}

	//fmt.Fprintln(out, ropes, sum, maxNum)
	if maxNum <= sum-maxNum {
		fmt.Fprintln(out, sum)
	} else {
		fmt.Fprintln(out, maxNum-(sum-maxNum))
	}
}
