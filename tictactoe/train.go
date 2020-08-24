package tictactoe

import (
	"errors"
	"math/rand"
	"time"
)

const (
	smallReward = 2
	reward      = 1
	punishment  = 10
)

// Train plays random games and modifies the weights
func Train(parent Node, n int, start bool) {
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < n; i++ {
		curNode := &parent
		choices := []*Choice{}

		// play up to 9 rounds
		for j := 0; j < 9; j++ {
			if curNode.Outcome != None || len(curNode.Choices) == 0 {
				break
			}
			var choice *Choice
			if j%2 == 0 != start {
				choice = makeChoice(curNode)
				choices = append(choices, choice)
			} else {
				choice = makeRandomChoice(curNode)
			}
			curNode = choice.Node
		}
		// use feedback
		switch {
		case curNode.Outcome == None:
			for _, choice := range choices {
				if choice.Weight <= 255-smallReward {
					choice.Weight += smallReward
				} else {
					choice.Weight = 255
				}
			}
		case curNode.Outcome == XWin == start || curNode.Outcome == OWin != start:
			for _, choice := range choices {
				if choice.Weight >= punishment {
					choice.Weight -= punishment
				} else {
					choice.Weight = 0
				}
			}
		case curNode.Outcome == XWin != start || curNode.Outcome == OWin == start:
			for _, choice := range choices {
				if choice.Weight <= 255-reward {
					choice.Weight += reward
				} else {
					choice.Weight = 255
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
	if sum == 0 {
		return makeRandomChoice(node)
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
