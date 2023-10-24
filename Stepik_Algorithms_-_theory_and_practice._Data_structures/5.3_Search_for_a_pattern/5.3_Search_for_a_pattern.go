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

const (
	p = 1000000007
	x = 263
)

func main() {
	defer Out.Flush()

	var (
		pattern string
		text    string
	)

	fmt.Fscanf(In, "%s\n%s", &pattern, &text)

	sizePatt := len(pattern)
	sizeText := len(text)

	powX := make([]int, sizePatt) // создаем слайс со значениями степени Х
	powX[0] = 1
	for i := 1; i < sizePatt; i++ { // заполняем
		powX[i] = powX[i-1] * x % p
	}

	res := make([]int, 0)

	pattHash := Hashing(pattern, powX) // хеш-значение образца

	nextInd := sizeText - sizePatt // индекс элемента, следующего за размером окна образца
	hash := Hashing(text[nextInd:], powX)

	if hash == pattHash { // если совпадение, записываем в результат
		res = append(res, nextInd)
	}

	for i := nextInd - 1; i >= 0; i-- { // обходим с конца текста
		//вычитаем последний символ в соответствующей степени, и добавляем новый
		hash = ((hash-int(text[i+sizePatt])*powX[sizePatt-1]%p+p)%p*x%p + int(text[i])) % p
		if hash == pattHash {
			res = append(res, i)
		}
	}

	for i := len(res) - 1; i >= 0; i-- { // выводим в обратном порядке, тк обходили текст с конца
		fmt.Fprint(Out, res[i], " ")
	}
	//fmt.Fprintln(Out, pattern, text)
}

func Hashing(str string, pows []int) int {
	hashSum := 0
	for k, v := range str {
		hashSum = ((hashSum + (int(v) * pows[k])) % p) % p
	}
	return hashSum
}
