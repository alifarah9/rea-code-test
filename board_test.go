package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestToyRobot(t *testing.T) {

	var testCommands = []string{
		"PLACE 1,2,EAST",
		"MOVE",
		"MOVE",
		"LEFT",
		"MOVE",
		"REPORT",
	}

	board := &Board{}
	_ = board
	for n, test := range testCommands {
		buf := strings.NewReader(test)
		s := bufio.NewScanner(buf)
		s.Split(bufio.ScanWords)
		token := s.Text()
		fmt.Println("d")

		_ = token
		_ = n
		_ = s

	}

}
