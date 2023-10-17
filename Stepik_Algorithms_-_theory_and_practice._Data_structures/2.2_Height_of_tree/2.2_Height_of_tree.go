package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count int
	fmt.Fscan(in, &count)

	arr := make(map[int][]int, count)

	head := math.MinInt

	for i := 0; i < count; i++ {
		buff := 0
		fmt.Fscan(in, &buff)
		if buff == -1 {
			head = i
		}
		arr[buff] = append(arr[buff], i)
	}

	//	height := 1 // высота дерева, состоящей только из вершины r

	//for k, v := range arr {
	//	if v == -1 {
	//		head = k
	//	}
	//}

	total := height(head, arr)

	fmt.Fprintln(out, arr, head, total)
}

func height(h int, arr map[int][]int) int {
	total := 1
	if arr[h] == nil {
		return total
	}
	for _, v := range arr[h] {
		subHeight := height(v, arr) + 1
		//total = int(math.Max(float64(total), float64(subHeight)))
		if subHeight > total {
			total = subHeight
		}
	}
	return total
}
