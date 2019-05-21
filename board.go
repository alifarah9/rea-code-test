package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	NORTH = "NORTH"
	SOUTH = "SOUTH"
	EAST  = "EAST"
	WEST  = "WEST"
)

type Board struct {
	x      int
	y      int
	Facing string
}

func (board *Board) executeCommand(command string) {
	switch command {
	case PLACE:
		board.Place()
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

func (b *Board) Move() {
	if b.Facing == "" {
		return
	}
	// Check the current facing position of the board and attempt to move
	// in that direction if new position is valid.
	switch b.Facing {
	case NORTH:
		if b.Validate(b.x, b.y+1) {
			b.y = b.y + 1
		}
	case SOUTH:
		if b.Validate(b.x, b.y-1) {
			b.y = b.y - 1
		}
	case EAST:
		if b.Validate(b.x-1, b.y) {
			b.x = b.x - 1
		}
	case WEST:
		if b.Validate(b.x+1, b.y) {
			b.x = b.x + 1
		}
	}
}

func (b *Board) Validate(x int, y int) bool {
	if b.Facing == "" {
		return false
	}
	// Check if given positions are valid
	if x > -1 && x < 5 && y > -1 && y < 5 {
		return true
	}
	return false
}

func (b *Board) MoveLeft() {
	if b.Facing == "" {
		return
	}
	// Check the current facing position of the board and update
	// the facing position accordingly. i.e 90 degrees counter clockwise
	switch b.Facing {
	case NORTH:
		b.Facing = WEST
	case WEST:
		b.Facing = SOUTH
	case SOUTH:
		b.Facing = EAST
	case EAST:
		b.Facing = NORTH
	}
}

func (b *Board) MoveRight() {
	if b.Facing == "" {
		return
	}

	// Check the current facing position of the board and update
	// the facing position accordingly. i.e 90 degrees clockwise
	switch b.Facing {
	case NORTH:
		b.Facing = EAST
	case WEST:
		b.Facing = NORTH
	case SOUTH:
		b.Facing = WEST
	case EAST:
		b.Facing = SOUTH
	}
}

func (b *Board) Place() {
	// Excecute scanner to get next token for the PLACE command i.e (X,Y,F)
	scanner.Scan()
	position := scanner.Text()
	match, _ := regexp.MatchString("^([0-5]),([0-5]),(NORTH|SOUTH|EAST|WEST)", position)
	if !match {
		scanner.Scan()
		return
	}
	s := strings.Split(position, ",")
	// Convert the string into integer, check for errors,
	// and then initialise board position
	x, err := strconv.Atoi(s[0])
	if err != nil {
		panic(err)
	}

	y, err := strconv.Atoi(s[1])

	if err != nil {
		panic(err)
	}

	Facing := s[2]
	b.x = x
	b.y = y
	b.Facing = Facing
}
