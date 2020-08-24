package tictactoe

import "testing"

func TestReduce(t *testing.T) {
	state1 := State{
		X, O, Empty,
		X, O, Empty,
		O, X, Empty,
	}
	state2 := State{
		X, X, O,
		O, O, X,
		Empty, Empty, Empty,
	}
	state3 := State{
		Empty, X, X,
		O, O, X,
		Empty, Empty, O,
	}

	node1 := &Node{}

	t.Run("Equivalent state", func(t *testing.T) {
		NodeMap = map[State]*Node{}
		NodeMap[state1] = node1
		want := state1
		got := reduceState(state2)
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Non-equivalent state", func(t *testing.T) {
		NodeMap = map[State]*Node{}
		NodeMap[state1] = node1
		dontWant := state1
		got := reduceState(state3)
		if got == dontWant {
			t.Errorf("got %d, dont want %d", got, dontWant)
		}
	})
}
