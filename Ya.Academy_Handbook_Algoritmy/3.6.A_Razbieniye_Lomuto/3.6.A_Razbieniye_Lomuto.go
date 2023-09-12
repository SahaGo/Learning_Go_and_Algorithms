package main

import (
	"bufio"
	"fmt"
	"os"
)

//Постройте разбиение Ломуто относительно первого числа.
//В первой строке задано число n (1≤n≤100000).
//Во второй строке задано n чисел ai(1≤ai≤10^9).
//Выведите разбиение Ломуто.

//Пример 1
//Ввод
//9
//4 7 2 5 3 1 8 9 6
//Вывод 1 2 3 4 7 5 8 9 6

//Пример 2
//Ввод
//4
//3 4 7 17
//Вывод 3 4 7 17

//Пример 3
//Ввод
//5
//1 3 2 9 10
//Вывод 1 3 2 9 10

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
	LomutoSort(items, 0, len(items)-1)
	fmt.Fprintln(out, items)
}

func LomutoSort(items []int, left int, right int) {
	if left < right {
		p := partLomutoSort(items, left, right)
		LomutoSort(items, left, p-1)
		LomutoSort(items, p+1, right)
	}
}

func partLomutoSort(items []int, start, end int) int {
	left := start
	for i := start; i < end; i++ {
		if items[i] <= items[end] {

			items[i], items[left] = items[left], items[i]
			left++
		}
	}
	items[left], items[end] = items[end], items[left]
	return left
}
