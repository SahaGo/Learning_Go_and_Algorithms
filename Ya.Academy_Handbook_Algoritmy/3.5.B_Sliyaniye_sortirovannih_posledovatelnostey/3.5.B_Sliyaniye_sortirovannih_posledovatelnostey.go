package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

//Задано n отсортированных по неубыванию последовательностей.
//Требуется найти отсортированную по неубыванию конкатенацию этих последовательностей.
//В первой строке задано одно число n (1≤n≤20) — количество отсортированных последовательностей.
//
//Каждая из следующих последовательностей задано в формате: В первой строке mi(1≤mi≤10^5) — количество элементов последовательности.
//Во второй сами элементы ai(1≤ai≤10^9).
//
//Гарантируется, что
//∑n,i=1 mi≤10^5.
//В единственной строке выведите ответ на задачу.

//Пример 1
//Ввод
//3
//3
//1 2 3
//2
//1 2
//4
//3 5 6 7
//Вывод 1 1 2 2 3 3 5 6 7

//Пример 2
//Ввод
//2
//2
//1 10
//3
//7 9 11
//Вывод 1 7 9 10 11

//Пример 3
//Ввод
//1
//10
//1 2 3 4 5 6 7 8 9 10
//Вывод 1 2 3 4 5 6 7 8 9 10

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	items := make([][]int, n)
	for i := 0; i < n; i++ {
		var m int
		fmt.Fscan(in, &m)

		var a int
		item := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &a)
			item[j] = a
		}
		items[i] = item
	}

	final := merge(items)

	fmt.Fprintln(out, "test", final)
}

func merge(items [][]int) []int {
	final := []int{}
	indexes := make([]int, len(items))

	for candidatesAvailable(items, indexes) {
		cand := make([]int, 0)
		candIndx := make([]*int, 0)

		for i := 0; i < len(indexes); i++ {
			if indexes[i] < len(items[i]) {
				cand = append(cand, items[i][indexes[i]])
				candIndx = append(candIndx, &indexes[i])
			}
		}

		candidateMin, candidateMinIndex := findMin(cand, candIndx)

		final = append(final, candidateMin)
		*candidateMinIndex++
	}

	return final
}

func findMin(candidates []int, indexes []*int) (candidateMin int, candidateMinIndex *int) {
	var currCandMin = math.MaxInt
	var currCandMinIndx *int = nil
	for indx := range candidates {
		currCand := candidates[indx]
		currCandIndx := indexes[indx]
		if currCand < currCandMin {
			currCandMin = currCand
			currCandMinIndx = currCandIndx
		}
	}

	return currCandMin, currCandMinIndx
}

func candidatesAvailable(items [][]int, indexes []int) bool {
	for i := 0; i < len(indexes); i++ {
		if indexes[i] < len(items[i]) {
			return true
		}
	}
	return false
}
