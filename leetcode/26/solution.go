package solution

func removeDuplicates(nums []int) int {
	moveTo := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[moveTo] {
			continue
		} else {
			moveTo++
			nums[moveTo] = nums[i]
		}
	}

	nums = nums[:moveTo+1]
	return len(nums)
}
