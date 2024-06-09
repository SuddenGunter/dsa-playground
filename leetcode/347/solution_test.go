package solution

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExamples(main *testing.T) {
	main.Run("1", func(t *testing.T) {
		input := []int{1, 1, 1, 2, 2, 3}
		k := 2

		result := topKFrequent(input, k)

		// result is allowed to be in any order, but it's easier to assert it in tests as sorted
		sort.Ints(result)

		assert.Equal(t, []int{1, 2}, result)
	})
	main.Run("2", func(t *testing.T) {
		input := []int{1}
		k := 1

		result := topKFrequent(input, k)

		// result is allowed to be in any order, but it's easier to assert it in tests as sorted
		sort.Ints(result)

		assert.Equal(t, []int{1}, result)
	})
}
