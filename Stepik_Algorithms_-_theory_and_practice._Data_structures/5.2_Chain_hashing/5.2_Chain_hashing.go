package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	p int64   = 1000000007
	x float64 = 263
)

var (
	Scan = bufio.NewScanner(os.Stdin)
	Out  = bufio.NewWriter(os.Stdout)
)

type Dictionary struct {
	record map[string][]string
}

func (d *Dictionary) AddString(str string, m int64) {
	hash := HashFunc(str, m)
	if v, ok := d.record[hash]; ok {
		for _, val := range v { // если повтор, то игнорим
			if val == str {
				return
			}
		}
		d.record[hash] = append([]string{str}, d.record[hash]...) //если слайс не пустой, то добавляем в начало
		return
	}
	// если нет, то создаем его
	d.record[hash] = make([]string, 0)
	d.record[hash] = append(d.record[hash], str)
	return
}

func (d *Dictionary) DeleteString(str string, m int64) {
	hash := HashFunc(str, m)

	if v, ok := d.record[hash]; ok {

		if len(d.record[hash]) == 1 { //если у нас только один элемент по хеш-ключу, то удаляем его
			d.record[hash] = make([]string, 0)
			return
		} else { // иначе перебираем слайс
			var ind int               // ind - индекс удаляемого элемента
			for key, val := range v { //поиск элемента в слайсе
				if val == str {
					ind = key
				}
			}
			d.record[hash] = append(d.record[hash][:ind], d.record[hash][ind+1:]...) // удаляем элемент по индексу ind
			return
		}
	}
	return
}

func (d *Dictionary) FindString(str string, m int64) {
	hash := HashFunc(str, m)

	if v, ok := d.record[hash]; ok {
		for _, val := range v {
			if val == str {
				fmt.Fprintln(Out, "yes")
				return
			}
		}
	}
	fmt.Fprintln(Out, "\nno")
	return
}

func (d *Dictionary) CheckI(i string) {
	if v, ok := d.record[i]; ok {
		if len(d.record[i]) == 0 {
			fmt.Fprint(Out, "\n")
		}
		for k, val := range v {
			if k == 0 {
				fmt.Fprint(Out, "\n")
			}
			fmt.Fprint(Out, val, " ")
		}
	}
}

func CreateDictionary() Dictionary {
	d := Dictionary{}
	d.record = make(map[string][]string)
	return d
}

func HashFunc(str string, m int64) string {
	var hashSum int64
	for i, v := range str {
		hashSum = hashSum + (int64(v)*int64(math.Pow(x, float64(i))))%p
		//fmt.Fprintln(Out, "\nhashing: v = ", v, ", x = ", x, ", i = ", i)
	}
	hashSum = hashSum % p % m

	return strconv.Itoa(int(hashSum))
}

type Methods interface {
	AddString(string, int64)
	DeleteString(string, int64)
	FindString(string, int64)
	CheckI(string)
}

func main() {
	defer Out.Flush()
	Scan.Scan()
	mInt, _ := strconv.Atoi(Scan.Text())
	m := int64(mInt)
	Scan.Scan()
	request, _ := strconv.Atoi(Scan.Text())

	dict := CreateDictionary()

	for i := 0; i < request; i++ {
		Scan.Scan()
		buff := strings.Split(Scan.Text(), " ")

		str := buff[1]

		switch buff[0] {
		case "add":
			dict.AddString(str, m)
		case "del":
			dict.DeleteString(str, m)
		case "check":
			dict.CheckI(str)
		case "find":
			dict.FindString(str, m)
		}
		//fmt.Fprintln(Out, "\nTest - ", i+1, ": ", buff[0], dict)
	}
}
