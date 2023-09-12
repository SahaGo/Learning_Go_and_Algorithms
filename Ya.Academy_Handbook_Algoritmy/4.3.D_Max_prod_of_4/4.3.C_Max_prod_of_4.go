// не проходит на тестах со всеми отрицательными числами

package main

import (
	"bufio"
	"fmt"
	"math"
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

	if len(arr) == 4 {
		fmt.Fprintln(out, arr, "\n", arr[0]*arr[1]*arr[2]*arr[3])
		return
	}

	sort.Slice(arr, func(i int, j int) bool {
		return arr[i] >= arr[j]
	})

	pretendersForMaxRes := []int{}

	pretendersForMaxRes = append(pretendersForMaxRes, arr[0]*arr[1]*arr[2]*arr[3])

	if secondCase(arr) {
		pretendersForMaxRes = append(pretendersForMaxRes, arr[0]*arr[1]*arr[len(arr)-1]*arr[len(arr)-2])
	}

	if thirdCase(arr) {
		pretendersForMaxRes = append(pretendersForMaxRes, arr[len(arr)-3]*arr[len(arr)-4]*arr[len(arr)-1]*arr[len(arr)-2])
	}

	max := math.MinInt
	for _, v := range pretendersForMaxRes {
		if v > max {
			max = v
		}
	}
	fmt.Fprintln(out, arr, "\n", pretendersForMaxRes, max)

}

func thirdCase(arr []int) bool {
	if arr[len(arr)-1] > 0 && arr[len(arr)-2] > 0 && arr[len(arr)-3] > 0 && arr[len(arr)-4] > 0 {
		return false
	}

	return true
}

func secondCase(arr []int) bool {
	if arr[0] < 0 && arr[1] < 0 && arr[len(arr)-1] > 0 && arr[len(arr)-2] > 0 {
		return false
	}

	return true
}
