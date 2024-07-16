package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	res := make(map[int]int) //[число]номер массива
	ans := make([]int, 0)
	store := make(map[int]int) //[число]номер массива

	for i := 0; i < 3; i++ {

		l := 0
		fmt.Fscan(in, &l)
		for j := 0; j < l; j++ {
			x := 0
			fmt.Fscan(in, &x)

			_, inStore := store[x]
			_, inRes := res[x]

			if !inStore { // add if first meet
				store[x] = i
			}

			//что, если это повторы в одном и том же массиве? проверим номер массива
			if inStore && !inRes && (store[x] != i) { // more meet and not written
				ans = append(ans, x)
				res[x] = i
			}
		}
	}
	sort.Ints(ans)
	s := strings.Trim(fmt.Sprint(ans), "[]")
	fmt.Fprint(out, s)
}
