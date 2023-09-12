package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	txt, err := os.Open("/Users/kisats/projects/stepik_3.5.13/taskdata.txt")
	if err != nil {
		fmt.Println("Ошибос")
		return
	}
	defer txt.Close()

	reader := bufio.NewReader(txt)
	var count int = 0
	for {
		num_0, err := reader.ReadString(59)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}

		count++
		if len(num_0) == 2 {
			fmt.Println(num_0, " Позиция:", count)
		}
	}

}
