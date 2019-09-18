package life

import "math"

// CellTester exposes a read-only function to check the state of
// a cell at a certain position.
type CellTester interface {
	IsAlive(Point) bool
}

// SiblingsAlive returns the count of siblings to a cell
// that are currently marked alive.
func SiblingsAlive(c Cell, ct CellTester) int {
	var i int

	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if x == 0 && y == 0 {
				continue
			}
			p := Point{X: c.X + int64(x), Y: c.Y + int64(y)}
			if ct.IsAlive(p) {
				i++
			}
		}
	}

	return i
}

// SetCellState sets the life state of a Cell
func SetCellState(c Cell, ct CellTester) Cell {
	alive := SiblingsAlive(c, ct)
	if alive > 3 {
		c.Alive = false
	} else if alive == 3 {
		c.Alive = true
	} else if alive < 2 {
		c.Alive = false
	}
	return c
}

// Point holds the position of a Cell as well as its key
type Point struct {
	X, Y int64
}

// Cell represents a cell on the Game of Life board
type Cell struct {
	Point
	Alive bool
}

// Grid represents the Game of Life board
type Grid struct {
	set map[Point]Cell
}

// IsAlive returns whether a cell at a certain point is alive
// of dead. The function will return false if the cell does
// not exist
func (g *Grid) IsAlive(p Point) bool {
	if c, ok := g.set[p]; ok {
		return c.Alive
	}
	return false
}

// NewGrid returns a new grid object
func NewGrid(count int64) (*Grid, int64) {
	bound := int64(math.Floor(math.Sqrt(float64(count))))

	return &Grid{generateEmptyMap(bound, bound)}, int64(math.Pow(float64(bound), 2.0))
}

func generateEmptyMap(x, y int64) map[Point]Cell {
	var i, j int64
	n := x * y
	m := make(map[Point]Cell, n)

	for i = 0; i < x; i++ {
		for j = 0; j < y; j++ {
			p := Point{i, j}
			m[p] = Cell{p, false}
		}
	}

	return m
}
