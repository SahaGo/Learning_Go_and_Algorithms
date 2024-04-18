package main

import (
	"fmt"
)

//uint8 == byte -> true !!!

const (
	empty    byte = 0
	occupied byte = 1
	r        byte = 'R'
	b        byte = 'B'
)

type char struct {
	typ uint8
	i   int
	j   int
}

type coord struct {
	i int
	j int
}

type square struct {
	status byte
	value  uint8
}

func main() {
	table, coordinates := readChessTable()

	var res int
	for i := 0; i < len(coordinates); i++ {
		if coordinates[i].typ == b {
			B(coordinates[i].i, coordinates[i].j, table, &res)
		} else {
			R(coordinates[i].i, coordinates[i].j, table, &res)
		}
	}

	fmt.Println(64 - res)
}

// Делаю матрицу с границами
func readChessTable() ([][]square, []char) {
	table := make([][]square, 10)
	coordinates := make([]char, 0)
	table[0] = make([]square, 10) // граница
	table[9] = make([]square, 10) // граница

	for i := 1; i < 9; i++ {
		table[i] = make([]square, 10) // инициализирую строку поля
		var str string
		fmt.Scanln(&str)

		for j := 0; j < 8; j++ {
			if str[j] == 'R' || str[j] == 'B' {
				coordinates = append(coordinates, char{typ: str[j], i: i, j: j + 1})
			}
			table[i][j+1] = square{value: str[j]}
		}
	}
	//fmt.Println(table)
	return table, coordinates
}

// R Ладья вверх, вниз
func R(line int, column int, table [][]square, res *int) {
	table[line][column].status = occupied // статус начального расположения фигуры - занято
	*res++

	//up
	for i := line - 1; i != 0; i-- { // идем до границ
		if table[i][column].value == b || table[i][column].value == r { // если фигура, то дальше идти не можем
			break
		}
		if table[i][column].status == occupied { // если уже был ход в этой клетке, то пропускаем
			continue
		}
		table[i][column].status = occupied
		*res++
	}

	//down
	for i := line + 1; i != 9; i++ {
		if table[i][column].value == b || table[i][column].value == r {
			break
		}

		if table[i][column].status == occupied { // если уже был ход в этой клетке, то пропускаем
			continue
		}
		table[i][column].status = occupied
		*res++
	}

	//left
	for j := column - 1; j != 0; j-- {
		if table[line][j].value == b || table[line][j].value == r {
			break
		}

		if table[line][j].status == occupied { // если уже был ход в этой клетке, то пропускаем
			continue
		}

		table[line][j].status = occupied
		*res++
	}

	//right
	for j := column + 1; j != 9; j++ {
		if table[line][j].value == b || table[line][j].value == r {
			break
		}
		if table[line][j].status == occupied { // если уже был ход в этой клетке, то пропускаем
			continue
		}
		table[line][j].status = occupied
		*res++
	}
}

// B слон по диагонали
func B(line int, column int, table [][]square, res *int) {
	table[line][column].status = occupied // статус начального расположения фигуры - занято
	*res++

	//left-up
	i, j := line, column
	for {
		i--
		j--

		if i == 0 || j == 0 { // если дошли до границы, выходим из цикла
			break
		}
		if table[i][j].value == b || table[i][j].value == r { // если встретили фигуру, останавливаем ходы
			break
		}
		if table[i][j].status == occupied { // если уже был ход в этой клетке, то пропускаем
			continue
		}

		table[i][j].status = occupied
		*res++
	}

	i, j = line, column
	//left-down
	for {
		i++
		j--

		if i == 9 || j == 0 {
			break
		}
		if table[i][j].value == b || table[i][j].value == r {
			break
		}
		if table[i][j].status == occupied { // если уже был ход в этой клетке, то пропускаем
			continue
		}

		table[i][j].status = occupied
		*res++
	}

	//right-up
	i, j = line, column
	for {
		i--
		j++

		if i == 0 || j == 9 {
			break
		}
		if table[i][j].value == b || table[i][j].value == r {
			break
		}
		if table[i][j].status == occupied { // если уже был ход в этой клетке, то пропускаем
			continue
		}

		table[i][j].status = occupied
		*res++
	}

	//right-down
	i, j = line, column
	for {
		i++
		j++

		if i == 9 || j == 9 {
			break
		}
		if table[i][j].value == b || table[i][j].value == r {
			break
		}
		if table[i][j].status == occupied { // если уже был ход в этой клетке, то пропускаем
			continue
		}

		table[i][j].status = occupied
		*res++
	}
}
