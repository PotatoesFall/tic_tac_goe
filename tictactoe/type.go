package tictactoe

// Field is a value in a tictactoe field
type Field uint8

// possible field values, E is empty
const (
	Empty = iota
	X
	O
)

// State is the state of a tictactoe game
type State [9]Field

// Node is a node of the game tree
type Node struct {
	Outcome
	State
	Choices []Choice
	Depth   uint8
}

// Choice is a weighted Node pointer
type Choice struct {
	Weight uint8
	Node   *Node
}

// Outcome is whether a game is won by either party
type Outcome uint8

// possible outcomes
const (
	None = iota
	XWin
	OWin
)
