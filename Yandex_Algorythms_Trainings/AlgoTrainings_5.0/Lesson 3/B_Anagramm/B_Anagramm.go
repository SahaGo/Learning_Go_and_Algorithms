package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	In  = bufio.NewReader(os.Stdin)
	Out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer Out.Flush()

	var first, second string
	fmt.Fscan(In, &first, &second)

	runes := make(map[rune]int)

	for _, v := range first {
		runes[v]++
	}

	for _, v := range second {
		if _, ok := runes[v]; ok {
			if runes[v] == 1 {
				delete(runes, v)
				continue
			}
		}
		runes[v]--
	}

	if len(runes) == 0 {
		fmt.Fprint(Out, "YES")
	} else {
		fmt.Fprint(Out, "NO")
	}

}
