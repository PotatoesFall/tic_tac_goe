package tictactoe

import (
	"math/rand"
	"time"
)

// Play plays games without modifying the weights
func Play(parent Node, n int, start bool) []Outcome {
	rand.Seed(time.Now().UTC().UnixNano())

	outcomes := []Outcome{}

	for i := 0; i < n; i++ {
		curNode := &parent

		// play up to 9 rounds
		for j := 0; j < 9; j++ {
			if curNode.Outcome != None || len(curNode.Choices) == 0 {
				break
			}
			var choice *Choice
			if j%2 == 0 != start {
				choice = makeChoice(curNode)
			} else {
				choice = makeRandomChoice(curNode)
			}
			curNode = choice.Node
		}

		outcomes = append(outcomes, curNode.Outcome)
	}
	return outcomes
}

func makeRandomChoice(node *Node) *Choice {
	randInt := rand.Intn(len(node.Choices))
	return &node.Choices[randInt]
}
