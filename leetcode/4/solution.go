package solution

func FindMedianSortedArrays(x []int, y []int) float64 {
	if len(x) > len(y) {
		y, x = x, y
	}

	start := 0
	end := len(x) - 1

	partitionX := (start + end) / 2
	partitionY := (len(x)+len(y)+1)/2 - partitionX

	for !foundMidpoint(x, y, partitionX, partitionY) {
		start = partitionX + 1
		end = len(x) - 1

		partitionX = (start + end) / 2
		partitionY = (len(x)+len(y)+1)/2 - partitionX
	}

	switch (len(x) + len(y)) % 2 {
	case 0:
		return float64(max(x[partitionX-1], y[partitionY-1])+min(x[partitionX], y[partitionY])) / 2
	case 1:
		return float64(max(x[partitionX-1], y[partitionY-1]))
	}

	return 0
}

func max(a, b int) int {
	if a > b {
		return int(a)
	}

	return int(b)
}

func min(a, b int) int {
	if a < b {
		return int(a)
	}

	return int(b)
}

func foundMidpoint(x []int, y []int, partitionX int, partitionY int) bool {
	return x[partitionX-1] <= y[partitionY] && y[partitionY-1] <= x[partitionX]
}
