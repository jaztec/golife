package golife

import "sync"

type step int

var (
	stepOne   step = 1
	stepTwo   step = 2
	stepThree step = 3
)

func setup(c int, points map[Point]struct{}) *Grid {
	g := &Grid{set: make(map[Point]Cell, c), lock: sync.RWMutex{}}

	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			p := Point{int32(x), int32(y)}
			a := false
			if _, ok := points[p]; ok {
				a = true
			}
			g.set[p] = Cell{a}
		}
	}

	return g
}

func simpleGliderSetup() *Grid {
	points := map[Point]struct{}{
		Point{1, 0}: {},
		Point{2, 1}: {},
		Point{0, 2}: {},
		Point{1, 2}: {},
		Point{2, 2}: {},
	}

	return setup(16, points)
}

func simpleGliderAtStep(s step) *Grid {
	points := map[Point]struct{}{}

	switch {
	case s == stepOne:
		points = map[Point]struct{}{
			Point{0, 1}: {},
			Point{2, 1}: {},
			Point{1, 2}: {},
			Point{2, 2}: {},
			Point{1, 3}: {},
		}
	case s == stepTwo:
		points = map[Point]struct{}{
			Point{2, 1}: {},
			Point{0, 2}: {},
			Point{2, 2}: {},
			Point{1, 3}: {},
			Point{2, 3}: {},
		}
	case s == stepThree:
		points = map[Point]struct{}{
			Point{1, 1}: {},
			Point{2, 2}: {},
			Point{3, 2}: {},
			Point{1, 3}: {},
			Point{2, 3}: {},
		}
	}

	return setup(16, points)
}
