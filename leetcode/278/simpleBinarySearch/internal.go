package solution

var IsBadVersion func(version int) bool

func isBadVersion(version int) bool {
	return IsBadVersion(version)
}

func Solve(n int) int {
	return firstBadVersion(n)
}
