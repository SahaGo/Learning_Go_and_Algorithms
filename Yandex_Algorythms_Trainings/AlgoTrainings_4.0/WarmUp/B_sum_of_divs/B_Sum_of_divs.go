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

	var a, b, c, d, m, n int
	fmt.Fscanln(In, &a, &b, &c, &d)

	// divisor is eq
	if isEq(b, d) {
		n = b
		m = a + c

		if isEq(n, m) {
			fmt.Fprint(Out, 1, 1)
			return
		} else {
			num := nod(m, n)
			fmt.Fprint(Out, m/num, n/num)
			return
		}
	}

	// div isn't eq
	n = b * d
	m = a*d + c*b
	num := nod(m, n)
	fmt.Fprint(Out, m/num, n/num)

}

func isEq(x, y int) bool {
	if x == y {
		return true
	}
	return false
}

// Nod -наибольший общий делитель, алгоритм Евклида
func nod(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a >= b {
		return nod(a%b, b)
	} else {
		return nod(a, b%a)
	}
}
