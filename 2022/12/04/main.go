package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic("file not found")
	}
	defer inputFile.Close()

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	poorlyPairedAssignmentsPartOne := 0
	poorlyPairedAssignmentsPartTwo := 0
	for fileScanner.Scan() {
		assignments := strings.Split(fileScanner.Text(), ",")
		if len(assignments) < 2 {
			panic("not enough assignments")
		}
		assignment1 := getMinAndMax(assignments[0])
		assignment2 := getMinAndMax(assignments[1])
		if overlapCheckPartOne(assignment1, assignment2) {
			poorlyPairedAssignmentsPartOne = poorlyPairedAssignmentsPartOne + 1
		}

		if overlapCheckPartTwo(assignment1, assignment2) {
			poorlyPairedAssignmentsPartTwo = poorlyPairedAssignmentsPartTwo + 1
		}
	}
	fmt.Println("part one", poorlyPairedAssignmentsPartOne)
	fmt.Println("part two", poorlyPairedAssignmentsPartTwo)
}

type MinMax struct {
	Min int
	Max int
}

func overlapCheckPartOne(p1, p2 MinMax) bool {
	if p1.Min <= p2.Min && p1.Max >= p2.Max {
		return true
	}

	if p2.Min <= p1.Min && p2.Max >= p1.Max {
		return true
	}

	return false
}

func overlapCheckPartTwo(p1, p2 MinMax) bool {
	if p1.Max < p2.Min || p2.Max < p1.Min {
		return false
	}

	return true
}

func getMinAndMax(assignment string) MinMax {
	rangeVals := strings.Split(assignment, "-")
	if len(rangeVals) < 2 {
		panic("bad range")
	}

	min, _ := strconv.Atoi(rangeVals[0])
	max, _ := strconv.Atoi(rangeVals[1])

	return MinMax{min, max}
}
