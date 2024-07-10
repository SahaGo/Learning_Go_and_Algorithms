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
	var n, k int
	fmt.Fscan(in, &n, &k)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &nums[i])
	}

	if sol(nums, k) {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}

func sol(nums []int, k int) bool {
	lastK := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := lastK[nums[i]]; ok {
			return true
		}
		lastK[nums[i]] = true
		if i >= k { // как только набрали окно, ток тогда удаляем просмотренный элемент
			delete(lastK, nums[i-k])
		}
	}
	return false
}
