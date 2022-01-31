package solution

func maximumWealth(accounts [][]int) int {
	max := accounts[0][0]

	sum := 0
	for _, customer := range accounts {
		sum = 0
		for _, bank := range customer {
			sum += bank
		}

		if sum >= max {
			max = sum
		}
	}

	return max
}
