package main

import (
	"fmt"
	"io"
	"os"

	framework "github.com/ppg/advent-of-code/2022/12/framework"
)

func main() {
	framework.Register(parser, solution0)
	framework.Run(os.Stdout)
}

var parser = func(line string) string { return line }

func solution0(w io.Writer, runner *framework.Runner[string]) {
	lines := runner.Lines()
	datastream := <-lines

	lookAhead := 14

	var length int
	var next4 [14]byte
	for i := 0; i < len(datastream)-lookAhead; i++ {
		for j := 0; j < lookAhead; j++ {
			next4[j] = datastream[i+j]
		}
		// fmt.Println(i, next4)

		found := false
		outerBreak := false
		for k := 0; k < lookAhead; k++ {
			for l := 0; l < lookAhead; l++ {
				if k != l && next4[k] == next4[l] {
					// fmt.Println("break")
					outerBreak = true
					break
				} else {
					if k == lookAhead-1 && l == lookAhead-1 {
						found = true
					}
				}
			}
			if outerBreak {
				break
			}
		}

		if found {
			// fmt.Println("found it!")
			length = i
			break
		}
	}

	fmt.Println("The messages starts at", length+4)
}
