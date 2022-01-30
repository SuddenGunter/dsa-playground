package solver

const size = 3

func tictactoe(moves [][]int) string {
	return Solve(moves)
}

func Solve(moves [][]int) string {
	aMoves := true
	var state [size][size]rune

	emptyFound := false
	won := false
	winner := rune(0)

	for _, m := range moves {
		if aMoves {
			state[m[0]][m[1]] = 'A'
		} else {
			state[m[0]][m[1]] = 'B'
		}

		winner, won, emptyFound = checkWinners(state)
		//prettyPrint(state)
		//fmt.Println()
		if won {
			return string(winner)
		}

		aMoves = !aMoves
	}

	if emptyFound {
		return "Pending"
	}

	return "Draw"
}

func checkWinners(state [size][size]rune) (rune, bool, bool) {
	emptyElementFound := false

	// by row
	for i := 0; i < size; i++ {
		prev := state[i][0]
		if prev == rune(0) {
			emptyElementFound = true
			break
		}
		prevWon := true

		for j := 1; j < size; j++ {
			if prev != state[i][j] {
				prevWon = false
				break
			}
		}

		if prevWon {
			return prev, prevWon, emptyElementFound
		}
	}

	// by col
	for j := 0; j < size; j++ {
		prev := state[0][j]
		if prev == rune(0) {
			emptyElementFound = true
			break
		}
		prevWon := true

		for i := 1; i < size; i++ {
			if prev != state[i][j] {
				prevWon = false
				break
			}
		}

		if prevWon {
			return prev, prevWon, emptyElementFound
		}
	}

	// by diag top left -> bottom right
	j := 0
	i := 0
	prev := state[i][j]
	prevWon := true

	if prev != rune(0) {
		for i < size {
			if prev != state[i][j] {
				prevWon = false
				break
			}

			i++
			j++
		}
	} else {
		emptyElementFound = true
		prevWon = false
	}

	if prevWon {
		return prev, prevWon, emptyElementFound
	}

	// by diag bottom left -> top right
	j = 0
	i = 2
	prev = state[i][j]
	prevWon = true

	if prev != rune(0) {
		for i >= 0 {
			if prev != state[i][j] {
				prevWon = false
				break
			}

			i--
			j++
		}
	} else {
		emptyElementFound = true
		prevWon = false
	}

	if prevWon {
		return prev, prevWon, emptyElementFound
	}

	if !emptyElementFound {
		for i := range state {
			for j := range state[i] {
				if state[i][j] == rune(0) {
					emptyElementFound = true
					break
				}
			}
		}
	}

	return 0, false, emptyElementFound
}
