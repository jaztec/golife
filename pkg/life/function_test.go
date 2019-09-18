package life

import (
	"testing"
)

func TestGrid(t *testing.T) {
	t.Run("create new grid", func(t *testing.T) {
		var expect int32 = 196
		t.Run("test exact number of cells", func(t *testing.T) {
			g, got, _ := NewGrid(expect)

			if got != expect {
				t.Errorf("expected %d cells but %d where created", expect, got)
			}

			if len(g.set) != int(expect) {
				t.Errorf("expected %d cells but %d where created", expect, len(g.set))
			}
		})
		t.Run("test calculated number of cells", func(t *testing.T) {
			g, got, _ := NewGrid(200)

			if got != expect {
				t.Errorf("expected %d cells but %d where created", expect, got)
			}

			if len(g.set) != int(expect) {
				t.Errorf("expected %d cells but %d where created", expect, len(g.set))
			}
		})
	})
	t.Run("test grid coords", func(t *testing.T) {
		g, _, _ := NewGrid(196)
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
			if l, _ := g.IsAlive(p); l == true {
				t.Errorf("the point at position %v shoould not be alive", p)
			}
		}
	})
}

func setAlive(i, x int) bool {
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
		tc.cells[Point{0, 0}] = Cell{setAlive(1, i)}
		tc.cells[Point{0, 1}] = Cell{setAlive(2, i)}
		tc.cells[Point{0, 2}] = Cell{setAlive(3, i)}
		tc.cells[Point{1, 0}] = Cell{setAlive(4, i)}
		tc.cells[Point{1, 2}] = Cell{setAlive(5, i)}
		tc.cells[Point{2, 0}] = Cell{setAlive(6, i)}
		tc.cells[Point{2, 1}] = Cell{setAlive(7, i)}
		tc.cells[Point{2, 2}] = Cell{setAlive(8, i)}

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

func BenchmarkGenerateEmptyMap5000x5000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateEmptyMap(5000, 5000)
	}
}

func BenchmarkGenerateEmptyMap2500x2500(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateEmptyMap(2500, 2500)
	}
}

func BenchmarkGenerateEmptyMap2000x2000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateEmptyMap(2000, 2000)
	}
}

func BenchmarkGenerateEmptyMap1000x1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateEmptyMap(1000, 1000)
	}
}

func BenchmarkGenerateEmptyMap100x100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateEmptyMap(100, 100)
	}
}

func BenchmarkGenerateEmptyMap10x10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateEmptyMap(10, 10)
	}
}
