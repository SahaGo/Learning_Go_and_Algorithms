package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	count := 0
	fmt.Fscanf(in, "%d\n", &count)

	nums := make(map[int]int)
	for i := 0; i < count; i++ {
		x := 0
		fmt.Fscan(in, &x)
		nums[x]++
	}

	maxRes := 0
	for k, _ := range nums {

		now := nums[k] + max(nums[k+1], nums[k-1])
		maxRes = max(maxRes, now)
	}
	fmt.Fprintln(out, count-maxRes)
}
