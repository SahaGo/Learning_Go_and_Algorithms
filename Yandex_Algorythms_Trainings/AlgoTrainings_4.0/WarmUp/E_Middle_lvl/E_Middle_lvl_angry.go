// time limit
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	In  = bufio.NewReader(os.Stdin)
	Out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer Out.Flush()

	var count int
	fmt.Fscan(In, &count)
	rang := make([]float64, count)
	for i := 0; i < count; i++ {
		fmt.Fscan(In, &rang[i])
	}

	for _, v := range rang {
		fmt.Fprint(Out, processing(v, rang), " ")
	}

}

func processing(stRang float64, rang []float64) float64 {
	var res float64
	for i := 0; i < len(rang); i++ {
		res += math.Abs(rang[i] - stRang)
	}
	return res
}
