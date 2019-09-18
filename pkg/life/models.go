package life

// Simulator described function needed for a game
type Simulator interface {
	Step() error
	Gen() int
}

// CellTester exposes a read-only function to check the state of
// a cell at a certain position.
type CellTester interface {
	IsAlive(Point) (alive bool, exists bool)
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
	set map[Point]Cell
}
