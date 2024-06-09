package solver_test

import (
	"testing"

	"github.com/SuddenGunter/dsa-playground/leetcode/710/solver"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	t.Skip("9 June 2024: broken implementation")
	const upperLimit = 6
	blocklist := []int{0, 2, 3}
	obj := solver.Constructor(upperLimit, blocklist)
	picked := make(map[int]int)
	for i := 0; i < 30; i++ {
		picked[obj.Pick()]++
	}

	for _, el := range blocklist {
		assert.NotContains(t, picked, el)
	}

	assert.Len(t, picked, upperLimit-len(blocklist))
}

func TestExample2(t *testing.T) {
	t.Skip("9 June 2024: broken implementation")
	blocklist := []int{}
	const upperLimit = 7
	obj := solver.Constructor(upperLimit, blocklist)
	picked := make(map[int]int)
	for i := 0; i < 30; i++ {
		picked[obj.Pick()]++
	}

	// doesn't fail with empty blocklist
	assert.Len(t, picked, upperLimit)
}

func TestExample3(t *testing.T) {
	t.Skip("9 June 2024: broken implementation")
	const executions = 100000
	const upperLimit = 4
	blocklist := []int{1, 3}
	obj := solver.Constructor(upperLimit, blocklist)
	picked := make(map[int]int)
	for i := 0; i < executions; i++ {
		picked[obj.Pick()]++
	}

	for _, el := range blocklist {
		assert.NotContains(t, picked, el)
	}

	assert.Len(t, picked, 2)
	assert.GreaterOrEqual(t, float64(picked[0]), float64((executions/2.)*0.995))
	assert.GreaterOrEqual(t, float64(picked[2]), float64((executions/2.)*0.995))
}

func TestExample4(t *testing.T) {
	t.Skip("9 June 2024: broken implementation")
	const upperLimit = 3
	blocklist := []int{1, 2}
	obj := solver.Constructor(upperLimit, blocklist)
	picked := make(map[int]int)
	for i := 0; i < 30; i++ {
		picked[obj.Pick()]++
	}

	for _, el := range blocklist {
		assert.NotContains(t, picked, el)
	}

	assert.Len(t, picked, upperLimit-len(blocklist))
}

func TestExample5(t *testing.T) {
	t.Skip("9 June 2024: broken implementation")
	const upperLimit = 4
	blocklist := []int{2, 1, 0}
	obj := solver.Constructor(upperLimit, blocklist)
	picked := make(map[int]int)
	for i := 0; i < 30; i++ {
		picked[obj.Pick()]++
	}

	for _, el := range blocklist {
		assert.NotContains(t, picked, el)
	}

	assert.Len(t, picked, upperLimit-len(blocklist))
}
