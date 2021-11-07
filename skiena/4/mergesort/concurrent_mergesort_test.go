package mergesort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcurrentSort_OnTwoElements(t *testing.T) {
	input := []float64{
		2, 1,
	}
	expected := []float64{
		1, 2,
	}

	result := ConcurrentSort(input)

	assert.Equal(t, expected, result)
}

func TestConcurrentSort_OnUnsortedArray(t *testing.T) {
	input := []float64{
		6, 1, 3, 8, 7, 9, 1, 9,
	}
	expected := []float64{
		1, 1, 3, 6, 7, 8, 9, 9,
	}

	result := ConcurrentSort(input)

	assert.Equal(t, expected, result)
}

func TestConcurrentSort_OnSortedArray(t *testing.T) {
	input := []float64{
		1, 1, 3, 6, 7, 8, 9, 9,
	}
	expected := []float64{
		1, 1, 3, 6, 7, 8, 9, 9,
	}

	result := ConcurrentSort(input)

	assert.Equal(t, expected, result)
}

func TestConcurrentSort_OnEmptyArray(t *testing.T) {
	input := []float64{}
	expected := []float64{}

	result := ConcurrentSort(input)

	assert.Equal(t, expected, result)

}
