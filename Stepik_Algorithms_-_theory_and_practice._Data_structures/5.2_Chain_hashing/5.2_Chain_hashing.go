// не проходит 7 тест, все примеры ввода отображаются коректно,
// вариант решения - переписать со структурой[][]string
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	p int64 = 1000000007
	x int64 = 263
)

var (
	Scan = bufio.NewScanner(os.Stdin)
	Out  = bufio.NewWriter(os.Stdout)
)

type Dictionary struct {
	record map[int64][]string
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
			//d.record[hash] = make([]string, 0)
			delete(d.record, hash)
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

func (d *Dictionary) FindString(str string, m int64) bool {
	hash := HashFunc(str, m)

	if v, ok := d.record[hash]; ok {
		for _, val := range v {
			if val == str {
				return true
			}
		}
	}
	return false
}

func (d *Dictionary) CheckI(i int64) []string {
	if v, ok := d.record[i]; ok {
		return v
	}
	return nil
}

func CreateDictionary() Dictionary {
	d := Dictionary{}
	d.record = make(map[int64][]string)
	return d
}

func HashFunc(str string, m int64) int64 {
	var hashSum int64

	var pow int64 = 1
	for _, v := range str {

		hashSum = hashSum + (int64(v) * pow)
		hashSum = (hashSum%p + p) % p
		pow = (pow * x) % p
	}
	hashSum = hashSum % p % m

	return hashSum
}

type Methods interface {
	AddString(string, int64)
	DeleteString(string, int64)
	FindString(string, int64) bool
	CheckI(int64) []string
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
			strConverted, _ := strconv.ParseInt(str, 10, 64)
			res := dict.CheckI(strConverted)

			fmt.Fprintln(Out, strings.Join(res, " "))

		case "find":
			if dict.FindString(str, m) {
				fmt.Fprintln(Out, "yes")
			} else {
				fmt.Fprintln(Out, "no")
			}
		}
	}
}
