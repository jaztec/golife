package golife_test

import (
	"fmt"
	"testing"

	"github.com/jaztec/golife"
)

func TestNewGrid(t *testing.T) {
	var expect int32 = 196
	t.Run("test exact number of cells", func(t *testing.T) {
		_, got, _ := golife.NewGrid(expect)

		if got != expect {
			t.Errorf("expected %d cells but %d where created", expect, got)
		}
	})
	t.Run("test calculated number of cells", func(t *testing.T) {
		_, got, _ := golife.NewGrid(200)

		if got != expect {
			t.Errorf("expected %d cells but %d where created", expect, got)
		}
	})
	t.Run("make sure the grid implements the proper interfaces", func(t *testing.T) {
		var g interface{}
		g, _, _ = golife.NewGrid(4)

		if _, ok := g.(golife.CellTester); !ok {
			t.Errorf("expected Grid to implement 'CellTester', got %T", g)
		}
		if _, ok := g.(golife.CellSetter); !ok {
			t.Errorf("expected Grid to implement 'CellSetter', got %T", g)
		}
		if _, ok := g.(golife.CellGetter); !ok {
			t.Errorf("expected Grid to implement 'CellGetter', got %T", g)
		}
		if _, ok := g.(golife.CellCounter); !ok {
			t.Errorf("expected Grid to implement 'CellCounter', got %T", g)
		}
	})
	t.Run("make sure the grid reports missing on invalid point call", func(t *testing.T) {
		g, _, _ := golife.NewGrid(4)

		if _, ok := g.IsAlive(golife.Point{5, 5}); ok {
			t.Errorf("expected Grid to implement 'CellTester'm got %T", g)
		}
	})
	t.Run("SetCell must work and actually set a cell", func(t *testing.T) {
		g, _, _ := golife.NewGrid(4)

		if err := g.SetCell(golife.Point{1, 1}, golife.Cell{true}); err != nil {
			t.Errorf("SetCell cannot return an error but instead got %s", err)
		}
		if l, ok := g.IsAlive(golife.Point{1, 1}); !ok || !l {
			t.Errorf("the cell at point %v should be alive, it is not", golife.Point{1, 1})
		}
	})
}

var (
	bP = golife.Point{1, 1}
	bC = golife.Cell{false}
)

func BenchmarkSiblingsAlive16(b *testing.B) {
	g, _, _ := golife.NewGrid(16)
	for n := 0; n < b.N; n++ {
		golife.SiblingsAlive(bP, bC, g)
	}
}

func BenchmarkSiblingsAlive1000(b *testing.B) {
	g, _, _ := golife.NewGrid(1000)
	for n := 0; n < b.N; n++ {
		golife.SiblingsAlive(bP, bC, g)
	}
}

func BenchmarkSetCellState16(b *testing.B) {
	g, _, _ := golife.NewGrid(16)
	for n := 0; n < b.N; n++ {
		golife.GetCellState(bP, bC, g)
	}
}

func BenchmarkSetCellState1000(b *testing.B) {
	g, _, _ := golife.NewGrid(1000)
	for n := 0; n < b.N; n++ {
		golife.GetCellState(bP, bC, g)
	}
}

func ExampleNewGrid() {
	g, n, err := golife.NewGrid(200)

	alive, exists := g.IsAlive(golife.Point{10, 10})
	fmt.Println(alive)  // false
	fmt.Println(exists) // true

	fmt.Println(n) // 196 cells are created (sqrt(int(200)) * 2)

	fmt.Println(err) // nil
}
