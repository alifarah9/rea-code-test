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
	f, err := os.Open("./data.txt")

	flag.BoolVar(&fileOpt, "file", true, "Read data from a file")

	if err != nil {
		panic(err)
	}

	board := &Board{}
	scanner := bufio.NewScanner(f)
	// Set the Split method to ScanWords.
	scanner.Split(bufio.ScanWords)

	// Scan the file line by line and run method corresponding to
	// the command
	// Note: If PLACE is found, scanner is executed again the to get the
	// position data
	for scanner.Scan() {
		token := scanner.Text()
		executeCommand(token)
	}
}

func executeCommand(command string) {
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
		fmt.Printf("%d,%d,%s", board.x, board.y, board.Facing)
	}
}
