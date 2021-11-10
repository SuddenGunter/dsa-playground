package quicksort

import (
	"math/rand"
	"sync"
)

func Sort(src []float64) []float64 {
	result := make([]float64, len(src))
	copy(result, src)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	sort(result, wg)
	wg.Wait()
	return result
}

func sort(result []float64, wg *sync.WaitGroup) {
	defer wg.Done()
	if len(result) < 2 {
		return
	}

	p := partition(result)

	wg.Add(2)
	go sort(result[:p], wg)
	go sort(result[p+1:], wg)
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
