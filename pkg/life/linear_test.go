package life

import "testing"

func TestBasicSimulator(t *testing.T) {
	s, err := NewLinearSimulator()
	if err != nil {
		t.Fatalf("no error may occur on creating a new linear simulator")
	}
	// check base values
	if err := s.Step(); err != nil {
		t.Errorf("base value for Step cannot return an error, %s", err)
	}
	if g := s.Gen(); g != 0 {
		t.Errorf("base value for Gen should be 0 instead of %d", g)
	}
}
