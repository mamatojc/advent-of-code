package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	framework "github.com/ppg/advent-of-code/2022/12/framework"
)

func main() {
	framework.Register(parser, solution0)
	framework.Run(os.Stdout)
}

var parser = func(line string) string { return line }

func solution0(w io.Writer, runner *framework.Runner[string]) {
	lines := runner.Lines()
	firstLine := <-lines
	if firstLine != "$ cd /" {
		panic("not starting at top of tree")
	}

	tree := Node{
		Name:   "/",
		Type:   Directory,
		Parent: nil,
	}

	currentNode := &tree
	for line := range lines {
		subParts := strings.Split(line, " ")

		// Command
		if subParts[0] == "$" {
			if subParts[1] == "cd" {
				if len(subParts) < 3 {
					fmt.Println(subParts)
					panic("bad cd command")
				}
				currentNode = processCdCommand(subParts[2], currentNode)
				if currentNode == nil {
					fmt.Println(subParts)
					panic("could not find dir")
				}
				continue
			} else {
				// implies an `ls` command, skip over it to process following lines
				continue
			}
		}

		if subParts[0] == "dir" {
			if len(subParts) < 2 {
				fmt.Println(subParts)
				panic("bad dir")
			}
			dir := Node{
				Name:   subParts[1],
				Type:   Directory,
				Parent: currentNode,
			}
			currentNode.addChild(dir)
		} else {
			// File
			if len(subParts) < 2 {
				fmt.Println(subParts)
				panic("bad file")
			}

			fileSize, err := strconv.Atoi(subParts[0])
			if err != nil {
				fmt.Println(subParts)
				panic("bad file size")
			}

			file := Node{
				Name:   subParts[1],
				Type:   File,
				Size:   fileSize,
				Parent: currentNode,
			}
			currentNode.addChild(file)
		}
	}

	fmt.Println(tree.countChildren())
	tree.calculateDirectorySizes()
	totalDiskSize := 70000000
	freeDiskSpace := totalDiskSize - tree.Size
	spaceNeeded := 30000000
	needToFree := spaceNeeded - freeDiskSpace
	fmt.Println("Total disk size is", totalDiskSize)
	fmt.Println("Total on disk is", tree.Size)
	fmt.Println("Free disk space", freeDiskSpace)
	fmt.Printf("Need %d. Need to free %d\n", spaceNeeded, needToFree)

	totalSizeDirsUnder100k := 0
	candidatesToDelete := []*Node{}
	queue := []*Node{}
	queue = append(queue, &tree)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Found the node
		if current.Type == Directory {
			// fmt.Printf("directory - name: %s, size: %d\n", current.Name, current.Size)
			if current.Size <= 100000 {
				totalSizeDirsUnder100k = totalSizeDirsUnder100k + current.Size
			}
			if current.Size > needToFree {
				// fmt.Printf("candidate - name: %s, size: %d\n", current.Name, current.Size)
				candidatesToDelete = append(candidatesToDelete, current)
			}
		}
		if len(current.Children) > 0 {
			queue = append(queue, current.Children...)
		}
	}

	dirToDelete := Node{
		Size: math.MaxInt,
	}
	for _, candidate := range candidatesToDelete {
		if candidate.Size < dirToDelete.Size {
			dirToDelete = *candidate
		}
	}

	fmt.Println("Part 1 answer", totalSizeDirsUnder100k)
	fmt.Println("Part 2 answer", dirToDelete.Size)
}

// Process a change directory command
func processCdCommand(name string, node *Node) *Node {
	// move up one directory
	if name == ".." {
		return node.Parent
	}
	return BFS(node, name)
}

type Type uint8

const (
	Directory Type = iota
	File
)

type Node struct {
	Name     string
	Type     Type
	Size     int
	Children []*Node
	Parent   *Node
}

func (n *Node) addChild(child Node) {
	n.Children = append(n.Children, &child)
}

func (n *Node) calculateDirectorySizes() {
	if len(n.Children) == 0 {
		return
	}

	if n.Type == File {
		return
	}

	// Children present, not a file
	directorySize := 0
	for _, child := range n.Children {
		child.calculateDirectorySizes()
		directorySize = directorySize + child.Size

	}
	n.Size = directorySize
}

func (n *Node) countChildren() int {
	if len(n.Children) == 0 {
		return 0
	}

	count := len(n.Children)
	for _, child := range n.Children {
		if len(child.Children) > 0 {
			count = count + len(child.Children)
		}
		count = count + child.countChildren()
	}

	return count
}

func BFS(root *Node, name string) *Node {
	queue := []*Node{}
	queue = append(queue, root)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Found the node
		if current.Name == name {
			return current
		}
		if len(current.Children) > 0 {
			queue = append(queue, current.Children...)
		}
	}
	return nil
}
