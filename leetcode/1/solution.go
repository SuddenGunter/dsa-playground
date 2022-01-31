package solution

func twoSum(nums []int, target int) []int {
	toFind := make(map[int]int, len(nums))

	for i, n := range nums {
		x, ok := toFind[n]
		if ok {
			return []int{x, i}
		}

		toFind[target-n] = i
	}

	return nil
}
