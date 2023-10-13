package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	In  = bufio.NewReader(os.Stdin)
	Out = bufio.NewWriter(os.Stdout)
	Res []int
)

func main() {
	defer Out.Flush()

	var count int

	fmt.Fscan(In, &count)

	arr := make([]int, count)

	for i := 0; i < count; i++ {
		var scan int
		fmt.Fscan(In, &scan)
		arr[i] = scan
	}

	BuildHeap(arr)

	fmt.Fprintln(Out, len(Res)/2)
	for i := 0; i < len(Res); {
		fmt.Fprintln(Out, Res[i], Res[i+1])
		i += 2
	}
}

func BuildHeap(arr []int) {
	size := len(arr)

	for i := size / 2; i >= 0; i-- {
		SiftDown(i, arr)
	}

}

func SiftDown(i int, arr []int) {
	minIndex := i

	leftChildIndex := LeftChildIndex(i)
	if leftChildIndex < len(arr) && arr[leftChildIndex] < arr[minIndex] {
		minIndex = leftChildIndex
	}

	rightChildIndex := RightChildIndex(i)
	if rightChildIndex < len(arr) && arr[rightChildIndex] < arr[minIndex] {
		minIndex = rightChildIndex
	}

	if i != minIndex {
		Res = append(Res, i, minIndex)
		//swap
		arr[i], arr[minIndex] = arr[minIndex], arr[i]

		SiftDown(minIndex, arr)
	}
}

func LeftChildIndex(i int) int {
	return 2*i + 1
}

func RightChildIndex(i int) int {
	return 2*i + 2
}
