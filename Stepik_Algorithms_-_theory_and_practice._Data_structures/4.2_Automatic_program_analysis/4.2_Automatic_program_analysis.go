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
	arr  []int
	rank []int
}

func main() {
	defer Out.Flush()

	var (
		n int // число переменных
		e int // число равенств х=х
		d int // число неравенств х!=х
	)

	fmt.Fscan(In, &n, &e, &d)

	data := Data{
		arr:  make([]int, n),
		rank: make([]int, n),
	}

	data.MakeSet()
	x1, x2 := 0, 0

	for i := 0; i < e; i++ {
		fmt.Fscan(In, &x1, &x2)
		data.Union(x1-1, x2-1)
	}

	var res int = 1

	for i := 0; i < d; i++ {
		fmt.Fscan(In, &x1, &x2)
		if data.Find(x1-1) == data.Find(x2-1) {
			res = 0
			break
		}
	}

	fmt.Fprintln(Out, res)
}

// MakeSet создание множества
func (d *Data) MakeSet() {
	for i := 0; i < len(d.arr); i++ {
		d.arr[i] = i
	}
}

// Find нахождение множества
func (d *Data) Find(i int) int {
	for i != d.arr[i] {
		i = d.arr[i]
	}
	return i
}

// Union объединение множеств
func (d *Data) Union(i, j int) {
	iId := d.Find(i)
	jId := d.Find(j)

	if iId == jId {
		return
	}

	if d.rank[iId] > d.rank[jId] {
		d.arr[jId] = iId
	} else {
		d.arr[iId] = jId
		if d.rank[iId] == d.rank[jId] {
			d.rank[jId] += 1
		}
	}
}
