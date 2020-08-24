package main

import (
	"fmt"

	"github.com/PotatoesFall/tic_tac_goe/tictactoe"
)

const (
	playingRounds  = 500000
	trainingRounds = 5000000
	start          = true
)

func main() {
	parent := tictactoe.MakeTree()

	for i := uint8(1); i <= 9; i++ {
		fmt.Printf("step %d: %d possibilities\n", i, tictactoe.Nums[i])
	}

	fmt.Println()
	fmt.Println("Pre-training round:")
	outcomes := tictactoe.CountOutcomes(tictactoe.Play(parent, playingRounds, start))
	draw, xwin, owin := calcPercent(outcomes)
	fmt.Printf("%%%2d Draw, %%%2d X, %%%2d O\n", draw, xwin, owin)

	fmt.Println()
	fmt.Printf("Doing %d training rounds...\n", trainingRounds)
	tictactoe.Train(parent, trainingRounds)

	fmt.Println()
	fmt.Println("Post-training round:")
	outcomes = tictactoe.CountOutcomes(tictactoe.Play(parent, playingRounds, start))
	draw, xwin, owin = calcPercent(outcomes)
	fmt.Printf("%%%2d Draw, %%%2d X, %%%2d O\n", draw, xwin, owin)
}

func calcPercent(outcomes [3]int) (int, int, int) {
	sum := outcomes[0] + outcomes[1] + outcomes[2]
	return 100 * outcomes[0] / sum, 100 * outcomes[1] / sum, 100 * outcomes[2] / sum
}
