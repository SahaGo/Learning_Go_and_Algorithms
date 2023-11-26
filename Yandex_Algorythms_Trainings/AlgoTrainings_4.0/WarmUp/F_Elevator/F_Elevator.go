package main

import (
	"bufio"
	"os"
)

var (
	In  = bufio.NewReader(os.Stdin)
	Out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer Out.Flush()

}
