package solver

func maxProfit(prices []int) int {
	currentWindowBuy := prices[0]
	currentWindowSell := prices[len(prices)-1]
	bestDeal := currentWindowSell - currentWindowBuy

	for _, x := range prices[1:] {
		if x < currentWindowBuy && currentWindowSell-x > bestDeal {
			currentWindowBuy = x
			continue
		}

		if x-currentWindowBuy > bestDeal {
			bestDeal = x - currentWindowBuy
			currentWindowSell = x
		}
	}

	if bestDeal < 0 {
		return 0
	}

	return bestDeal
}
