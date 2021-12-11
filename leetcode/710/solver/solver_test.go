package solver_test

import (
	"testing"

	"github.com/SuddenGunter/dsa-playground/leetcode/710/solver"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	blocklist := []int{0, 2, 3}
	obj := solver.Constructor(6, blocklist)
	picked := make(map[int]int)
	for i := 0; i < 30; i++ {
		picked[obj.Pick()]++
	}

	for _, el := range blocklist {
		assert.NotContains(t, picked, el)
	}
}

func TestExample2(t *testing.T) {
	blocklist := []int{}
	obj := solver.Constructor(7, blocklist)
	picked := make(map[int]int)
	for i := 0; i < 30; i++ {
		picked[obj.Pick()]++
	}

	// doesn't fail with empty blocklist
	assert.Len(t, picked, 30)
}

func TestExample3(t *testing.T) {
	const executions = 10000

	blocklist := []int{2, 3}
	obj := solver.Constructor(4, blocklist)
	picked := make(map[int]int)
	for i := 0; i < executions; i++ {
		picked[obj.Pick()]++
	}

	assert.Len(t, picked, 2)
	assert.GreaterOrEqual(t, picked[0], (executions/2)*0.45)
	assert.GreaterOrEqual(t, picked[1], (executions/2)*0.45)
}
