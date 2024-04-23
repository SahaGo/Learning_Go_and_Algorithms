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

	var count int
	fmt.Fscan(in, &count)

	hourSequence := make([]int, 0, count)
	antiHourSequence := make([]int, 0, count)

	j := count - 1

	for i := 0; i < count; i++ {
		buff := 0
		fmt.Fscan(in, &buff)

		hourSequence = append(hourSequence, buff)
		antiHourSequence = append([]int{buff}, antiHourSequence...)
		j--
	}

	antiHourSequence = append(antiHourSequence[len(antiHourSequence)-1:], antiHourSequence[:len(antiHourSequence)-1]...)
	var a, b, k int

	fmt.Fscan(in, &a, &b, &k)

	maxRouletteNum := math.MinInt
	for i := a; i <= b; i++ {
		res := i / k
		ind := res % count
		if i == k {
			ind = 0
		}
		if hourSequence[ind] >= maxRouletteNum {
			maxRouletteNum = hourSequence[ind]
		}
		if antiHourSequence[ind] >= maxRouletteNum {
			maxRouletteNum = antiHourSequence[ind]
		}
	}

	fmt.Fprintln(out, maxRouletteNum, hourSequence, antiHourSequence)
}
