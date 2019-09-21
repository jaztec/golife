package golife

import (
	"fmt"
	"testing"
)

func TestBasicSimulatorBases(t *testing.T) {
	s, err := NewLinearSimulator(16)
	if err != nil {
		t.Fatalf("no error may occur on creating a new linear simulator")
	}
	// check base values
	if g := s.Gen(); g != 0 {
		t.Errorf("base value for Gen should be 0 instead of %d", g)
	}
	if err := s.Step(); err != nil {
		t.Errorf("base value for Step cannot return an error, %s", err)
	}
	if g := s.Gen(); g != 1 {
		t.Errorf("after a step the generation should be 1 instead of %d", g)
	}
	g, _, _ := NewGrid(4)
	if err := s.SetGrid(g); err != nil {
		t.Errorf("SetGrid cannot produce an error, instead got %s", err)
	}
}

func TestLinearSimulatorSteps(t *testing.T) {
	s, _ := NewLinearSimulator(36)
	g := simpleGliderSetup()
	s.SetGrid(g)
	// validate step 1
	s.Step()
	if err := compareGrids(s.Grid(), simpleGliderAtStep(stepOne)); err != nil {
		t.Errorf("error at step 1: %v", err)
	}
	// validate step 2
	s.Step()
	if err := compareGrids(s.Grid(), simpleGliderAtStep(stepTwo)); err != nil {
		t.Errorf("error at step 2: %v", err)
	}
	// validate step 3
	s.Step()
	if err := compareGrids(s.Grid(), simpleGliderAtStep(stepThree)); err != nil {
		t.Errorf("error at step 3: %v", err)
	}
}

func compareGrids(a, b CellGetter) (equal error) {
	aC := a.GetCells()
	bC := b.GetCells()

	if len(aC) != len(bC) {
		return fmt.Errorf("grid sizes don't equal, got %d but want %d", len(aC), len(bC))
	}

	missed := ""
	counter := 0
	for p, c := range aC {
		cB, _ := b.Cell(p)
		if c != cB {
			missed += fmt.Sprintf("%v: (%v != %v),\n", p, c, cB)
			counter++
		}
	}

	if counter > 0 {
		return fmt.Errorf("mismatch between cells (%d):\n%s", counter, missed)
	}

	return nil
}
