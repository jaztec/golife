package golife

import (
	"sync"
)

// Simulator described function needed for a game
type Simulator interface {
	// SetGrid replaces the grid with which this simulation is working
	SetGrid(*Grid) error
	// Step calculates the next state of the internal grid
	Step() error
	// Gen reports the current step since the last grid change
	Gen() int
}

// CellTester exposes a read-only function to check the state of
// a cell at a certain position.
type CellTester interface {
	// IsAlive returns whether a cell at a certain point is alive
	// or dead. The function should return false if the cell does
	// not exist as well as the existance
	IsAlive(Point) (alive bool, exists bool)
}

// CellSetter defines a function to update a cell inside the
// struct
type CellSetter interface {
	// SetCell replaces a cell at a certain position with an updated
	// one
	SetCell(Point, Cell) error
}

// Point holds the position of a Cell as well as its key
type Point struct {
	// X and Y represent the position of this point inside a 2D grid
	X, Y int32
}

// Cell represents a cell on the Game of Life board
type Cell struct {
	// A Cell keeps track whether its alive
	Alive bool
}

// Grid represents the Game of Life board
type Grid struct {
	set  map[Point]Cell
	lock sync.RWMutex
}
