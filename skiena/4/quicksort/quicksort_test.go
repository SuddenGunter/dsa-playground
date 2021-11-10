package quicksort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort_OnUnsortedArray(t *testing.T) {
	input := []float64{
		6, 8, 2, 4, 1, 0, 9, 5, 3, 7,
	}
	expected := []float64{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	}

	result := Sort(input)

	assert.Equal(t, expected, result)
}

func TestSort_OnSortedArray(t *testing.T) {
	input := []float64{
		1, 1, 3, 6, 7, 8, 9, 9,
	}
	expected := []float64{
		1, 1, 3, 6, 7, 8, 9, 9,
	}

	result := Sort(input)

	assert.Equal(t, expected, result)
}

func TestSort_OnEmptyArray(t *testing.T) {
	input := []float64{}
	expected := []float64{}

	result := Sort(input)

	assert.Equal(t, expected, result)
}

func TestSort_OnTwoElements(t *testing.T) {
	input := []float64{
		2, 1,
	}
	expected := []float64{
		1, 2,
	}

	result := Sort(input)

	assert.Equal(t, expected, result)
}
