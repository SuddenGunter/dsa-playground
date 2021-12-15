package solution_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	solution "github.com/SuddenGunter/dsa-playground/leetcode/278/simpleBinarySearch"
)

func TestExample1(t *testing.T) {
	solution.IsBadVersion = func(version int) bool {
		if version > 5 || version <= 0 {
			panic("out of bounds")
		}
		return version >= 4
	}

	result := solution.Solve(5)

	assert.Equal(t, 4, result)
}

func TestExample2(t *testing.T) {
	solution.IsBadVersion = func(version int) bool {
		if version > 5 || version <= 0 {
			panic("out of bounds")
		}
		return version == 5
	}

	result := solution.Solve(5)

	assert.Equal(t, 5, result)
}

func TestExample3(t *testing.T) {
	solution.IsBadVersion = func(version int) bool {
		if version > 5 || version <= 0 {
			panic("out of bounds")
		}

		return version >= 1
	}

	result := solution.Solve(5)

	assert.Equal(t, 1, result)
}

func TestExample4(t *testing.T) {
	solution.IsBadVersion = func(version int) bool {
		if version > 3 || version <= 0 {
			panic("out of bounds")
		}

		return version >= 2
	}

	result := solution.Solve(3)

	assert.Equal(t, 2, result)
}

func TestExample6(t *testing.T) {
	solution.IsBadVersion = func(version int) bool {
		if version > 3 || version <= 0 {
			panic("out of bounds")
		}

		return version >= 1
	}

	result := solution.Solve(4)

	assert.Equal(t, 1, result)
}

func TestExample5(t *testing.T) {
	solution.IsBadVersion = func(version int) bool {
		if version > 3 || version <= 0 {
			panic("out of bounds")
		}

		return version >= 1
	}

	result := solution.Solve(3)

	assert.Equal(t, 1, result)
}
