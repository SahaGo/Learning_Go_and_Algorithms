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

	// считываем количество полков и количество предстоящих вылазок
	troopsCount, research := 0, 0
	fmt.Fscan(in, &troopsCount, &research)

	// считываем число орков в i-ом полке
	trArr := make([]int, troopsCount+1)
	for i := 1; i < troopsCount+1; i++ {
		fmt.Fscan(in, &trArr[i])
	}

	// считаем префиксные суммы
	prefSum := make([]int, troopsCount+1)
	for i := 1; i < troopsCount+1; i++ {
		prefSum[i] = prefSum[i-1] + trArr[i]
	}

	// количество полков, которые должны будут отправиться в вылазку,
	// и суммарное количество орков в этих полках
	troopsForR, orks := 0, 0
	for i := 0; i < research; i++ {
		fmt.Fscan(in, &troopsForR, &orks)

		// бинпоиск
		l, r := 1, troopsCount+1-troopsForR
		for l < r {
			mid := (l + r) / 2
			if prefSum[mid+troopsForR-1]-prefSum[mid-1] >= orks {
				r = mid
			} else {
				l = mid + 1
			}
		}

		// проверяем и выводим
		if prefSum[l+troopsForR-1]-prefSum[l-1] == orks {
			fmt.Fprintln(out, l)
		} else {
			fmt.Fprintln(out, -1)
		}
	}
}
