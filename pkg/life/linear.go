package life

type linearSimulator struct{}

func (bs *linearSimulator) Step() error { return nil }
func (bs *linearSimulator) Gen() int    { return 0 }

// NewLinearSimulator returns a new simulator that processes a
// by game linear processing
func NewLinearSimulator() (Simulator, error) {
	return &linearSimulator{}, nil
}
