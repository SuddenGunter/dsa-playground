package solution_test

import (
	"testing"

	solution "github.com/SuddenGunter/dsa-playground/skiena/5/1_2"
	"github.com/stretchr/testify/assert"
)

func Test_FindMaxShiftedTwoTimes(t *testing.T) {
	src := []int{35, 42, 5, 15, 27, 29}

	max := solution.FindMax(src)

	assert.Equal(t, 5, max)
}

func Test_FindMaxShiftedFourTimes(t *testing.T) {
	src := []int{27, 29, 35, 42, 5, 15}

	max := solution.FindMax(src)

	assert.Equal(t, 5, max)
}

func Test_FindMaxNotShifted(t *testing.T) {
	src := []int{5, 15, 27, 29, 35, 42}

	max := solution.FindMax(src)

	assert.Equal(t, 5, max)
}
