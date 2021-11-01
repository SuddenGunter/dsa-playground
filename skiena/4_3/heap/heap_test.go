package heap_test

import (
	"fmt"
	"testing"

	"github.com/SuddenGunter/dsa-playground/skiena/4_3/heap"
)

func Test_FromSlice_ArrangesDataCorrectly(t *testing.T) {
	src := []rune{'c', 'd', 'a', 'b', 'e', 'f'}
	h := heap.FromSlice(src)

	fmt.Printf("%v", h)
}
