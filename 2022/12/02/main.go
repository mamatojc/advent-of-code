package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	totalScore := 0

	for fileScanner.Scan() {
		movePair := fileScanner.Text()
		moves := strings.Split(movePair, " ")
		moveA := NewMove(moves[0])
		moveB := NewCounterMove(moveA, moves[1])
		// fmt.Println("Player A", moves[0])
		// fmt.Println("Player B", moves[1])

		result := CheckResult(moveA, moveB)
		// fmt.Println("result", result)
		totalScore = totalScore + result.score
	}
	fmt.Println("total score", totalScore)
}

type Move struct {
	name  string
	moves []string
	score int
}

func NewMove(s string) Move {
	if s == "A" || s == "X" {
		return rock
	}

	if s == "B" || s == "Y" {
		return paper
	}

	if s == "C" || s == "Z" {
		return scissors
	}

	panic("unknown move")
}

func NewCounterMove(move Move, counter string) Move {
	if move.name == "rock" {
		if counter == "X" {
			return scissors
		}

		if counter == "Y" {
			return rock
		}

		if counter == "Z" {
			return paper
		}
	}

	if move.name == "paper" {
		if counter == "X" {
			return rock
		}

		if counter == "Y" {
			return paper
		}

		if counter == "Z" {
			return scissors
		}
	}

	if move.name == "scissors" {
		if counter == "X" {
			return paper
		}

		if counter == "Y" {
			return scissors
		}

		if counter == "Z" {
			return rock
		}
	}

	panic("unknown move")
}

var rock = Move{
	name:  "rock",
	moves: []string{"A", "X"},
	score: 1,
}

var paper = Move{
	name:  "paper",
	moves: []string{"B", "Y"},
	score: 2,
}

var scissors = Move{
	name:  "scissors",
	moves: []string{"C", "Z"},
	score: 3,
}

type Result struct {
	result string
	score  int
}

// Gives back the result for player B
func CheckResult(a, b Move) Result {

	if a.name == "rock" {
		// draw
		if b.name == "rock" {
			return Result{"draw", b.score + 3}
		}

		// lose
		if b.name == "scissors" {
			return Result{"lose", b.score + 0}
		}

		// win
		if b.name == "paper" {
			return Result{"win", b.score + 6}
		}
	}

	if a.name == "paper" {
		// lose
		if b.name == "rock" {
			return Result{"lose", b.score + 0}
		}

		// win
		if b.name == "scissors" {
			return Result{"win", b.score + 6}
		}

		// draw
		if b.name == "paper" {
			return Result{"draw", b.score + 3}
		}
	}

	if a.name == "scissors" {
		// win
		if b.name == "rock" {
			return Result{"win", b.score + 6}
		}

		// win
		if b.name == "scissors" {
			return Result{"draw", b.score + 3}
		}

		// draw
		if b.name == "paper" {
			return Result{"lose", b.score + 0}
		}
	}

	panic("cannot determine result")
}
