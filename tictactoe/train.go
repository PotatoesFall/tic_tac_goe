package tictactoe

import (
	"errors"
	"math/rand"
	"time"
)

const (
	smallReward = 5
	reward      = 30
	punishment  = 20
)

// Train plays random games and modifies the weights
func Train(parent Node, n int) {
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < n; i++ {
		curNode := &parent
		choices := []*Choice{}

		// play up to 9 rounds
		for j := 0; j < 9; j++ {
			if curNode.Outcome != None || len(curNode.Choices) == 0 {
				break
			}
			choice := makeChoice(curNode)
			choices = append(choices, choice)
			curNode = choice.Node
		}
		// change AI around
		switch curNode.Outcome {
		case None:
			for _, choice := range choices {
				if choice.Weight < 255-smallReward {
					choice.Weight += smallReward
				} else {
					choice.Weight = 255
				}
			}
		case XWin:
			for i, choice := range choices {
				if i%2 == 0 {
					if choice.Weight <= 255-reward {
						choice.Weight += reward
					} else {
						choice.Weight = 255
					}
				} else {
					if choice.Weight > punishment {
						choice.Weight -= punishment
					} else {
						choice.Weight = 1
					}
				}
			}
		case OWin:
			for i, choice := range choices {
				if i%2 != 0 {
					if choice.Weight <= 255-reward {
						choice.Weight += reward
					} else {
						choice.Weight = 255
					}
				} else {
					if choice.Weight > punishment {
						choice.Weight -= punishment
					} else {
						choice.Weight = 1
					}
				}
			}
		}
	}
}

func makeChoice(node *Node) *Choice {
	sum := 0
	for _, choice := range node.Choices {
		sum += int(choice.Weight)
	}
	randInt := rand.Intn(sum)
	sum = 0
	for i, choice := range node.Choices {
		sum += int(choice.Weight)
		if sum > randInt {
			return &node.Choices[i]
		}
	}
	panic(errors.New("unable to make a choice"))
}

// CountOutcomes counts X wins, Y wins and ties
func CountOutcomes(outcomes []Outcome) [3]int {
	var counts [3]int
	for _, outcome := range outcomes {
		counts[outcome]++
	}
	return counts
}
