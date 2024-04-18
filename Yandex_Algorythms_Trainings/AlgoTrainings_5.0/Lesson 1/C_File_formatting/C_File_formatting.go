package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var linesCount int
	fmt.Fscanln(in, &linesCount)

	var num int
	ans := 0
	for i := 0; i < linesCount; i++ {
		fmt.Fscanln(in, &num)
		//ans += num / 4

		// мое решение не прошло по памяти
		//remains := num % 4
		//switch remains {
		//case 1:
		//	ans += 1
		//case 2, 3:
		//	ans += 2
		//}

		// решение разбора тоже не прошло по памяти!!!
		//if num%4 == 1 || num%4 == 2 {
		//	ans += num % 4
		//}
		//if num%4 == 3 {
		//	ans += 2
		//}

		//РЕШЕНИЯ НЕ ПРОХОДИЛИ ИЗ-ЗА СТАНДАРТНОГО ВВОДА-ВЫВОДА, ПОМОГ БУФЕРИЗИРОВАННЫЙ
		ans = ans + num/4 + int(math.Min(float64(num%4), 2))
	}

	fmt.Fprintln(out, ans)
}
