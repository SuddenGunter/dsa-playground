package solver_test

import (
	"testing"

	"github.com/SuddenGunter/dsa-playground/leetcode/20/solver"
	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	s := "()"

	res := solver.IsValid(s)

	assert.True(t, res)
}

func TestExample2(t *testing.T) {
	s := "()[]{}"

	res := solver.IsValid(s)

	assert.True(t, res)
}

func TestExample3(t *testing.T) {
	s := "(]"

	res := solver.IsValid(s)

	assert.False(t, res)
}

func TestExample4(t *testing.T) {
	s := "([)]"

	res := solver.IsValid(s)

	assert.False(t, res)
}

func TestExample5(t *testing.T) {
	s := "{[]}"

	res := solver.IsValid(s)

	assert.True(t, res)
}

func TestCase1(t *testing.T) {
	s := "["

	res := solver.IsValid(s)

	assert.False(t, res)
}

func TestCase2(t *testing.T) {
	s := "(("

	res := solver.IsValid(s)

	assert.False(t, res)
}
