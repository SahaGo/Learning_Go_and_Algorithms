package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	x int
	y int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	byType := make(map[point][]point)
	for i := 0; i < n; i++ {
		var x1, y1, x2, y2 int
		fmt.Fscan(in, &x1, &y1, &x2, &y2)
		if (x1 > x2) || (x1 == x2 && y1 > y2) { // питонское (x1, y1) > (x2, y2)
			x1, y1, x2, y2 = x2, y2, x1, y1
		}
		dx := x2 - x1
		dy := y2 - y1

		key := point{x: dx, y: dy}
		_, dInMap := byType[key]

		if !dInMap {
			byType[key] = []point{}
		}
		byType[key] = append(byType[key], point{x: x1, y: y1})
	}

	mxMy := make(map[point]int)
	for i := 0; i < n; i++ {
		var x1, y1, x2, y2 int
		fmt.Fscan(in, &x1, &y1, &x2, &y2)
		if (x1 > x2) || (x1 == x2 && y1 > y2) {
			x1, y1, x2, y2 = x2, y2, x1, y1
		}
		dx := x2 - x1
		dy := y2 - y1

		key := point{x: dx, y: dy}
		for _, points := range byType[key] {
			xp, yp := points.x, points.y
			mx := x1 - xp
			my := y1 - yp
			mxMy[point{mx, my}]++
		}
	}

	var maxInPlace int
	for _, count := range mxMy {
		maxInPlace = max(maxInPlace, count)
	}

	fmt.Fprintln(out, n-maxInPlace)
}
