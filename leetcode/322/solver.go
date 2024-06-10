package solver

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	max := amount + 1
	for i := range dp {
		dp[i] = max
	}
	dp[0] = 0

	for i := 1; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = min(dp[i], 1+dp[i-coin])
			}
		}
	}

	if dp[amount] == max {
		return -1
	}

	return dp[amount]
}
