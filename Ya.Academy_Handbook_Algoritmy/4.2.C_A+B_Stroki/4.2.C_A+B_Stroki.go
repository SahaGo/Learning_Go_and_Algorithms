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
	var leeen int
	var b string

	fmt.Fscan(in, &leeen)
	fmt.Fscan(in, &b)
	arrA := []rune(b)

	fmt.Fscan(in, &leeen)
	fmt.Fscan(in, &b)
	arrB := []rune(b)

	for i := 0; i < len(arrA); i++ {
		fmt.Fprintf(out, "%c%c", arrA[i], arrB[i])
	}
}
