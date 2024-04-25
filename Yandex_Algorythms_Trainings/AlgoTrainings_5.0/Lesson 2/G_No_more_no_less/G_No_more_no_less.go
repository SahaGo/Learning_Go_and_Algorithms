package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	In  = bufio.NewReader(os.Stdin)
	Out = bufio.NewWriter(os.Stdout)
)

func main() {

	defer Out.Flush()

	var testsCount int
	fmt.Fscan(In, &testsCount)

	for i := 0; i < testsCount; i++ {
		var numCount int
		fmt.Fscan(In, &numCount)
		res := make([]int, 0, numCount)
		var length, num, minNum, prevNum int
		minNum = math.MaxInt
		for j := 0; j < numCount; j++ {

			fmt.Fscan(In, &num)
			length++
			//fmt.Fprintln(Out, "test:", i, length, num)

			if num <= minNum {
				minNum = num
			}
			//if minNum < length {
			if num < length || minNum < length || prevNum < length {
				length--
				res = append(res, length)
				length = 1
				minNum = math.MaxInt
				prevNum = num
			}
			if j == numCount-1 {
				res = append(res, length)
			}
		}
		//fmt.Fprintln(Out, "test:", res)
		printResult(res, i)
	}
}

func printResult(arr []int, i int) {
	//defer Out.Flush()
	if i != 0 {
		fmt.Fprintln(Out)
	}
	l := len(arr)
	fmt.Fprintln(Out, l-1)
	for k := 1; k < l; k++ {
		if k == l-1 {
			fmt.Fprint(Out, arr[k])
		} else {
			fmt.Fprint(Out, arr[k], " ")
		}

	}
}
