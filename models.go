package golife

import (
	"sync"
)

// Simulator described function needed for a game
type Simulator interface {
	SetGrid(*Grid) error
	Step() error
	Gen() int
}

// CellTester exposes a read-only function to check the state of
// a cell at a certain position.
type CellTester interface {
	IsAlive(Point) (alive bool, exists bool)
}

// CellSetter defines a function to update a cell inside the
// struct
type CellSetter interface {
	SetCell(Point, Cell) error
}

// Point holds the position of a Cell as well as its key
type Point struct {
	X, Y int32
}

// Cell represents a cell on the Game of Life board
type Cell struct {
	Alive bool
}

// Grid represents the Game of Life board
type Grid struct {
	set  map[Point]Cell
	lock sync.RWMutex
}
