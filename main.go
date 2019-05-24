package main

import (
	"bufio"
	"flag"
  "os"
  "log"
)

const (
	PLACE  = "PLACE"
	MOVE   = "MOVE"
	LEFT   = "LEFT"
	RIGHT  = "RIGHT"
	REPORT = "REPORT"
)

var scanner *bufio.Scanner

func main() {
	var fileOpt bool
	file, err := os.Open("./data.txt")

	if err != nil {
    log.Fatalf("Error, cannot open file %v", err)
  }

	flag.BoolVar(&fileOpt, "file", true, "Read data from a file")
	flag.Parse()


	board := &Board{}

	// Check if fileOpt is passed to decide whether to get data from a file or standard input
	if fileOpt {
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	scanner.Split(bufio.ScanWords)

	// Scan the file line by line and run method corresponding to the command
	for scanner.Scan() {
		token := scanner.Text()
		board.executeCommand(token)
	}
}
