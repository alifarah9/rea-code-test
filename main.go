package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

const (
	PLACE  = "PLACE"
	MOVE   = "MOVE"
	LEFT   = "LEFT"
	RIGHT  = "RIGHT"
	REPORT = "REPORT"
)

func main() {
	var fileOpt bool
	file, err := os.Open("./data.txt")

	flag.BoolVar(&fileOpt, "file", true, "Read data from a file")
	flag.Parse()
	if err != nil {
		panic(err)
	}
	var scanner *bufio.Scanner

	board := &Board{}

	// Check if fileOpt is passed to decide whether to get data
	// from a file or standard input
	if fileOpt {
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	// Set the Split method to ScanWords.
	scanner.Split(bufio.ScanWords)

	// Scan the file line by line and run method corresponding to
	// the command
	// Note: If PLACE is found, scanner is executed again the to get the
	// position data
	for scanner.Scan() {
		token := scanner.Text()
		board.executeCommand(token, scanner)
	}
}

func (board *Board) executeCommand(command string, scanner *bufio.Scanner) {
	switch command {
	case PLACE:
		board.Place(scanner)
	case MOVE:
		board.Move()
	case LEFT:
		board.MoveLeft()
	case RIGHT:
		board.MoveRight()
	case REPORT:
		fmt.Printf("%d,%d,%s\n", board.x, board.y, board.Facing)
	}
}
