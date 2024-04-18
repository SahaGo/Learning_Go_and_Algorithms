package main

import (
	"fmt"
	"math"
)

func main() {
	var firstPrev, secPrev, firstNow, secNow, isHome int

	fmt.Scanf("%d:%d\n%d:%d\n%d", &firstPrev, &secPrev, &firstNow, &secNow, &isHome)

	ans := 0

	// неполное реш
	//if firstPrev+firstNow == secPrev+secNow {
	//	if isHome == 2 && firstPrev < secPrev {
	//		ans += 1
	//	}
	//} else if firstPrev+firstNow < secPrev+secNow {
	//	ans = secPrev + secNow - firstPrev - firstNow
	//	if isHome == 2 {
	//		ans += 1
	//	}
	//}

	// реш с разбора
	if isHome == 1 {
		score1 := firstPrev*100 + firstNow*101
		score2 := secPrev*101 + secNow*100
		ans = int(math.Max(0, float64((score2-score1+101)/101)))
	} else {
		score1 := firstPrev*101 + firstNow*100
		score2 := secPrev*100 + secNow*101
		ans = int(math.Max(0, float64((score2-score1+100)/100)))
	}
	//fmt.Println(firstPrev, secPrev, firstNow, secNow, isHome)
	fmt.Println(ans)
}
