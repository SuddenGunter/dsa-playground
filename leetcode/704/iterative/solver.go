package solver

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	low := 0
	high := len(nums) - 1

	for low <= high {
		middle := (low + high) / 2

		if nums[middle] > target {
			high = middle - 1
			continue
		}

		if nums[middle] < target {
			low = middle + 1
			continue
		}

		return middle
	}

	return -1
}

func Solve(nums []int, target int) int {
	return search(nums, target)
}
