package main

import (
	"bufio"
	"fmt"
	"os"
)

//Реализуйте QuickSort.
//В первой строке задано число n (1≤n≤100000).
//Во второй строке задано n чисел ai(1≤ai≤10^9).
//Выведите остмортированный массив чисел.

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

// Пример 3
//Ввод
//3
//1 11 1
//Вывод 1 1 11

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
	quickSort(items, 0, len(items)-1)
	fmt.Fprintln(out, items)
}

func quickSort(items []int, left int, right int) {
	if left < right {
		p := partQuickSort(items, left, right)
		quickSort(items, left, p-1)
		quickSort(items, p, right)
	}
}

func partQuickSort(items []int, left int, right int) int {
	pivot := items[(left+right)/2]

	for left <= right {

		for items[left] < pivot {
			left++
		}

		for items[right] > pivot {
			right--
		}

		if left <= right {
			items[left], items[right] = items[right], items[left]
			left++
			right--
		}
	}
	return left
}
