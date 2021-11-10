package quicksort

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
	pivot := 0

	firstHigh := len(src) - 1
	lastLow := 0

	for lastLow < firstHigh {
		if src[lastLow] > src[pivot] {
			for firstHigh > lastLow {
				if src[firstHigh] < src[pivot] {
					src[lastLow], src[firstHigh] = src[firstHigh], src[lastLow]
					firstHigh--
					break
				}
				firstHigh--
			}
		} else {
			lastLow++
		}
	}

	if src[lastLow] < src[pivot] {
		src[lastLow], src[pivot] = src[pivot], src[lastLow]
		return lastLow
	}

	return pivot
}
