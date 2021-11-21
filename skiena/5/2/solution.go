package solution

func FindMissing(src []int) int {
	if len(src) == 0 {
		panic("invalid input")
	}

	low, high := 0, len(src)-1
	for {
		mid := (low + high) / 2

		if src[mid]+1 != src[mid+1] {
			return src[mid] + 1
		}
		if src[mid]-1 != src[mid-1] {
			return src[mid] - 1
		}

		if src[mid] <= src[high]/2 {
			// check right subarray

			low = mid + 1

		} else {

			// check left subarray
			high = mid - 1
		}
	}
}
