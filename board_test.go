package main

import (
	"testing"
)

func TestMoveLeft(t *testing.T) {
	tests := []struct {
		name  string
		board Board
		want  string
	}{
		{
			// Changes the facing direction to WEST when MOVE LEFT
			name: "facing north",
			board: Board{
				Facing: "NORTH",
			},
			want: "WEST",
		},
		{
			name: "facing south",
			board: Board{
				Facing: "SOUTH",
			},
			want: "EAST",
		},
		{
			name: "facing east",
			board: Board{
				Facing: "NORTH",
			},
			want: "WEST",
		},
		{
			name: "facing west",
			board: Board{
				Facing: "WEST",
			},
			want: "SOUTH",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				x:      tt.board.x,
				y:      tt.board.y,
				Facing: tt.board.Facing,
			}
			b.MoveLeft()

			if b.Facing != tt.want {
				t.Fatal()
			}
		})
	}
}

func TestMoveRight(t *testing.T) {
	type board struct {
		x      int
		y      int
		Facing string
	}
	tests := []struct {
		name  string
		board board
		want  string
	}{
		{
			// Changes the facing direction correctly when MOVE RIGHT
			name: "facing north",
			board: board{
				Facing: "NORTH",
			},
			want: "EAST",
		},
		{
			name: "facing south",
			board: board{
				Facing: "SOUTH",
			},
			want: "WEST",
		},
		{
			name: "facing east",
			board: board{
				Facing: "EAST",
			},
			want: "SOUTH",
		},
		{
			name: "facing west",
			board: board{
				Facing: "WEST",
			},
			want: "NORTH",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Board{
				x:      tt.board.x,
				y:      tt.board.y,
				Facing: tt.board.Facing,
			}
			b.MoveRight()

			if b.Facing != tt.want {
				t.Fatal()
			}
		})
	}
}

// Test Move() method
func TestMove(t *testing.T) {
	type board struct {
		x      int
		y      int
		Facing string
	}
	tests := []struct {
		name  string
		board board
		want  int
	}{
		{
			// Changes the facing direction correctly when MOVE RIGHT
			name: "moves north",
			board: board{
				Facing: "NORTH",
				x:      1,
				y:      2,
			},
			want: 3,
		},
		{
			name: "moves south",
			board: board{
				Facing: "SOUTH",
				x:      3,
				y:      3,
			},
			want: 2,
		},
		{
			name: "moves east",
			board: board{
				Facing: "EAST",
				x:      1,
				y:      1,
			},
			want: 0,
		},
		{
			name: "moves west",
			board: board{
				Facing: "WEST",
				x:      1,
				y:      4,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := &Board{
				x:      tt.board.x,
				y:      tt.board.y,
				Facing: tt.board.Facing,
			}
			board.Move()

			if board.Facing == NORTH && board.y != tt.board.y+1 {
				t.Fatal()
			}
			if board.Facing == EAST && board.x != tt.board.x-1 {
				t.Fatal()
			}
			if board.Facing == SOUTH && board.y != tt.board.y-1 {
				t.Fatal()
			}
			if board.Facing == WEST && board.x != tt.board.x+1 {
				t.Fatal()
			}
		})
	}
}

// Test validate method
func TestValidate(t *testing.T) {
	type board struct {
		x      int
		y      int
		Facing string
	}
	tests := []struct {
		name  string
		board board
		want  bool
	}{

		{
			name: "facing value missing",
			board: board{
				Facing: "",
				x:      -1,
				y:      3,
			},
			want: false,
		},
		{
			name: "facing value is invalid string",
			board: board{
				Facing: "XEAST",
				x:      -1,
				y:      3,
			},
			want: false,
		},
		{
			name: "negative x value",
			board: board{
				Facing: "NORTH",
				x:      -1,
				y:      3,
			},
			want: false,
		},
		{
			name: "above 5 in x direction",
			board: board{
				Facing: "NORTH",
				x:      5,
				y:      1,
			},
			want: false,
		},
		{
			name: "above 5 in y direction",
			board: board{
				Facing: "NORTH",
				x:      2,
				y:      5,
			},
			want: false,
		},
		{
			name: "inside bounds",
			board: board{
				Facing: "EAST",
				x:      1,
				y:      2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := &Board{
				x:      tt.board.x,
				y:      tt.board.y,
				Facing: tt.board.Facing,
			}

			result := board.Validate(board.x, board.y)

			if result != tt.want {
				t.Fatal()
			}
		})
	}
}
