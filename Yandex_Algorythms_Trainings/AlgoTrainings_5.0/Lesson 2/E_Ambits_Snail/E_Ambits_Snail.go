package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type data struct {
	up   int64
	down int64
	i    int
}

var (
	In  = bufio.NewReader(os.Stdin)
	Out = bufio.NewWriter(os.Stdout)
)

func main() {

	defer Out.Flush()

	var count int
	var up, down int64
	fmt.Fscanln(In, &count)
	good := make([]data, 0, count)
	bad := make([]data, 0, count)

	indUpBad := -1
	//maxIndBgood := -1
	var maxGoodDown, maxBadUp int64
	for i := 0; i < count; i++ {
		fmt.Fscanln(In, &up, &down)
		if up-down < 0 {

			bad = append(bad, data{up: up, down: down, i: i + 1})

			if up > maxBadUp {
				maxBadUp = up
				indUpBad = len(bad) - 1
			}

		} else {
			good = append(good, data{up: up, down: down, i: i + 1})

			if down > maxGoodDown {
				//maxIndBgood = len(good) - 1
				maxGoodDown = down
			}
		}
	}

	res1, sum, lastSortedGoodDown := variant1(good, bad, count)

	res2, sumWithBadUp := variant2(good, bad, sum, lastSortedGoodDown, count, indUpBad)

	if sum > sumWithBadUp {
		printAnswer(sum, res1)
	} else {
		printAnswer(sumWithBadUp, res2)
	}

	//	fmt.Println("\ntest1", sum, sumWithBadUp)
	//	fmt.Println("test2", sum, "\ngood:", good, "\nbad:", bad, "\nres1:", res1, "\nres2:", res2)

}

func printAnswer(sum int64, res []data) {

	fmt.Fprintln(Out, sum)

	for i := 0; i < len(res); i++ {
		if i == len(res)-1 {
			fmt.Fprint(Out, res[i].i)
		} else {
			fmt.Fprint(Out, res[i].i, " ")
		}
	}
}

func variant1(good, bad []data, count int) ([]data, int64, int64) {
	if len(good) == 0 {
		return nil, 0, 0
	}
	res1 := make([]data, 0, count/2)
	var sum int64

	res1 = append(res1, good...)
	sort.Slice(res1, func(i, j int) bool {
		if res1[i].down <= res1[j].down {
			return true
		}
		return false
	})
	for i := 0; i < len(res1); i++ {
		sum += res1[i].up
		sum -= res1[i].down
	}

	sum += res1[len(res1)-1].down

	lastDown := res1[len(res1)-1].down
	res1 = append(res1, bad...)

	return res1, sum, lastDown
}

func variant2(good, bad []data, sum int64, lastSortedGoodDown int64, count int, indUpBad int) ([]data, int64) {
	if len(bad) == 0 {
		return bad, -1
	}
	sumFor2Variant := sum
	res2 := make([]data, 0, count/2)
	res2 = append(res2, good...)

	res2 = append(res2, bad[indUpBad])

	if len(bad) != 0 {
		sumFor2Variant += bad[indUpBad].up
	}
	sumFor2Variant -= lastSortedGoodDown
	bad = append(bad[:indUpBad], bad[indUpBad+1:]...)
	res2 = append(res2, bad...)
	return res2, sumFor2Variant
}
