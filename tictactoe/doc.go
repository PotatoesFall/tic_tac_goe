package tictactoe

import (
	"fmt"
	"os"
)

// MakeDoc creates a document with the ai output
func MakeDoc(parent Node, docName string) error {
	f, err := os.Create(docName)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintln(f, "L,M,R,,,L,M,R,,,L,M,R,,,L,M,R,,,L,M,R,,,L,M,R,,,L,M,R,,,L,M,R,,,L,M,R,,,L,M,R")
	if err != nil {
		return err
	}
	for state, node := range NodeMap {
		states := []State{}
		weights := []uint8{}
		for _, choice := range node.Choices {
			states = append(states, choice.Node.State)
			weights = append(weights, choice.Weight)
		}
		fmt.Fprintln(f)
		fmt.Fprintln(f, makeLine(0, state, states, weights))
		fmt.Fprintln(f, makeLine(1, state, states, weights))
		fmt.Fprintln(f, makeLine(2, state, states, weights))
	}

	return nil
}

func makeLine(i int, state State, states []State, weights []uint8) string {
	out := getCSVString(state[0+3*i], state[1+3*i], state[2+3*i])
	for _, state := range states {
		out += getCSVString(state[0+3*i], state[1+3*i], state[2+3*i])
	}
	return out
}

func getCSVString(fields ...Field) string {
	out := ""
	for _, field := range fields {
		switch field {
		case Empty:
			out += "_,"
		case X:
			out += "X,"
		case O:
			out += "O,"
		}
	}
	return out + ",,"
}
