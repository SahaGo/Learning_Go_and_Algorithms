package main

import (
	"fmt"
	"math"
)

func main() {
	var troops, HP, enemiesProduct int

	fmt.Scan(&troops, &HP, &enemiesProduct)
	ans := math.MaxInt

	// моделирование

	// до какого уровня здоровья казармы мы будем
	// нападать на нее первым способом
	for t := 0; t < HP+1; t++ {
		ans = min(ans, calc(t, troops, HP, enemiesProduct))
	}
	if ans != math.MaxInt {
		fmt.Print(ans)
	} else {
		fmt.Print(-1)
	}
}

func calc(t, myTroops, barrHP, enemiesProduct int) int {
	rounds := 0
	enemyTroops := 0
	for barrHP >= t { // пока не достигли предела, до которого мы стукаем казарму
		if enemyTroops >= myTroops { // если противника больше, ьо не выигрываем
			return math.MaxInt
		}
		barrHP -= myTroops - enemyTroops // добиваем противников и бьем казарму
		enemyTroops = 0
		if barrHP >= 0 { // уцелевшая казарма производит впротивников
			enemyTroops += enemiesProduct
		}
		rounds++
	}

	// стратегия - сломать казарм, бросив на это все силы
	for barrHP > 0 {
		if myTroops <= 0 { // если нашего войска не осталось, то проигрыш
			return math.MaxInt
		}
		if barrHP >= myTroops { // если войска есть, то все на разрушение казармы
			barrHP -= myTroops
		} else { //если наших больше, чем хп казармы, то дробиваем казарму и бьем противников, если естались
			enemyTroops -= myTroops - barrHP
			barrHP = 0
		}
		myTroops -= enemyTroops

		//если казарма еще жива, то в конце хода она производит противников
		if barrHP > 0 {
			enemyTroops += enemiesProduct
		}
		rounds++
	}

	// добиваем армию противника
	for enemyTroops > 0 {
		if myTroops <= 0 {
			return math.MaxInt
		}
		enemyTroops -= myTroops
		if enemyTroops > 0 {
			myTroops -= enemyTroops
		}
		rounds++
	}
	return rounds
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
