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
	sizes   []int
	parents []int
	maxSize int
}

func main() {
	defer Out.Flush()

	var (
		n int // число таблиц
		m int // число запросов
	)

	fmt.Fscan(In, &n, &m)

	data := Data{
		sizes:   make([]int, n),
		parents: make([]int, n),
	}

	for i := 0; i < n; i++ {
		buff := 0
		fmt.Fscan(In, &buff)

		data.sizes[i] = buff
		data.parents[i] = i
		if buff > data.maxSize {
			data.maxSize = buff
		}
	}

	for i := 0; i < m; i++ {
		var (
			dest   int
			source int
		)
		fmt.Fscan(In, &dest, &source)
		fmt.Fprintln(Out, data.Union(dest-1, source-1))
	}
}

// Find находим корень и делаем сжатие путей
func (d *Data) Find(i int) int {
	if i != d.parents[i] {
		d.parents[i] = d.Find(d.parents[i])
	}
	return d.parents[i]
}

// Union объединяем таблицы в соответсвии с условием задачи
func (d *Data) Union(i, j int) int {
	iId := d.Find(i) //destination id
	jId := d.Find(j) // source id

	if iId != jId {
		d.sizes[iId] += d.sizes[jId]
		d.sizes[jId] = 0

		if d.sizes[iId] > d.maxSize {
			d.maxSize = d.sizes[iId]
		}

		d.parents[jId] = iId
	}

	return d.maxSize
}
