package main

import (
	"bufio"
	"fmt"
	"os"
)

//Реализуйте сортировку слиянием
//В первой строке задано число n (1≤n≤100000)
//Во второй строке задано n чисел ai (0≤ai≤10^9)
//Выведите отсортированнй список

//Пример 1
//Ввод
//7
//13 17 37 73 31 19 23
//Вывод 13 17 19 23 31 37 73

//Пример 2
//Ввод
//4
//18 20 3 17
//Вывод 3 17 18 20

//Пример 3
//Ввод
//3
//0 11 0
//Вывод 0 0 11

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n int
	var m int

	fmt.Fscanln(in, &n)
	items := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &m)
		items[i] = m
	}

	final := mergeSort(items)
	fmt.Fprintln(out, final)
}

func mergeSort(items []int) []int {
	if len(items) == 1 {
		return items
	} else {
		firstHalf := items[:len(items)/2]
		secondHalf := items[len(items)/2:]
		sortedFirstHalf := mergeSort(firstHalf)
		sortedSecondHalf := mergeSort(secondHalf)
		final := merge(sortedFirstHalf, sortedSecondHalf)
		return final
	}
}

func merge2(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0

	for i < len(a) || j < len(b) {

		if i <= len(a)-1 && j > len(b)-1 {
			final = append(final, a[i])
			i++
			continue
		}

		if j <= len(b)-1 && i > len(a)-1 {
			final = append(final, b[j])
			j++
			continue
		}
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	return final
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0

	for i < len(a) || j < len(b) {

		if j > len(b)-1 {
			final = append(final, a[i])
			i++
			continue
		}

		if i > len(a)-1 {
			final = append(final, b[j])
			j++
			continue
		}
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	return final
}
