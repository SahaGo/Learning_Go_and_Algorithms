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
	pointsLst := make([]point, 0, n)
	pointsSet := make(map[point]struct{})

	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		coord := point{x, y}
		pointsLst = append(pointsLst, coord)
		pointsSet[coord] = struct{}{}
	}

	x0, y0 := pointsLst[0].x, pointsLst[0].y                     // на случай, если в дано только одна точка
	ans := []point{{x0 + 1, y0 + 1}, {x0 + 1, y0}, {x0, y0 + 1}} // на этот случай делаем ответ

	for i := 0; i < n; i++ { // и - точка 1
		for j := 0; j < n; j++ { // жи - точка 2
			if i != j { // и!=жи, чтоб корректно можно было обработать сторону квадрата
				x1, y1 := pointsLst[i].x, pointsLst[i].y
				x2, y2 := pointsLst[j].x, pointsLst[j].y
				//x2, y2 := pointsLst[i].x, pointsLst[i].y
				//x1, y1 := pointsLst[j].x, pointsLst[j].y

				dx := x2 - x1
				dy := y2 - y1

				x3 := x1 + dy
				y3 := y1 - dx

				x4 := x3 + dx
				y4 := y3 + dy

				_, point3InSet := pointsSet[point{x3, y3}] // лежит ли точка 3 в мн-ве
				_, point4InSet := pointsSet[point{x4, y4}] // лежит ли точка 4 в мн-ве

				if point3InSet && point4InSet { // если есть обе, значит у нас уже есть квадрат
					ans = []point{} // добавлять ничего не надо
				}

				if len(ans) > 1 && point3InSet { // есл есть более длинный ответ и только точка 3 есть в мн-ве
					ans = []point{{x4, y4}} // добавить надо только точку 4
				}

				if len(ans) > 1 && point4InSet { // длинный ответ и точка 4 есть в мн-ве
					ans = []point{{x3, y3}} // добавляем точку 3
				}

				if len(ans) > 2 { // когда дано больше одной точки, нужно достроить всего две
					ans = []point{{x3, y3}, {x4, y4}} // найденные 3 и 4 добавим в ответ
				}
			}
		}
	}

	fmt.Fprintln(out, len(ans))
	for i := 0; i < len(ans); i++ {
		fmt.Fprintln(out, ans[i].x, ans[i].y)
	}
}
