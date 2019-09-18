package golife

import "testing"

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
