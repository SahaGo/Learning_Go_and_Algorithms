package main

import (
	"bufio"
	"fmt"
	"os"
)

//Задано n интервалов. Требуется найти максимальное количество взаимно непересекающихся интервалов.
//Два интервала пересекаются, если они имеют хотя бы одну общую точку.
//В первой строке задано одно число n (1≤n≤100) — количество интервалов.
//В следующих n строках заданы интервалы li,ri (1≤li≤r≤50).

// Пример 1
//3
//1 3
//2 3
//4 5
// Вывод 2

// Пример 2
//5
//1 2
//2 3
//4 5
//4 5
//5 6
// Вывод 2

// Пример 3
//2
//1 50
//49 50
// Вывод 1

type interval struct {
	start  int
	finish int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	intervals := make(map[int]interval)

	inputCount := 0
	fmt.Fscanln(in, &inputCount)

	for i := 0; i < inputCount; i++ {
		start := 0
		finish := 0

		fmt.Fscan(in, &start)
		fmt.Fscan(in, &finish)

		intervals[i] = interval{start, finish}

	}

	res := make([]interval, 0)

	for _, v := range intervals {

		leaderId, leaderInterval := leadInterval(intervals, v)

		res = append(res, intervals[leaderId])
		delete(intervals, leaderId)

		intersection(intervals, leaderInterval)
	}

	fmt.Fprintln(out, len(res))
}

func leadInterval(intervals map[int]interval, v interval) (int, interval) {
	min := v.finish
	var leaderId int
	var leaderInterval interval

	for id, val := range intervals {

		if val.finish <= min {
			min = v.finish
			leaderId = id
			leaderInterval = val
		}
	}
	return leaderId, leaderInterval
}

func intersection(intervals map[int]interval, leaderInterval interval) {

	for k, v := range intervals {
		if leaderInterval.finish >= v.start {
			delete(intervals, k)
		}
	}
}