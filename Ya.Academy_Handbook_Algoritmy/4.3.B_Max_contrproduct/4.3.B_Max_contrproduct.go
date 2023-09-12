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
	var a int

	fmt.Fscan(in, &a)
	arr := make([]int, a)

	if a <= 9 {
		fmt.Fprintln(out, "No")
		return
	}
	for i := 1; i < a; i++ {
		arr[i] = i
	}

	fmt.Fprintln(out, "Yes")

	arr[0] = a
	for i := 0; i < a; i++ {
		fmt.Fprint(out, arr[i], " ")
	}
}
