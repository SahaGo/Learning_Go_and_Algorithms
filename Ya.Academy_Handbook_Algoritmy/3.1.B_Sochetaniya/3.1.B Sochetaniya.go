package main

import "fmt"

//Выведите число сочетаний C(n,k).
//В первой строке находится два числа n(1≤n≤7), k(1≤k≤7).

func main() {
	//var n uint = 3 // result 3
	//var k uint = 2

	//var n uint = 7 // result 21
	//var k uint = 5

	var n uint = 1 // result 1
	var k uint = 1

	var answer uint = 1

	if k == 0 {
		fmt.Println(answer)
		return
	}

	if k == 1 {
		fmt.Println(n)
		return
	}

	//if k == 2 {
	//	answer = n * ((n - 1) / 2)
	//	fmt.Println(answer)
	//	return
	//}

	var numerator uint = 1
	var denominator1 uint = 1
	var denominator2 uint = 1

	var i uint
	for i = 1; i <= n; i++ {
		numerator = numerator * i
	}

	for i = 1; i <= n-k; i++ {
		denominator1 = denominator1 * i
	}

	for i = 1; i <= k; i++ {
		denominator2 = denominator2 * i
	}

	answer = numerator / (denominator1 * denominator2)
	fmt.Println(answer)
}
