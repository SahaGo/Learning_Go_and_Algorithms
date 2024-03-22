package main

import (
	"fmt"
	"math"
)

type interval struct {
	begin float64
	end   float64
}

func main() {
	var p, v, q, m float64
	fmt.Scan(&p, &v, &q, &m)

	//превышение по памяти!
	//paintedTrees := make(map[int]bool)
	//count(p, v, paintedTrees)
	//count(q, m, paintedTrees)
	//fmt.Println(len(paintedTrees))

	firstInt := interval{p - v, p + v}
	secondInt := interval{q - m, q + m}

	intervalIntersection(firstInt, secondInt)
}

//func count(x, y int, pT map[int]bool) {
//
//	for i := y; i >= 0; i-- {
//		indMinus := i * -1
//		if _, ok := pT[x+i]; !ok {
//			pT[x+i] = true
//		}
//		if _, ok := pT[x+indMinus]; !ok {
//			pT[x+indMinus] = true
//		}
//	}
//}

func intervalIntersection(first, second interval) {
	if math.Max(first.begin, second.begin) <= math.Min(first.end, second.end) {
		fmt.Print(int64(math.Max(first.end, second.end)) - int64(math.Min(first.begin, second.begin)) + 1)
	} else {
		fmt.Print(int64(first.end-first.begin+1) + int64(second.end-second.begin+1))
	}
}
