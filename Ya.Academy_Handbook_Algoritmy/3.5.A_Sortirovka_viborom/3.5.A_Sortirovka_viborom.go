package main

import (
	"bufio"
	"fmt"
	"os"
)

//Реализуйте сортировку выбором.
//В первой строке задано число n (1≤n≤1000).
//
//Во второй строке задано n чисел ai (1≤ai≤10^9).
//Выведите отсортированный массив чисел.

//Пример 1
//Ввод
//7
//13 17 37 73 31 19 23
//Вывод 13 17 19 23 31 37 73

//Пример 2
//Ввод
//6
//12 18 7 11 5 17
//Вывод 5 7 11 12 17 18

//Пример 3
//Ввод
//3
//1 2 3
//Вывод 1 2 3

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

	for i := range items {
		min := i
		for k := range items {
			if items[k] > items[min] {
				min = k
			}
			items[i], items[min] = items[min], items[i]
		}
	}

	for i := range items {
		fmt.Fprint(out, items[i], " ")
	}
}
