package solution

func firstBadVersion(n int) int {
	if n == 1 {
		return 1
	}

	low := 1
	high := n

	return binarySearch(n, low, high)
}

func binarySearch(n int, low, high int) int {
	if high-low == 1 {
		if isBadVersion(low) {
			return low
		} else {
			return high
		}
	}

	middle := (low + high) / 2

	if isBadVersion(middle) {
		return binarySearch(n, low, middle)
	} else {
		return binarySearch(n, middle, high)
	}
}
