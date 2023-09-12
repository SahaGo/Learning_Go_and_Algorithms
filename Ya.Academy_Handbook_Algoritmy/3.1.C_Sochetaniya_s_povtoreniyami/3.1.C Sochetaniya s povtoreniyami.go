package main

import "fmt"

//Выведите число сочетаний с повторением C‾ (n,k).
//В первой строке находятся два числа n(1≤n≤4), k(1≤k≤4).
func main() {
	var n uint = 1 // result 1
	var k uint = 1

	//var n uint = 4 // result 20
	//var k uint = 3

	//var n uint = 2 // result 3
	//var k uint = 2

	var answer uint = 1

	if k == 0 {
		fmt.Println(answer)
		return
	}

	if k == 1 {
		fmt.Println(n)
		return
	}

	var numerator uint = 1
	var denominator1 uint = 1
	var denominator2 uint = 1

	var i uint
	for i = 1; i <= n+k-1; i++ {
		numerator = numerator * i
	}

	for i = 1; i <= n-1; i++ {
		denominator1 = denominator1 * i
	}

	for i = 1; i <= k; i++ {
		denominator2 = denominator2 * i
	}

	answer = numerator / (denominator1 * denominator2)
	fmt.Println(answer)
}
