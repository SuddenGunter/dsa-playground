package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExamples(main *testing.T) {
	main.Run("1", func(t *testing.T) {
		resultTree := invertTree(fromSlice([]int{4, 2, 7, 1, 3, 6, 9}))
		result := asSlice(resultTree)
		assert.Equal(t, []int{4, 7, 2, 9, 6, 3, 1}, result)
	})
	main.Run("2", func(t *testing.T) {
		resultTree := invertTree(fromSlice([]int{2, 1, 3}))
		result := asSlice(resultTree)
		assert.Equal(t, []int{2, 3, 1}, result)
	})
	main.Run("3", func(t *testing.T) {
		resultTree := invertTree(fromSlice([]int{}))
		result := asSlice(resultTree)
		assert.Equal(t, []int{}, result)
	})
}
