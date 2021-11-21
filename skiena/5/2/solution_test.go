package solution_test

import (
	"testing"

	solution "github.com/SuddenGunter/dsa-playground/skiena/5/2"
	"github.com/stretchr/testify/assert"
)

func Test_FindMissingEven(t *testing.T) {
	src := []int{1, 2, 3, 5}

	missing := solution.FindMissing(src)

	assert.Equal(t, 4, missing)
}

func Test_FindMissingOdd(t *testing.T) {
	src := []int{1, 3, 4, 5, 6}

	missing := solution.FindMissing(src)

	assert.Equal(t, 2, missing)
}

func Test_FindMissingFirstMid(t *testing.T) {
	src := []int{1, 3, 4, 5}

	missing := solution.FindMissing(src)

	assert.Equal(t, 2, missing)
}

// todo: test if missing first mid
