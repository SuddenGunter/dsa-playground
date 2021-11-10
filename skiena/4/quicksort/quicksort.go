package quicksort

import "math/rand"

func Sort(src []float64) []float64 {
	result := make([]float64, len(src))
	copy(result, src)

	sort(result)

	return result
}

func sort(result []float64) {
	p := partition(result)
	if p == -1 {
		return
	}

	sort(result[:p])
	sort(result[p+1:])
}

func partition(result []float64) int {
	if len(result) <= 1 {
		return -1
	}

	pivot := rand.Int() % len(result)

	firstHigh := len(result) - 1
	lastLow := 0
	for lastLow < firstHigh {
		if result[lastLow] > result[pivot] {
			for result[firstHigh] > result[pivot] {
				firstHigh--
			}
			if firstHigh > lastLow {
				result[lastLow], result[firstHigh] = result[firstHigh], result[lastLow]
			}
		}

		lastLow++
	}

	result[pivot], result[firstHigh] = result[firstHigh], result[pivot]

	return pivot
}
