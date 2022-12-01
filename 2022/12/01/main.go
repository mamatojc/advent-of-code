package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("reading file")
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic("file not found")
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	counter := 1
	elfCalories := 0
	maxCalories := 0
	maxElf := 0
	calories := []int{}

	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			// fmt.Println(fmt.Sprintf("The %v elf is carrying %v calories", counter, elfCalories))
			calories = append(calories, elfCalories)

			if elfCalories > maxCalories {
				maxCalories = elfCalories
				maxElf = counter
			}

			counter = counter + 1
			elfCalories = 0
		}
		calories, _ := strconv.Atoi(fileScanner.Text())
		elfCalories = elfCalories + calories
	}

	fmt.Println(
		fmt.Sprintf(
			"The %v elf is carrying %v calories. This is the most of all of the elves.",
			maxElf, maxCalories),
	)

	sumCalories := 0
	sort.Ints(calories)
	for i := len(calories) - 3; i < len(calories); i++ {
		fmt.Println(calories[i])
		sumCalories = sumCalories + calories[i]
	}

	fmt.Println(fmt.Sprintf("The top 3 elves are carrying %v calories", sumCalories))
}
