package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int //= 1
	fmt.Fscan(in, &testCount)

	for test := 0; test < testCount; test++ {

		var countRow, countColumn int

		fmt.Fscan(in, &countRow, &countColumn)

		//часть 1 считываю талицу
		table := make([][]int, countRow)
		for i := range table {
			table[i] = make([]int, countColumn)
		}

		var buff int

		for i := 0; i < countRow; i++ {
			for j := 0; j < countColumn; j++ {
				fmt.Fscan(in, &buff)
				table[i][j] = buff
			}
		}

		//  часть 2 считываю столбцы, по которым кликнули
		var countClick int
		fmt.Fscan(in, &countClick)
		//fmt.Println(countClick)
		clickOrder := make([]int, countClick)
		for i := 0; i < countClick; i++ {
			fmt.Fscan(in, &buff)
			clickOrder[i] = buff
		}
		// часть 3 Сортировка по столбцам
		sliceSorted := make([]int, countRow)

		currTable := table
		for _, v := range clickOrder {

			for indexRow, row := range table {
				for r, val := range row {

					if r == v-1 {
						//fmt.Print("\n", val)

						sliceSorted[indexRow] = val
					}

				}

			}

			sort.Ints(sliceSorted)

			currTable = sortTable(countRow, countColumn, sliceSorted, currTable, v)
			//fmt.Println(tableRes)

		}
		// for i := 0; i < countRow; i++ {
		// 	for j := 0; j < countColumn-1; j++ {

		// 		fmt.Fprintf(out, "%d\t", currTable[i][j])
		// 	}
		// 	fmt.Fprintf(out, "%d\n", currTable[i][countColumn-1])
		// }

		for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {

			fmt.Fprintf(out, "%d,%d. %s\t", i, j, table[i][j])
		}
		fmt.Fprintf(out, "\n")
	}

		fmt.Fprintf(out, "\n")
	}


}

func sortTable(countRow int, countColumn int, sliceSorted []int, table [][]int, click int) [][]int {
	//copyTable:=copy(table)//create `table` copy with deep copy = copyTable
	copyTable := make([][]int, countRow)
	for i := range copyTable {
		copyTable[i] = make([]int, countColumn)
	}

	for i := 0; i < countRow; i++ {
		buff := make([]int, countColumn)
		copy(buff, table[i])
		copyTable[i] = buff
	}

	//fmt.Println(copyTable)
	resTable := make([][]int, countRow)
	for i := range resTable {
		resTable[i] = make([]int, countColumn)
	}

	var znach int

	for i := 0; i < countRow; i++ {
		znach = sliceSorted[i]

		for j := 0; j < countRow; j++ {
			if copyTable[j][click-1] == znach { //если в этой строчке copyTable лежит целевой элемент, то копируем эту строчку и
				copy(resTable[i], copyTable[j]) //вставляем в результат.
				copyTable[j][click-1] = -1      // А в строчку в copyTable пишем -1 и дальше break
				break
			}
		}

	
	}

	//fmt.Fprintln(out, resTable)
	return resTable
}
