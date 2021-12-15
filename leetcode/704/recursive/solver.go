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
	return binarySearch(nums, target, 0, len(nums)-1)
}

func binarySearch(nums []int, target, low, high int) int {
	if low > high {
		return -1
	}

	middle := (low + high) / 2

	if nums[middle] > target {
		return binarySearch(nums, target, low, middle-1)
	}
	if nums[middle] < target {
		return binarySearch(nums, target, middle+1, high)
	}

	return middle
}

func Solve(nums []int, target int) int {
	return search(nums, target)
}
