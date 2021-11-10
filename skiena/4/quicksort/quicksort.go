package quicksort

import "math/rand"

func Sort(src []float64) []float64 {
	result := make([]float64, len(src))
	copy(result, src)

	sort(result)

	return result
}

func sort(result []float64) {
	if len(result) < 2 {
		return
	}

	p := partition(result)

	sort(result[:p])
	sort(result[p+1:])
}

func partition(src []float64) int {
	firstHigh := 0

	pivot := rand.Int() % len(src)
	src[len(src)-1], src[pivot] = src[pivot], src[len(src)-1]

	for i := range src[:len(src)-1] {
		if src[i] < src[len(src)-1] {
			src[i], src[firstHigh] = src[firstHigh], src[i]
			firstHigh++
		}
	}

	src[len(src)-1], src[firstHigh] = src[firstHigh], src[len(src)-1]

	return firstHigh
}
