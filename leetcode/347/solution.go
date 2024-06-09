package solution

func topKFrequent(nums []int, k int) []int {
	numOccurrences := make(map[int]int)
	for _, v := range nums {
		numOccurrences[v]++
	}

	topK := make([][]int, len(nums))
	for k, v := range numOccurrences {
		topK[v-1] = append(topK[v-1], k)
	}

	result := make([]int, 0, k)
	for i := len(topK) - 1; i >= 0 && len(result) < k; i-- {
		if len(topK[i]) > 0 {
			result = append(result, topK[i]...)
		}
	}

	return result[:k]
}
