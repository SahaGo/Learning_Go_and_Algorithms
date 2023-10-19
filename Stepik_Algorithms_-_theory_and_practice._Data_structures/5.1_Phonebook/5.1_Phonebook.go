package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PhoneBook struct {
	records map[string]string
}

// AddNumberName добавление в записную книжку
func (pb *PhoneBook) AddNumberName(num string, name string) {
	num = pb.HashFunc(num)
	pb.records[num] = name
}

// DeleteNumber удалить номер
func (pb *PhoneBook) DeleteNumber(num string) {
	num = pb.HashFunc(num)
	delete(pb.records, num)
}

// FindNumber найти номер
func (pb *PhoneBook) FindNumber(num string) {
	num = pb.HashFunc(num)
	if v, ok := pb.records[num]; ok {
		fmt.Fprintln(Out, v)
		return
	}
	fmt.Fprintln(Out, "not found")
	return
}

// HashFunc хешировать по алгоритму МД5
func (pb *PhoneBook) HashFunc(num string) string {
	key := []byte(num)
	return fmt.Sprintf("%x", md5.Sum(key))
}

type Methods interface {
	HashFunc(string) string
	AddNumberName(string, string)
	DeleteNumber(string)
	FindNumber(string)
}

var (
	Scan = bufio.NewScanner(os.Stdin)
	Out  = bufio.NewWriter(os.Stdout)
)

func main() {
	defer Out.Flush()
	Scan.Scan()
	request, _ := strconv.Atoi(Scan.Text())

	phoneBook := CreatePhoneBook()

	for i := 0; i < request; i++ {

		Scan.Scan()
		buff := strings.Split(Scan.Text(), " ")

		num := buff[1]

		switch buff[0] {
		case "find":
			phoneBook.FindNumber(num)
		case "del":
			phoneBook.DeleteNumber(num)
		case "add":
			phoneBook.AddNumberName(num, buff[2])
		default:
			fmt.Fprintln(Out, "how does it happen?")
		}
	}
}

// CreatePhoneBook создать записную книжку
func CreatePhoneBook() PhoneBook {
	pB := PhoneBook{}
	pB.records = make(map[string]string)
	return pB
}
