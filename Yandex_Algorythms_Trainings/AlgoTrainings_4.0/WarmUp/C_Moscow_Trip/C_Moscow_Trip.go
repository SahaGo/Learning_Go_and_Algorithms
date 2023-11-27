package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	In  = bufio.NewReader(os.Stdin)
	Out = bufio.NewWriter(os.Stdout)
)

type dotChar struct {
	r       float64
	x       float64
	y       float64
	atan    float64
	quarter float64
}

func main() {
	defer Out.Flush()
	var a, b dotChar
	fmt.Fscan(In, &a.x, &a.y, &b.x, &b.y)

	// определяю четверть, выношу особые случаи, когда Х или У лежат на координатных прямых
	checkQuarter(&a)
	checkQuarter(&b)

	// считаем радиусы
	a.r = rO(a.x, a.y)
	b.r = rO(b.x, b.y)

	// маршрут через центр
	throughCenter := a.r + b.r

	var lowerR, diffR float64

	// определяем меньший радиус и расстояние между ними, если б они были на одном векторе
	if a.r <= b.r {
		lowerR = a.r
		diffR = b.r - a.r
	} else {
		lowerR = b.r
		diffR = a.r - b.r
	}

	//считаем угол координаты от оси оХ // АОХ и ВОХ
	atanFind(&a)
	atanFind(&b)

	//fmt.Fprintf(Out, "Четверть А = %f\nЧетверть В = %f\nr A = %f\nr B = %f\nАтангенс AОХ = %f\nАтангенс BОХ = %f\n", a.quarter, b.quarter, a.r, b.r, a.atan, b.atan)

	// вычисляю два угла, тк мы можем двитаться по окружности по часовой и против часовой стрелки
	var corner1, corner2 float64
	corner1 = math.Abs(a.atan - b.atan)
	corner2 = 360 - corner1

	//fmt.Fprintf(Out, "Угол 1 = %f\nУгол 2 = %f\n", corner1, corner2)

	//для вычисления длины дуги окружности и оптимального маршрута
	//возьмем меньший радиус, умножим на угол и домножим на П/180 чтоб перевести градусы в радианы и правильный результат
	var sDugi float64
	if corner1 <= corner2 {
		sDugi = lowerR * corner1 * math.Pi / 180
	} else {
		sDugi = lowerR * corner2 * math.Pi / 180
	}

	//найдем маршрут по дуге
	throughCirqle := sDugi + diffR

	// выберем оптимальный маршрут
	if throughCirqle <= throughCenter {
		fmt.Fprintln(Out, throughCirqle)
	} else {
		fmt.Fprintln(Out, throughCenter)
	}

	//fmt.Fprintln(Out, "duga = ", throughCirqle, "\ncenter = ", throughCenter)
}

func sBetweenXY(xA, yA, xB, yB float64) float64 {
	return math.Sqrt(math.Pow(xB-xA, 2) + math.Pow(yB-yA, 2))
}

func rO(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func checkQuarter(dot *dotChar) {
	switch {
	case dot.x >= 0 && dot.y >= 0:
		dot.quarter = 1
	case dot.x > 0 && dot.y < 0:
		dot.quarter = 4
	case dot.x < 0 && dot.y > 0:
		dot.quarter = 2
	case dot.x < 0 && dot.y < 0:
		dot.quarter = 3

	// special cases for atan func
	case dot.x == 0:
		if dot.y > 0 {
			dot.quarter = 5
		} else {
			dot.quarter = 6
		}

	case dot.y == 0:
		if dot.x > 0 {
			dot.quarter = 7
		} else {
			dot.quarter = 8
		}
	}
}

func atanFind(dot *dotChar) {
	switch dot.quarter {

	case 4:
		dot.atan = 360 - (math.Abs(math.Atan2(dot.y, dot.x)) * 180 / math.Pi)
	case 2:
		dot.atan = 180 - (math.Abs(math.Atan2(dot.y, dot.x)) * 180 / math.Pi)
	case 3:
		dot.atan = (math.Abs(math.Atan2(dot.y, dot.x)) * 180 / math.Pi) + 180
	case 1:
		dot.atan = math.Abs(math.Atan2(dot.y, dot.x)) * 180 / math.Pi
	case 5:
		dot.atan = 90
	case 6:
		dot.atan = 270
	case 7:
		dot.atan = 0
	case 8:
		dot.atan = 180
	default:
		dot.atan = math.Abs(math.Atan2(dot.y, dot.x)) * 180 / math.Pi
	}
}
