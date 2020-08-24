package tictactoe

import "fmt"

// statistical variables for debugging
var Nums map[uint8]int

// NodeMap contains all nodes to prevent duplicate nodes
var NodeMap map[State]*Node

const defaultWeight = 50

// MakeTree makes a full tree and
func MakeTree() Node {
	NodeMap = map[State]*Node{}
	Nums = map[uint8]int{}
	parent := Node{}
	NodeMap[parent.State] = &parent
	makeChildren(&parent)
	return parent
}

// PrintState prints a state to Stdout
func PrintState(state State) {
	fmt.Println("_______")
	for i := 0; i < 9; i += 3 {
		fmt.Printf("|%c %c %c|\n", getField(state[i]), getField(state[i+1]), getField(state[i+2]))
	}
	fmt.Println("¯¯¯¯¯¯¯")
}

func makeChildren(node *Node) {
	node.Choices = []Choice{}

	newDepth := node.Depth + 1
	newField := Field(node.Depth%2 + 1)

	for i, field := range node.State {
		// only make children for empty fields
		if field != None {
			continue
		}

		// set new state
		newState := node.State
		newState[i] = newField

		// make a choice and get or create node for new state
		choice := Choice{
			Weight: defaultWeight,
			Node:   NodeMap[newState],
		}
		if choice.Node == nil {
			choice.Node = &Node{
				Outcome: getOutcome(newState),
				State:   newState,
				Depth:   newDepth,
			}
			NodeMap[choice.Node.State] = choice.Node
			Nums[newDepth]++
			// go deeper
			if choice.Node.Outcome == None && newDepth != 9 {
				makeChildren(choice.Node)
			}
		}

		// add choice to node
		node.Choices = append(node.Choices, choice)
	}
}

// Helper for PrintState
func getField(field Field) rune {
	switch field {
	case Empty:
		return '·'
	case X:
		return 'X'
	case O:
		return 'O'
	}

	return ' '
}

func getOutcome(state State) Outcome {
	// check rows
	for i := 0; i < 9; i += 3 {
		if state[i] == state[i+1] && state[i] == state[i+2] {
			return Outcome(state[i])
		}
	}
	// check cols
	for i := 0; i < 3; i++ {
		if state[i] == state[i+3] && state[i] == state[i+6] {
			return Outcome(state[i])
		}
	}

	// check diagonal
	if (state[0] == state[4] && state[0] == state[8]) ||
		state[2] == state[4] && state[2] == state[6] {
		return Outcome(state[4])
	}

	return None
}
