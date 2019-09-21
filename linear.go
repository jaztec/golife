package golife

import (
	"fmt"
	"math"
)

type linearSimulator struct {
	grid *Grid
	gen  int
}

func (bs *linearSimulator) Step() error {
	n := int32(math.Sqrt(float64(bs.grid.Count())))
	g, n, err := NewGridFromCells(int32(n), int32(n), bs.grid.GetCells())
	if n != int32(bs.Grid().Count()) {
		return fmt.Errorf("old size %d does not match new size %d", n, bs.Grid().Count())
	}
	if err != nil {
		return err
	}
	for p, c := range bs.grid.set {
		g.SetCell(p, SetCellState(p, c, bs.grid))
	}

	bs.grid = g
	bs.gen++
	return nil
}

func (bs *linearSimulator) Gen() int {
	return bs.gen
}

func (bs *linearSimulator) Grid() *Grid {
	return bs.grid
}

func (bs *linearSimulator) SetGrid(g *Grid) error {
	bs.grid = g
	return nil
}

// NewLinearSimulator returns a new simulator that processes a
// by game linear processing
func NewLinearSimulator(cellCount int32) (Simulator, error) {
	g, _, err := NewGrid(cellCount)
	if err != nil {
		return nil, err
	}
	return &linearSimulator{
		grid: g,
		gen:  0,
	}, nil
}
