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

	var nStr string
	var k int64
	var d int
	fmt.Fscanf(in, "%s %d %d", &nStr, &k, &d)

	// остаток от деления
	m := M(nStr, k)
	l := len(nStr)

	nRune := make([]rune, l, l+d)
	for key, v := range nStr {
		nRune[key] = v
	}

	var remains int64
	for i := 0; i < d; i++ { //проходимся по дням
		remains = (m*10/k*k + k - m*10) % k //искомый остаток

		if remains > 9 {
			fmt.Fprintln(out, -1)
			return
		}

		nRune = append(nRune, rune(remains)+'0')
		m = (m*10 + remains) % k
	}

	nStr = string(nRune)
	fmt.Fprintln(out, nStr)
}

func M(nStr string, k int64) int64 {
	var m int64
	for _, vRune := range nStr {
		v := int64(vRune - '0') // получаем из руны число грязным способом

		m = (m*10 + v) % k
	}
	return m
}
