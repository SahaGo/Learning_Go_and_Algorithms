package main

import "fmt"

// Выведите число перестановок P(n).
// В первой строке находится одно число n(1≤n≤7).

func main() {
	var input = 3

	var answer = 1

	for i := 1; i <= input; i++ {
		answer = answer * i
	}

	fmt.Println(answer)
}
