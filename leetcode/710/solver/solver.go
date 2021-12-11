package solver

import "math/rand"

type Solution struct {
	n                int32
	blockedToAllowed map[int]int
}

func Constructor(n int, blacklist []int) Solution {
	divider := n - len(blacklist)
	tempAllowList := make([]int, 0, n-divider)
	for i := divider; i < n; i++ {
		tempAllowList = append(tempAllowList, i)
	}

	for _, v := range blacklist {
		if v >= divider {
			tempAllowList = remove(tempAllowList, v-divider)
		}
	}

	blockedToAllowed := make(map[int]int)
	iterator := 0

	for _, v := range blacklist {
		if v < divider {
			blockedToAllowed[v] = tempAllowList[iterator]
			iterator++
		}
	}

	return Solution{n: int32(divider), blockedToAllowed: blockedToAllowed}
}

func (this *Solution) Pick() int {
	num := rand.Int31n(this.n)
	replacement, found := this.blockedToAllowed[int(num)]

	if found {
		return replacement
	}

	return int(num)
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
