package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type PhoneBook struct {
	records map[string]string
}

func (pb *PhoneBook) AddNumberName(num string, name string) {
	pb.records[num] = name
}

func (pb *PhoneBook) DeleteNumber(num string) {
	delete(pb.records, num)
}

func (pb *PhoneBook) FindNumber(num string) {
	if v, ok := pb.records[num]; ok {
		fmt.Fprintln(Out, v)
		return
	}
	fmt.Fprintln(Out, "not found")
	return
}

type Methods interface {
	AddNumberName(string, string)
	DeleteNumber(string)
	FindNumber(string)
}

var (
	In  = bufio.NewReader(os.Stdin)
	Out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer Out.Flush()
	var request int
	fmt.Fscanln(In, &request)

	phoneBook := CreatePhoneBook()

	for i := 0; i < request; i++ {

		line, err := In.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		line = strings.Trim(line, "\n")
		buff := strings.Split(line, " ")

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

func CreatePhoneBook() PhoneBook {
	pB := PhoneBook{}
	pB.records = make(map[string]string)
	return pB
}
