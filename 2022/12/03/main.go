package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic("file not found")
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	// closer, fileScanner := file.GetFileScanner("input.txt")
	// defer closer()

	// i := 1
	commonItems := []string{}
	for fileScanner.Scan() {

		// Part 1
		// 	items := fileScanner.Text()
		// 	halfway := len(items) / 2
		// 	compartment1 := items[0:halfway]
		// 	compartment2 := items[halfway:]
		// 	if len(compartment1) != len(compartment2) {
		// 		panic("items are not evenly divided")
		// 	}

		// 	// fmt.Println("sack", i)
		// 	// fmt.Println(compartment1)
		// 	// fmt.Println(compartment2)

		var commonItem string
		// 	outerBreak := false
		// 	for _, a := range compartment1 {
		// 		for _, b := range compartment2 {
		// 			if a == b {
		// 				commonItem = string(a)
		// 				outerBreak = true
		// 				break
		// 			}
		// 		}
		// 		if outerBreak {
		// 			break
		// 		}
		// 	}

		// 	if commonItem == "" {
		// 		panic("no common item found")
		// 	}
		// 	commonItems = append(commonItems, commonItem)
		// 	i = i + 1
		// }

		// fmt.Println(commonItems)
		// totalPriority := 0
		// priorities := populatePriorities()
		// for _, item := range commonItems {
		// 	totalPriority = totalPriority + priorities[string(item)]
		// 	fmt.Println(totalPriority, priorities[string(item)])
		// }
		// fmt.Println(totalPriority)

		group1 := fileScanner.Text()
		fileScanner.Scan()
		group2 := fileScanner.Text()
		fileScanner.Scan()
		group3 := fileScanner.Text()

		outerBreak := false
		for _, a := range group1 {
			for _, b := range group2 {
				if a == b {
					for _, c := range group3 {
						if b == c {
							commonItem = string(a)
							outerBreak = true
							break
						}
					}
				}
				if outerBreak {
					break
				}
			}
			if outerBreak {
				break
			}
		}

		commonItems = append(commonItems, commonItem)
	}

	fmt.Println(commonItems)
	fmt.Println(len(commonItems))
	totalPriority := 0
	priorities := populatePriorities()
	for _, item := range commonItems {
		totalPriority = totalPriority + priorities[string(item)]
		fmt.Println(totalPriority, priorities[string(item)])
	}
	fmt.Println(totalPriority)

}

func populatePriorities() map[string]int {
	key := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	priorities := make(map[string]int)
	for x, item := range key {
		priorities[string(item)] = x + 1
	}

	fmt.Println(priorities)
	return priorities
}

var priorities = map[string]int{
	"a": 1,
}

type Priority int32

const (
	a Priority = iota + 1
	b
	c
	d
	e
	f
	g
	h
	k
	l
	m
	n
	o
	p
	q
	r
	s
	t
	u
	v
	w
	x
	y
	z
	A
	B
	C
	D
	E
	F
	G
	H
	I
	J
	K
	L
	M
	O
	P
	Q
	R
	S
	T
	U
	V
	W
	X
	Y
	Z
)
