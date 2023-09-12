package main

import (
	"bufio"
	"fmt"
	"os"
)

//«freeze» — охлаждение. В этом режиме кондиционер может только уменьшать температуру. Если температура в комнате и так не больше желаемой, то он выключается.
//«heat» — нагрев. В этом режиме кондиционер может только увеличивать температуру. Если температура в комнате и так не меньше желаемой, то он выключается.
//«auto» — автоматический режим. В этом режиме кондиционер может как увеличивать, так и уменьшать температуру в комнате до желаемой.
//«fan» — вентиляция. В этом режиме кондиционер осуществляет только вентиляцию воздуха и не изменяет температуру в комнате.

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tRoom int
	var tCond int

	fmt.Fscanln(in, &tRoom, &tCond)

	var programm string

	fmt.Fscanln(in, &programm)

	switch programm {
	case "freeze":
		switch {
		case tRoom < tCond:
			fmt.Fprint(out, tRoom)
		}
	case "heat":

	case "auto":

	case "fan":

	}

}
