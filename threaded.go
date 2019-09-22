package golife

import (
	"fmt"
)

type msg struct {
	c  Cell
	p  Point
	ct CellTester
}

type threadedSimulator struct {
	grid *Grid
	gen  int
	in   chan msg
	out  chan msg
	done chan struct{}
}

func (ts *threadedSimulator) Step() error {
	g, n, err := NewGrid(ts.grid.Count())
	if n != ts.Grid().Count() {
		return fmt.Errorf("old size %d does not match new size %d", n, ts.Grid().Count())
	}
	if err != nil {
		return err
	}

	for p, c := range ts.grid.set {
		ts.in <- msg{c, p, ts.grid}
	}

	var i int32
	for i = 0; i < ts.grid.Count(); i++ {
		select {
		case msg := <-ts.out:
			g.SetCell(msg.p, msg.c)
		}
	}

	ts.grid = g
	ts.gen++
	return nil
}

func (ts *threadedSimulator) Gen() int {
	return ts.gen
}

func (ts *threadedSimulator) Grid() *Grid {
	return ts.grid
}

func (ts *threadedSimulator) SetGrid(g *Grid) error {
	ts.grid = g
	return nil
}

func startWorker(in <-chan msg, out chan<- msg, done <-chan struct{}) {
	for {
		select {
		case m := <-in:
			m.c = GetCellState(m.p, m.c, m.ct)
			out <- m
		case <-done:
			return
		}
	}
}

// NewThreadedSimulator returns a new simulator that processes a
// by game threaded processing
func NewThreadedSimulator(cellCount, threadCount int32) (Simulator, error) {
	g, _, err := NewGrid(cellCount)
	if err != nil {
		return nil, err
	}
	ts := &threadedSimulator{
		grid: g,
		gen:  0,
		in:   make(chan msg, cellCount),
		out:  make(chan msg, cellCount),
		done: make(chan struct{}, 1),
	}
	var i int32
	for i = 0; i < threadCount; i++ {
		go startWorker(ts.in, ts.out, ts.done)
	}
	return ts, nil
}
