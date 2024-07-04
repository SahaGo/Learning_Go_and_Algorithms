package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	listenersCount := 0
	fmt.Fscan(in, &listenersCount)

	if listenersCount == 1 {
		processOne(in, out)
		return
	}
	likes := make(map[string]int) // хранилище треков

	res := make([]string, 0)
	for i := 0; i < listenersCount; i++ {
		trackCount := 0 // количество треков
		fmt.Fscan(in, &trackCount)

		for j := 0; j < trackCount; j++ { // считываем треки в хранилище
			trackName := ""
			fmt.Fscan(in, &trackName)

			if _, ok := likes[trackName]; !ok { // трека нет - добавляем скольким людям еще должен понравится
				likes[trackName] = listenersCount
			} else {
				likes[trackName]--
				if likes[trackName] == 1 { // если 0, то знвчит трек нравится всем и добавляем в ответ
					res = append(res, trackName)
				}
			}
		}

	}
	slices.Sort(res)
	s := strings.Trim(fmt.Sprint(res), "[]")
	//fmt.Fprintln(out, likes)
	fmt.Fprintf(out, "%d\n%s", len(res), s)
}

func processOne(in *bufio.Reader, out *bufio.Writer) {
	tracks := 0
	fmt.Fscan(in, &tracks)
	res := make([]string, tracks)

	for i := 0; i < tracks; i++ {
		fmt.Fscan(in, &res[i])
	}

	slices.Sort(res)

	fmt.Fprintf(out, "%d\n%s", len(res), strings.Join(res, " "))
}
