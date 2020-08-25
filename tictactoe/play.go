package tictactoe

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var playState State

// Play play against the computer in the command line
func Play(parent Node) {
	rand.Seed(time.Now().UTC().UnixNano())

	input := bufio.NewReader(os.Stdin)

	userWins := 0
	pcWins := 0
	ties := 0

	nextRound := true
	userStarts := true

	for nextRound {
		switch userStarts {
		case true:
			fmt.Println("You start this round!")
		case false:
			fmt.Println("This round the computer starts.")
		}

		curNode := &parent
		playState = parent.State
		for step := 0; step < 9; step++ {
			if curNode.Outcome != None || len(curNode.Choices) == 0 {
				break
			}
			var choice *Choice
			if step%2 == 0 != userStarts {
				choice = makeChoice(curNode)
			} else {
				PrintState(curNode.State)
				choice = makeUserChoice(curNode, input, getField(step))
			}
			curNode = choice.Node
		}

		PrintState(curNode.State)
		switch {
		case curNode.Outcome == None:
			ties++
			fmt.Println("It's a tie.")
		case curNode.Outcome == XWin == userStarts:
			userWins++
			fmt.Println("You win!")
		case curNode.Outcome == XWin != userStarts:
			pcWins++
			fmt.Println("You lose...")
		}

		fmt.Print("Play another round? [Y/n] ")
		another, err := input.ReadString('\n')
		another = strings.Trim(another, " \n\t")
		if err != nil {
			panic(err)
		}
		if strings.ToLower(another) != "y" {
			break
		}
		userStarts = !userStarts
	}

	fmt.Printf("You won %d games, lost %d and tied %d.\n", userWins, pcWins, ties)
	fmt.Println("Thank you for playing!")
}

func makeUserChoice(node *Node, input *bufio.Reader, nextField Field) *Choice {
	move := 0
	for move > 9 || move < 1 || node.State[move-1] != Empty {
		fmt.Println("You are " + string(getFieldRune(nextField)))
		fmt.Print("What is your next move [1-9]? ")
		moveStr, _ := input.ReadString('\n')
		move, _ = strconv.Atoi(strings.Trim(moveStr, " \r\n\t"))
	}

	newState := node.State
	newState[move-1] = nextField
	newState = reduceState(newState)

	return &Choice{
		Node: NodeMap[newState],
	}
}

func getField(step int) Field {
	if step%2 == 0 {
		return X
	}
	return O
}
