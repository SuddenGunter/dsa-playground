package solution

import "math"

func MaxSubArraySum(nums []int) int {
	bestSum := float64(math.MinInt64)
	currentSum := float64(math.MinInt64)
	for _, x := range nums {
		currentSum = math.Max(float64(x), currentSum+float64(x))
		bestSum = math.Max(bestSum, currentSum)
	}

	return int(bestSum)
}
