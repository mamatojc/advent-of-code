package file

import (
	"bufio"
	"os"
)

func GetFileScanner(path string) (fileCloser func(), fileScanner *bufio.Scanner) {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		panic("file not found")
	}
	fileCloser = func() { inputFile.Close() }

	fileScanner = bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	return
}
