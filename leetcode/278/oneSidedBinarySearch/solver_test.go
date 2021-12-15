package solution_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	solution "github.com/SuddenGunter/dsa-playground/leetcode/278/oneSidedBinarySearch"
)

func TestExample1(t *testing.T) {
	solution.IsBadVersion = func(version int) bool {
		return version >= 4
	}

	result := solution.Solve(5)

	assert.Equal(t, 4, result)
}

func TestExample2(t *testing.T) {
	solution.IsBadVersion = func(version int) bool {
		return version >= 4
	}

	result := solution.Solve(5)

	assert.Equal(t, 4, result)
}

func TestExample3(t *testing.T) {
	solution.IsBadVersion = func(version int) bool {
		return version >= 4
	}

	result := solution.Solve(5)

	assert.Equal(t, 4, result)
}

func TestExample4(t *testing.T) {
	solution.IsBadVersion = func(version int) bool {
		return version >= 4
	}

	result := solution.Solve(5)

	assert.Equal(t, 4, result)
}
