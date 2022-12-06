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

	// i := 1
	commonItems := []rune{}
	for fileScanner.Scan() {

		// Part 1 logic
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

		var commonItem rune
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

		// Part 2 logic
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
							commonItem = a
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

	totalPriority := 0
	for _, item := range commonItems {
		totalPriority = totalPriority + priority(item)
	}
	fmt.Println(totalPriority)

}

func priority(r rune) int {
	if r > 96 {
		return int(r) - 96
	}
	return int(r) - 38
}
