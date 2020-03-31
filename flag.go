package clicmdflags

const flagPattern = "-"

// Flag -
type Flag struct {
	Name        string
	Description string
	Type        string
	Default     string
}
