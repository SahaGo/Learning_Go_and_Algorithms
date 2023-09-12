package main

import "fmt"

//"This is an example!" ==> "sihT si na !elpmaxe"
//"double  spaces"      ==> "elbuod  secaps"

func main() {
	fmt.Println(ReverseWords("The quick brown fox jumps over the lazy dog."))
}

func ReverseWords(str string) string {

	var kStop int = 0
	var res string

	for k, v := range str {
		if v == 32 { // руна пробела
			for i := k - 1; i >= kStop; i-- {

				res = res + string(str[i])

			}
			res = res + " "
			kStop = k + 1
		}

		if k+1 == len(str) {
			for i := k; i >= kStop; i-- {

				res = res + string(str[i])

			}

		}
	}
	return res
}
