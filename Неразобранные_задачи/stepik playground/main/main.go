package main

import (
	"fmt"
	"strconv"
	"time"
)

const now = 1589570165

func main() {
	var min = 12
	var sec = 13
	//fmt.Scanf("%d мин. %d", &min, &sec)
	//minTime,_:=time.Parse("04",min)
	minst := strconv.Itoa(min)
	secst := strconv.Itoa(sec)
	toParse := minst + "m" + secst + "s"
	//secTime,_:=time.ParseDuration("05",sec)

	tt, _ := time.ParseDuration(toParse)

	tm := time.Unix(now, 0).UTC()
	tm.Add(tt)
	fmt.Println(tm.Format(time.UnixDate))
}

//	scanner := "13.03.2018 14:00:15,12.03.2018 14:00:15"
//	layot := "02.01.2006 15:04:05"
//
//	t1, t2, _ := strings.Cut(scanner, ",")
//
//	time1, _ := time.Parse(layot, t1)
//
//	time2, _ := time.Parse(layot, t2)
//
//	//dur := time1.Sub(time2)
//
//	fmt.Println((time.Since(time2) - time.Since(time1)).Round(time.Second))
//
//	//fmt.Println(abs(time1, time2))
//}

//func absDiff(a, b time.Duration) time.Duration {
//	if a >= b {
//		return a - b
//	}
//	return b - a
//}
