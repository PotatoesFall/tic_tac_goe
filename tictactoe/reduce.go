package tictactoe

func reduceState(state State) State {
	for i := 0; i < 4; i++ {
		if _, exists := NodeMap[state]; exists {
			return state
		}
		state = rotate(state)
	}

	state = flip(state)

	for i := 0; i < 4; i++ {
		if _, exists := NodeMap[state]; exists {
			return state
		}
		state = rotate(state)
	}

	return state
}

func rotate(state State) State {
	state[0], state[2], state[8], state[6] = state[6], state[0], state[2], state[8]
	state[1], state[5], state[7], state[3] = state[3], state[1], state[5], state[7]
	return state
}

func flip(state State) State {
	state[0], state[6] = state[6], state[0]
	state[1], state[7] = state[7], state[1]
	state[2], state[8] = state[8], state[2]
	return state
}
