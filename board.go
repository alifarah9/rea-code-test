package main

import (
	"bufio"
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

func (b *Board) Move() {
	if b.Facing == "" {
		return
	}
	// Check the current facing position of the board and attempt to move
	// in that direction if new position is valid.
	switch b.Facing {
	case NORTH:
		b.Validate(b.x, b.y+1)
	case SOUTH:
		b.Validate(b.x, b.y-1)
	case EAST:
		b.Validate(b.x+1, b.y)
	case WEST:
		b.Validate(b.x-1, b.y)
	}
}

func (b *Board) Validate(x int, y int) bool {
	if b.Facing == "" {
		return false
	}
	// Check if given positions are valid
	if x > -1 && x < 6 && y > -1 && y < 6 {
		b.x = x
		b.y = y
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

func (b *Board) Place(scanner *bufio.Scanner) {
	// Excecute scanner to get next token for the PLACE command i.e (X,Y,F)
	scanner.Scan()
	//
	s := strings.Split(scanner.Text(), ",")
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
