package main

import (
	"bufio"
	"fmt"
	"os"
)

//Вы играете в игру <<Камни>>: игру для двух игроков с двумя наборами камней по n m штук.
//С каждым ходом один игрок может забрать следующие комбинации камней:
//1. Взять один камень из любого набора.
//2. Взять два камня из какого-то одного набора
//3. Взять два камня из одного и один из другого.
// Когда камень забрали, он выходит из игры. Побеждает игрок, который заберет последний камень. Первый ход за вами.
// Вы и ваш оппонент играете оптимально.
// В первой строке содержится два числа n (1≤n≤1000), m (1≤m≤1000) — количество ваших камней и количество камней у вашего оппонента.

//В единственной строке выведите Loose, если вы заведомо проиграете, и Win, иначе.

// Пример 1
//10 2
// Вывод Loose

// Пример 2
//4 5
// Вывод Win

// Пример 3
//6 8
// Вывод Loose

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var (
		n int //row
		m int //col
		//res string = "Win"
	)
	fmt.Fscanln(in, &n, &m)

	n++
	m++

	table := make([][]string, n)
	for i := range table {
		table[i] = make([]string, m)
	}
	table[0][0] = "L"
	table[1][0] = "W"
	table[0][1] = "W"
	table[1][1] = "L"

	for i := 2; i < n; i++ {
		if table[i-1][0] == "W" && table[i-2][0] == "W" {
			table[i][0] = "L"
		} else {
			table[i][0] = "W"
		}
	}

	for i := 2; i < m; i++ {
		if table[0][i-1] == "W" && table[0][i-2] == "W" {
			table[0][i] = "L"
		} else {
			table[0][i] = "W"
		}
	}

	for i := 2; i < n; i++ {
		if table[i-1][1] == "W" && table[i-2][1] == "W" && table[i][0] == "W" && table[i-2][0] == "W" {
			table[i][1] = "L"
		} else {
			table[i][1] = "W"
		}
	}

	for i := 2; i < m; i++ {
		if table[1][i-1] == "W" && table[1][i-2] == "W" && table[0][i] == "W" && table[0][i-2] == "W" {
			table[1][i] = "L"
		} else {
			table[1][i] = "W"
		}
	}

	for i := 2; i < n; i++ {
		for j := 2; j < m; j++ {
			if table[i-1][j] == "W" && table[i-2][j] == "W" && table[i][j-1] == "W" && table[i-2][j-1] == "W" && table[i][j-2] == "W" && table[i-1][j-2] == "W" {
				table[i][j] = "L"
			} else {
				table[i][j] = "W"
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {

			fmt.Fprintf(out, "%d,%d. %s\t", i, j, table[i][j])
		}
		fmt.Fprintf(out, "\n")
	}

	fmt.Fprintln(out, table[n-1][m-1])
}
