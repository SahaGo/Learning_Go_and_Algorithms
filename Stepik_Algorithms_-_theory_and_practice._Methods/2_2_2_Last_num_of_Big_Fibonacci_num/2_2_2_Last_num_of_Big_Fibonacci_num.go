package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	arr := make([]int, x+1)
	arr[1] = 1
	for i := 2; i <= x; i++ {
		arr[i] = (arr[i-2] + arr[i-1]) % 10
	}
	fmt.Print(arr[x])
}
