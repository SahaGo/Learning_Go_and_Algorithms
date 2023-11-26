package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	In  = bufio.NewReader(os.Stdin)
	Out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer Out.Flush()
	var a, b int
	fmt.Fscan(In, &a, &b)

	fmt.Fprint(Out, euclidGCD(a, b))

}

func euclidGCD(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a >= b {
		return euclidGCD(a%b, b)
	} else {
		return euclidGCD(a, b%a)
	}
}
