package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic("file not found")
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	data := map[int]*Stack{}
	for i := 1; i <= 9; i++ {
		data[i] = new(Stack)
	}

	/*
	   [T]     [Q]             [S]
	   [R]     [M]             [L] [V] [G]
	   [D] [V] [V]             [Q] [N] [C]
	   [H] [T] [S] [C]         [V] [D] [Z]
	   [Q] [J] [D] [M]     [Z] [C] [M] [F]
	   [N] [B] [H] [N] [B] [W] [N] [J] [M]
	   [P] [G] [R] [Z] [Z] [C] [Z] [G] [P]
	   [B] [W] [N] [P] [D] [V] [G] [L] [T]
	    1   2   3   4   5   6   7   8   9
	*/

	data[1].Push("B")
	data[1].Push("P")
	data[1].Push("N")
	data[1].Push("Q")
	data[1].Push("H")
	data[1].Push("D")
	data[1].Push("R")
	data[1].Push("T")

	data[2].Push("W")
	data[2].Push("G")
	data[2].Push("B")
	data[2].Push("J")
	data[2].Push("T")
	data[2].Push("V")

	data[3].Push("N")
	data[3].Push("R")
	data[3].Push("H")
	data[3].Push("D")
	data[3].Push("S")
	data[3].Push("V")
	data[3].Push("M")
	data[3].Push("Q")

	data[4].Push("P")
	data[4].Push("Z")
	data[4].Push("N")
	data[4].Push("M")
	data[4].Push("C")

	data[5].Push("D")
	data[5].Push("Z")
	data[5].Push("B")

	data[6].Push("V")
	data[6].Push("C")
	data[6].Push("W")
	data[6].Push("Z")

	data[7].Push("G")
	data[7].Push("Z")
	data[7].Push("N")
	data[7].Push("C")
	data[7].Push("V")
	data[7].Push("Q")
	data[7].Push("L")
	data[7].Push("S")

	data[8].Push("L")
	data[8].Push("G")
	data[8].Push("J")
	data[8].Push("M")
	data[8].Push("D")
	data[8].Push("N")
	data[8].Push("V")

	data[9].Push("T")
	data[9].Push("P")
	data[9].Push("M")
	data[9].Push("F")
	data[9].Push("Z")
	data[9].Push("C")
	data[9].Push("G")

	fmt.Println("processing moves")
	for fileScanner.Scan() {
		parsed := strings.Split(fileScanner.Text(), " ")
		if len(parsed) < 5 {
			panic("bad move")
		}

		// parsed[0] = "move"
		quantity, err := strconv.Atoi(parsed[1])
		if err != nil {
			panic("bad quantity")
		}
		// parsed[2] = "from"
		from, err := strconv.Atoi(parsed[3])
		if err != nil {
			panic("bad from")
		}
		// parsed[4] = "to"
		to, err := strconv.Atoi(parsed[5])
		if err != nil {
			panic("bad to")
		}

		// fmt.Printf("moving %d crates from %d to %d\n", quantity, from, to)
		// fmt.Printf("top of %d - %v\n", from, data[from].top.value)
		// Part One logic
		// var crate string
		// for i := 0; i < quantity; i++ {
		// 	var valid bool
		// 	crate, valid = data[from].Pop()
		// 	if !valid {
		// 		fmt.Println("stack empty", from)
		// 	}
		// 	data[to].Push(crate)
		// }

		// Part Two logic
		var crates []string
		for i := 0; i < quantity; i++ {
			crate, valid := data[from].Pop()
			crates = append(crates, crate)
			if !valid {
				fmt.Println("stack empty", from)
			}
		}
		fmt.Println(crates)

		for i := len(crates) - 1; i >= 0; i-- {
			data[to].Push(crates[i])
		}
		fmt.Printf("top of %d - %v\n", to, data[to].top.value)
	}

	for i := 1; i <= 9; i++ {
		fmt.Print(data[i].top.value)
	}
}

type item struct {
	value string
	next  *item
}

type Stack struct {
	top  *item
	size int64
	lock sync.Mutex
}

func (stack *Stack) Len() int64 {
	return stack.size
}

func (stack *Stack) Push(value string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.top = &item{
		value: value,
		next:  stack.top,
	}
	stack.size++
}

func (stack *Stack) Pop() (value string, valid bool) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if stack.Len() > 0 {
		value = stack.top.value
		stack.top = stack.top.next
		stack.size--
		return value, true
	}

	return "", false
}
