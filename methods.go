package golife

import (
	"errors"
	"fmt"
	"math"
)

// SiblingsAlive returns the count of siblings to a cell
// that are currently marked alive.
func SiblingsAlive(p Point, c Cell, ct CellTester) int {
	var i int

	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if x == 0 && y == 0 {
				continue
			}
			p2 := Point{X: p.X + int32(x), Y: p.Y + int32(y)}
			if a, _ := ct.IsAlive(p2); a == true {
				i++
			}
		}
	}

	return i
}

// GetCellState sets the life state of a Cell
func GetCellState(p Point, c Cell, ct CellTester) Cell {
	alive := SiblingsAlive(p, c, ct)
	if alive > 3 {
		c.Alive = false
	} else if alive == 3 {
		c.Alive = true
	} else if alive < 2 {
		c.Alive = false
	}
	return c
}

// NewGrid returns a new grid object
func NewGrid(count int32) (*Grid, int32, error) {
	if count == 0 {
		return nil, 0, errors.New("a grid with 0 cells does not make sense")
	}
	bound := int32(math.Floor(math.Sqrt(float64(count))))

	return &Grid{set: generateEmptyMap(bound, bound)}, int32(math.Pow(float64(bound), 2.0)), nil
}

// IsAlive returns whether a cell at a certain point is alive
// or dead. The function will return false if the cell does
// not exist as well as the existance
func (g *Grid) IsAlive(p Point) (bool, bool) {
	g.lock.RLock()
	defer g.lock.RUnlock()
	if c, ok := g.set[p]; ok {
		return c.Alive, ok
	}
	return false, false
}

// SetCell replaces a cell at a certain position with an updated
// one
func (g *Grid) SetCell(p Point, c Cell) error {
	g.lock.Lock()
	defer g.lock.Unlock()

	g.set[p] = c

	return nil
}

// Cell returns a cell at some point
func (g *Grid) Cell(p Point) (Cell, error) {
	g.lock.RLock()
	defer g.lock.RUnlock()

	c, ok := g.set[p]
	if !ok {
		return Cell{}, fmt.Errorf("at position %v no cell was found", p)
	}
	return c, nil
}

// GetCells gets a reference to a cell iterator
func (g *Grid) GetCells() map[Point]Cell {
	g.lock.RLock()
	defer g.lock.RUnlock()

	r := make(map[Point]Cell, len(g.set))
	it := 0
	for p, c := range g.set {
		r[p] = c
		it++
	}
	return r
}

// Count returns the internal number of cells
func (g *Grid) Count() int32 {
	return int32(len(g.set))
}

func generateEmptyMap(x, y int32) map[Point]Cell {
	var i, j int32
	n := x * y
	m := make(map[Point]Cell, n)

	for i = 0; i < x; i++ {
		for j = 0; j < y; j++ {
			m[Point{i, j}] = Cell{}
		}
	}

	return m
}
