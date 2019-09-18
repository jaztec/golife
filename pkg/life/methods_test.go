package life

import (
	"testing"
)

func TestGrid(t *testing.T) {
	t.Run("create new grid", func(t *testing.T) {
		var expect int32 = 196
		t.Run("test exact number of cells", func(t *testing.T) {
			g, got := NewGrid(expect)

			if got != expect {
				t.Errorf("expected %d cells but %d where created", expect, got)
			}

			if len(g.set) != int(expect) {
				t.Errorf("expected %d cells but %d where created", expect, len(g.set))
			}
		})
		t.Run("test calculated number of cells", func(t *testing.T) {
			g, got := NewGrid(200)

			if got != expect {
				t.Errorf("expected %d cells but %d where created", expect, got)
			}

			if len(g.set) != int(expect) {
				t.Errorf("expected %d cells but %d where created", expect, len(g.set))
			}
		})
	})
	t.Run("test grid coords", func(t *testing.T) {
		g, _ := NewGrid(196)
		coords := []Point{
			{0, 0},
			{2, 13},
			{4, 11},
			{7, 4},
			{7, 10},
			{8, 9},
			{9, 1},
			{10, 6},
			{10, 8},
			{13, 13},
		}
		for _, p := range coords {
			if _, ok := g.set[p]; !ok {
				t.Errorf("a Point was expected at position %v, none found", p)
			}
		}
	})
}

func isAlive(i, x int) bool {
	if i <= x {
		return true
	}
	return false
}
func TestSiblings(t *testing.T) {
	type tC struct {
		alive int
		cells map[Point]Cell
	}
	testPoint := Point{1, 1}
	grids := make([]tC, 8)
	for i := 0; i < 8; i++ {
		tc := tC{alive: i, cells: make(map[Point]Cell, 9)}
		tc.cells[testPoint] = Cell{false}

		// generate some cell values
		tc.cells[Point{0, 0}] = Cell{isAlive(1, i)}
		tc.cells[Point{0, 1}] = Cell{isAlive(2, i)}
		tc.cells[Point{0, 2}] = Cell{isAlive(3, i)}
		tc.cells[Point{1, 0}] = Cell{isAlive(4, i)}
		tc.cells[Point{1, 2}] = Cell{isAlive(5, i)}
		tc.cells[Point{2, 0}] = Cell{isAlive(6, i)}
		tc.cells[Point{2, 1}] = Cell{isAlive(7, i)}
		tc.cells[Point{2, 2}] = Cell{isAlive(8, i)}

		grids[i] = tc
	}
	t.Run("test siblings alive", func(t *testing.T) {
		for _, tc := range grids {
			g := &Grid{tc.cells}
			c := tc.cells[testPoint]
			got := SiblingsAlive(testPoint, c, g)
			if got != tc.alive {
				t.Errorf("expected %d cells to be alive but found %d", tc.alive, got)
			}
		}
	})
	t.Run("test set cell state", func(t *testing.T) {
		for _, tc := range grids {
			g := &Grid{tc.cells}
			c := tc.cells[testPoint]

			c = SetCellState(testPoint, c, g)
			if (tc.alive < 2 || tc.alive > 3) && c.Alive {
				t.Errorf("expected cell to be dead with %d life siblings but it lives", tc.alive)
			}
			if tc.alive > 2 && tc.alive < 4 && !c.Alive {
				t.Errorf("expected cell to be alive with %d life siblings but it is not", tc.alive)
			}
		}
	})
}

func BenchmarkGenerateEmptyMap1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateEmptyMap(1000, 1000)
	}
}

func BenchmarkGenerateEmptyMap100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateEmptyMap(100, 100)
	}
}

func BenchmarkGenerateEmptyMap10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateEmptyMap(10, 10)
	}
}

func BenchmarkSiblingsAlive16(b *testing.B) {
	g, _ := NewGrid(16)
	for n := 0; n < b.N; n++ {
		SiblingsAlive(Point{1, 1}, Cell{false}, g)
	}
}

func BenchmarkSiblingsAlive1000(b *testing.B) {
	g, _ := NewGrid(1000)
	for n := 0; n < b.N; n++ {
		SiblingsAlive(Point{1, 1}, Cell{false}, g)
	}
}

func BenchmarkSetCellState16(b *testing.B) {
	g, _ := NewGrid(16)
	for n := 0; n < b.N; n++ {
		SetCellState(Point{1, 1}, Cell{false}, g)
	}
}

func BenchmarkSetCellState1000(b *testing.B) {
	g, _ := NewGrid(1000)
	for n := 0; n < b.N; n++ {
		SetCellState(Point{1, 1}, Cell{false}, g)
	}
}
