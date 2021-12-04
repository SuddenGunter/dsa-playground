package solution

import "math"

func FindMedianSortedArrays(x, y []int) float64 {
	if len(x) > len(y) {
		y, x = x, y
	}

	start := 0
	end := len(x)

	partitionX := (start + end) / 2
	partitionY := (len(x)+len(y)+1)/2 - partitionX

	for !foundMidpoint(x, y, partitionX, partitionY) {
		if maxLeft(x, partitionX) > minRight(y, partitionY) {
			end = partitionX - 1
		} else {
			start = partitionX + 1
		}

		partitionX = (start + end) / 2
		partitionY = (len(x)+len(y)+1)/2 - partitionX
	}

	switch (len(x) + len(y)) % 2 {
	case 0:
		return (max(maxLeft(x, partitionX), maxLeft(y, partitionY)) + min(minRight(x, partitionX), minRight(y, partitionY))) / 2
	case 1:
		return max(maxLeft(x, partitionX), maxLeft(y, partitionY))
	}

	return 0
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}

	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}

	return b
}

func foundMidpoint(x, y []int, partitionX, partitionY int) bool {
	return maxLeft(x, partitionX) <= minRight(y, partitionY) && maxLeft(y, partitionY) <= minRight(x, partitionX)
}

func maxLeft(arr []int, partition int) float64 {
	if partition == 0 {
		return math.Inf(-1)
	}

	return float64(arr[partition-1])
}

func minRight(arr []int, partition int) float64 {
	if partition == len(arr) {
		return math.Inf(1)
	}

	return float64(arr[partition])
}
