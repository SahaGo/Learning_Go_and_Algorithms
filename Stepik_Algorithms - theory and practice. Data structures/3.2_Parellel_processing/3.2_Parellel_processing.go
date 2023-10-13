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

type Data struct {
	procTime int
	proc     int
}

func main() {
	defer Out.Flush()

	var (
		n int
		m int
	)

	fmt.Fscan(In, &n, &m)

	heap := make([]Data, n)

	for i := 0; i < n; i++ {
		heap[i].proc = i
	}

	for i := 0; i < m; i++ {
		time := 0
		fmt.Fscan(In, &time)

		res1, res2 := ExtractMin(heap)
		fmt.Fprintln(Out, res1, res2)

		heap[0].procTime += time

		SiftDown(0, heap)
	}
}

// SiftDown просеиваем вниз
func SiftDown(i int, h []Data) {
	minIndex := i

	l := LeftChild(i)
	if l < len(h) && (h[l].procTime < h[minIndex].procTime ||
		(h[l].procTime == h[minIndex].procTime && h[l].proc < h[minIndex].proc)) {
		minIndex = l
	}

	r := RightChild(i)
	if r < len(h) && (h[r].procTime < h[minIndex].procTime ||
		(h[r].procTime == h[minIndex].procTime && h[r].proc < h[minIndex].proc)) {
		minIndex = r
	}

	if i != minIndex {

		h[i], h[minIndex] = h[minIndex], h[i]

		SiftDown(minIndex, h)
	}
}

// LeftChild находим левого ребенка
func LeftChild(i int) int {
	return 2*i + 1
}

// RightChild находим индекс правого ребенка
func RightChild(i int) int {
	return 2*i + 2
}

// ExtractMin возвращаем значение вершины, т.е. наш минимум
func ExtractMin(h []Data) (int, int) {
	return h[0].proc, h[0].procTime
}
