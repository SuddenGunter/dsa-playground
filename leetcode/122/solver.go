package solver

func maxProfit(prices []int) int {
	totalProfit := 0
	prevVal := prices[0]
	for _, x := range prices[1:] {
		if x > prevVal {
			totalProfit += x - prevVal
		}
		prevVal = x
	}

	return totalProfit
}
