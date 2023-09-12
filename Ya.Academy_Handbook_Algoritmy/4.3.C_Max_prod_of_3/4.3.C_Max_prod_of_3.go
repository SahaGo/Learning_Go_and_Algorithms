package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var a, b int

	fmt.Fscan(in, &a)
	arr := make([]int, a)
	for i := 0; i < a; i++ {
		fmt.Fscan(in, &b)
		arr[i] = b
	}

	sort.Slice(arr, func(i int, j int) bool {
		return arr[i] >= arr[j]
	})

	fmt.Fprintln(out, arr[0]*arr[1]*arr[2])
}
