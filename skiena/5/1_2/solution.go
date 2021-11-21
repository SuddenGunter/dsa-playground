package solution

func FindMax(src []int) int {
	if len(src) == 0 {
		panic("invalid input")
	}

	if len(src) < 2 {
		return src[0]
	}

	// not shifted
	if src[0] < src[len(src)-1] {
		return src[0]
	}

	low, high := 0, len(src)-1
	for {
		mid := (low + high) / 2

		// Transition lies in left half of the array
		if src[mid] < src[high] {
			if src[mid-1] > src[mid] {
				return src[mid]
			}

			high = mid - 1
			continue
		}

		// Transition lies in right half of the array
		if src[mid] > src[mid+1] {
			if src[mid-1] > src[mid] {
				return src[mid]
			}

			high = mid - 1
			continue
		}

		if src[mid] > src[mid+1] {
			return mid + 1
		}

		low = mid + 1
	}
}
