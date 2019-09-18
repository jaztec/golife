package life_test

import (
	"testing"

	"github.com/jaztec/golife/pkg/life"
)

func TestNewGrid(t *testing.T) {
	var expect int32 = 196
	t.Run("test exact number of cells", func(t *testing.T) {
		_, got := life.NewGrid(expect)

		if got != expect {
			t.Errorf("expected %d cells but %d where created", expect, got)
		}
	})
	t.Run("test calculated number of cells", func(t *testing.T) {
		_, got := life.NewGrid(200)

		if got != expect {
			t.Errorf("expected %d cells but %d where created", expect, got)
		}
	})
	t.Run("make sure the grid implements the proper interfaces", func(t *testing.T) {
		var g interface{}
		g, _ = life.NewGrid(4)

		if _, ok := g.(life.CellTester); !ok {
			t.Errorf("expected Grid to implement 'CellTester'm got %T", g)
		}
	})
}

func BenchmarkSiblingsAlive16(b *testing.B) {
	g, _ := life.NewGrid(16)
	for n := 0; n < b.N; n++ {
		life.SiblingsAlive(life.Point{1, 1}, life.Cell{false}, g)
	}
}

func BenchmarkSiblingsAlive1000(b *testing.B) {
	g, _ := life.NewGrid(1000)
	for n := 0; n < b.N; n++ {
		life.SiblingsAlive(life.Point{1, 1}, life.Cell{false}, g)
	}
}

func BenchmarkSetCellState16(b *testing.B) {
	g, _ := life.NewGrid(16)
	for n := 0; n < b.N; n++ {
		life.SetCellState(life.Point{1, 1}, life.Cell{false}, g)
	}
}

func BenchmarkSetCellState1000(b *testing.B) {
	g, _ := life.NewGrid(1000)
	for n := 0; n < b.N; n++ {
		life.SetCellState(life.Point{1, 1}, life.Cell{false}, g)
	}
}
