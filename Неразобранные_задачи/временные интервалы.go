package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

//type Interval struct {
//	start time.Time
//	end   time.Time
//}
type Element struct {
	time    time.Time
	isStart bool
}

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscanln(in, &testCount)
Loop:
	for test := 0; test < testCount; test++ {
		var countInterval int

		fmt.Fscanln(in, &countInterval)

		var buff, t1, t2 string

		//intervals := make([]Interval, countInterval)
		sliceintervals := make([]Element, countInterval*2)
		wrongInterval := false
		for i := 0; i < countInterval; i++ {

			fmt.Fscanln(in, &buff)

			t1, t2, _ = strings.Cut(buff, "-")

			if WrongTime(t1) || WrongTime(t2) {
				wrongInterval = true
				continue
			}

			t1, _ := time.Parse("15:04:05", t1)

			t2, _ := time.Parse("15:04:05", t2)

			//intervals[i] = Interval{t1, t2}

			sliceintervals = append(sliceintervals, Element{t1, true}, Element{t2, false})
			if t2.Before(t1) {
				wrongInterval = true
				continue
			}
		}
		if wrongInterval {
			fmt.Fprintln(out, "NO")
			continue Loop
		}

		//for i := 0; i < len(intervals)-1; i++ {
		//	first := intervals[i]
		//	for j := i + 1; j < len(intervals); j++ {
		//		second := intervals[j]
		//		if intersect(first, second) {
		//			fmt.Fprintln(out, "NO")
		//			continue Loop
		//		}
		//	}
		//}

		sort.Slice(sliceintervals, func(i, j int) bool {
			return sliceintervals[i].time.Before(sliceintervals[j].time)
		})
		for _, v := range sliceintervals {
			var ch int
			if ch > 1 {
				fmt.Fprintln(out, "NO")
				continue Loop
			} else if v.isStart {
				ch += 1
			} else {
				ch -= 1
			}

		}
		fmt.Fprintln(out, "YES")
	}
}

//func intersect(first Interval, second Interval) bool {
//	if first.end.Before(second.start) || second.end.Before(first.start) {
//		return false
//	}
//	return true
//}

func WrongTime(t string) bool {
	var hourst, minst, secst string
	var hour, min, sec int
	hourst = string(t[0]) + string(t[1])
	hour, _ = strconv.Atoi(hourst)
	minst = string(t[3]) + string(t[4])
	min, _ = strconv.Atoi(minst)
	secst = string(t[6]) + string(t[7])
	sec, _ = strconv.Atoi(secst)

	switch {
	case hour > 23 || hour < 00:
		return true
	case 00 > min || min > 59:
		return true
	case sec > 59 || sec < 00:
		return true
	}
	return false
}

//elements []Element len 4
//sort(elements)
