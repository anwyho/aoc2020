package states

// Represents a state in the state grid
type State byte

const (
	Floor    State = '.'
	Empty    State = 'L'
	Occupied State = '#'
)

func (s State) String() string {
	switch s {
	case Floor:
		return "."
	case Empty:
		return "L"
	case Occupied:
		return "#"
	default:
		return "?"
	}
}

