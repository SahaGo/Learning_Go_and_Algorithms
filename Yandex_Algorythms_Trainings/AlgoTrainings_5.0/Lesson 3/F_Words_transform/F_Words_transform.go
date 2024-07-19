package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	shortSet := make(map[string]bool)

	line, err := in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	line, _ = strings.CutSuffix(line, "\n")
	arrSet := strings.Split(line, " ")
	for _, s := range arrSet {
		shortSet[s] = true
	}

	line, err = in.ReadString('\n')

	arrSet = strings.Split(line, " ")

	ans := make([]string, 0, len(arrSet))
	for _, s := range arrSet { // words
		for i := 0; i < len(s); i++ { //letters
			part := s[:i]
			if _, ok := shortSet[part]; ok {
				ans = append(ans, part)
				break
			}
			if i == len(s)-1 {
				ans = append(ans, s)
			}
		}
	}
	fmt.Fprintf(out, "%s", strings.Join(ans, " "))
}
