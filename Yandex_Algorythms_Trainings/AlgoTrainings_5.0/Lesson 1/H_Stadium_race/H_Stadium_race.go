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

	var l float64      // длина стадиона
	var x1, v1 float64 //начальная точка и скорость Кирилла
	var x2, v2 float64 //начальная точка и скорость Антона

	fmt.Fscan(in, &l, &x1, &v1, &x2, &v2)
	ans := math.MaxFloat64

	for i := 0; i < 2; i++ {
		deltaX := float64(int64(x2-x1+l) % int64(l))
		deltaV := v1 - v2
		if deltaV < 0 {
			deltaV = -deltaV
			deltaX = float64(int64(-deltaX+l) % int64(l))
		}
		if deltaX == 0 {
			ans = 0
		}
		if deltaV != 0 {
			ans = minimum(ans, deltaX/deltaV)
		}
		x2 = float64(int64(-x2+l) % int64(l))
		v2 = -v2
	}
	if ans == math.MaxFloat64 {
		fmt.Fprint(out, "NO")
	} else {
		fmt.Fprintln(out, "YES")
		fmt.Fprintln(out, ans)
	}
}

func minimum(a, b float64) float64 {
	if a <= b {
		return a
	}
	return b
}
