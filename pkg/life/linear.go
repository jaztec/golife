package life

type linearSimulator struct {
	grid *Grid
	gen  int
}

func (bs *linearSimulator) Step() error {
	bs.gen++
	return nil
}

func (bs *linearSimulator) Gen() int {
	return bs.gen
}

func (bs *linearSimulator) SetGrid(g *Grid) error {
	bs.grid = g
	return nil
}

// NewLinearSimulator returns a new simulator that processes a
// by game linear processing
func NewLinearSimulator(cellCount int32) (Simulator, error) {
	g, _, err := NewGrid(cellCount)
	if err != nil {
		return nil, err
	}
	return &linearSimulator{
		grid: g,
		gen:  0,
	}, nil
}
