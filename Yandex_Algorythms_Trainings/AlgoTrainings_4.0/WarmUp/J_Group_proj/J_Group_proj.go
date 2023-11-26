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

	var test, studentsCount, aFrom, bTo int

	for i := 0; i < test; i++ {
		fmt.Fscanf(In, "%d %d %d", &studentsCount, &aFrom, &bTo)

	}
}
