package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	pluS string = "+"
	mult string = "x"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count int
	fmt.Fscan(in, &count)

	//ans := make([]rune, 0, count/2)
	state := "no_odd_summand"
	for i := 0; i < count; i++ {
		var num int64
		fmt.Fscan(in, &num)

		switch state {
		case "no_odd_summand":
			if num%2 == 0 {
				fmt.Fprint(out, pluS)
				//ans = append(ans, pluS)
			} else {
				state = "last_odd"
			}
		case "last_odd":
			if num%2 == 0 {
				fmt.Fprint(out, pluS)
				state = "multiply_even"
			} else {
				fmt.Fprint(out, mult)
			}
		case "multiply_even":
			fmt.Fprint(out, mult)
		}
	}
}
