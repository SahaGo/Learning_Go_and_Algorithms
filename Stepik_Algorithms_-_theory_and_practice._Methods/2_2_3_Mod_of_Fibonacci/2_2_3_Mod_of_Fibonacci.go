package main

import "fmt"

func main() {
	var x, m int64
	fmt.Scan(&x, &m)
	arr := []int64{0, 1}

	for {

		l := len(arr)
		temp := (arr[l-2] + arr[l-1]) % m
		arr = append(arr, temp)

		l = len(arr)
		if arr[l-2]%m == 0 && arr[l-1]%m == 1 {
			break
		}
	}
	var l int64 = int64(len(arr))
	fmt.Print(arr[x%(l-2)])
}
