package golife

import (
	"testing"
)

func TestThreadedSimulatorBases(t *testing.T) {
	s, err := NewThreadedSimulator(16, 8)
	if err != nil {
		t.Fatalf("no error may occur on creating a new threaded simulator")
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

func TestThreadedSimulatorSteps(t *testing.T) {
	s, _ := NewThreadedSimulator(36, 8)
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

func BenchmarkThreadedSimulator100_100(b *testing.B) {
	s, _ := NewThreadedSimulator(100, 8)
	for n := 0; n < b.N; n++ {
		for i := 0; i < 100; i++ {
			s.Step()
		}
	}
}

func BenchmarkThreadedSimulator1000_100(b *testing.B) {
	s, _ := NewThreadedSimulator(1000, 8)
	for n := 0; n < b.N; n++ {
		for i := 0; i < 100; i++ {
			s.Step()
		}
	}
}

func BenchmarkThreadedSimulator1000_1(b *testing.B) {
	s, _ := NewThreadedSimulator(1000, 8)
	for n := 0; n < b.N; n++ {
		s.Step()
	}
}

// func BenchmarkThreadedSimulator10000_100(b *testing.B) {
// 	s, _ := NewThreadedSimulator(10000, 8)
// 	for n := 0; n < b.N; n++ {
// 		for i := 0; i < 100; i++ {
// 			s.Step()
// 		}
// 	}
// }

// func BenchmarkThreadedSimulator1000000_100(b *testing.B) {
// 	s, _ := NewThreadedSimulator(1000000, 8)
// 	for n := 0; n < b.N; n++ {
// 		for i := 0; i < 100; i++ {
// 			s.Step()
// 		}
// 	}
// }

// func BenchmarkThreadedSimulator1000000_100x16(b *testing.B) {
// 	s, _ := NewThreadedSimulator(1000000, 16)
// 	for n := 0; n < b.N; n++ {
// 		for i := 0; i < 100; i++ {
// 			s.Step()
// 		}
// 	}
// }
