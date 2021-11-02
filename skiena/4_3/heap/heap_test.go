package heap

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func isValidHeap(heap *Heap) bool {
	for i, v := range heap.body {
		if heap.parentIndexOf(i) != -1 {
			if heap.comparator(heap.body[heap.parentIndexOf(i)], v) == 1 {
				return false
			}
		}
	}

	return true
}

func Test_FromSlice_ArrangesDataCorrectly(t *testing.T) {
	src := []float64{1492, 1941, 2001, 1918, 1963, 1865}
	h := FromSlice(src)

	assert.True(t, isValidHeap(h))
}

func Test_TakeTop_ArrangesDataCorrectly(t *testing.T) {
	src := []float64{1492, 1941, 2001, 1918, 1963, 1865}
	h := FromSlice(src)
	require.True(t, isValidHeap(h))

	top, err := h.TakeTop()
	assert.NoError(t, err)
	assert.True(t, isValidHeap(h))
	assert.Equal(t, 1492., top)

	top, err = h.TakeTop()
	assert.NoError(t, err)
	assert.True(t, isValidHeap(h))
	assert.Equal(t, 1865., top)
}

func Test_Heapsort_SortsCorrectly(t *testing.T) {
	src := []float64{1492, 1941, 2001, 1918, 1963, 1865}
	expected := []float64{1492, 1865, 1918, 1941, 1963, 2001}

	result := Heapsort(src)

	assert.ElementsMatch(t, expected, result)
}
