package solver

import "math/rand"

type Solution struct {
	n                int32
	blockedToAllowed map[int32]int32
}

func Constructor(n int, blacklist []int) Solution {
	divider := n - len(blacklist)
	tempAllowList := make([]int32, 0, n-divider)
	for i := divider; i < n; i++ {
		tempAllowList = append(tempAllowList, int32(i))
	}

	for _, v := range blacklist {
		if v >= divider {
			tempAllowList = remove(tempAllowList, find(tempAllowList, int32(v)))
		}
	}

	blockedToAllowed := make(map[int32]int32)
	iterator := int32(0)

	for _, v := range blacklist {
		if v < divider {
			blockedToAllowed[int32(v)] = tempAllowList[iterator]
			iterator++
		}
	}

	return Solution{n: int32(divider), blockedToAllowed: blockedToAllowed}
}

func (this *Solution) Pick() int {
	num := rand.Int31n(this.n)
	replacement, found := this.blockedToAllowed[num]

	if found {
		return int(replacement)
	}

	return int(num)
}

func find(s []int32, v int32) int32 {
	for i := range s {
		if s[i] == v {
			return int32(i)
		}
	}

	return -1
}

func remove(s []int32, i int32) []int32 {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
