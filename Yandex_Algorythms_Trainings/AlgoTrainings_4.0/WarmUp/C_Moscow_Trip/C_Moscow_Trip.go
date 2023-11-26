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
	var xA, yA, xB, yB float64
	fmt.Fscan(In, &xA, &yA, &xB, &yB)

	r1 := rO(xA, yA)
	r2 := rO(xB, yB)
	throughCenter := r1 + r2
	//poDuge := 0
	//if r1 == r2 {
	//
	//}
	fmt.Fprint(Out, throughCenter)
}

func sBetweenXY(xA, yA, xB, yB float64) float64 {
	return math.Sqrt(math.Pow(xB-xA, 2) + math.Pow(yB-yA, 2))
}

func rO(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}
