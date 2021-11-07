package mergesort

import (
	"math/rand"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Benchmark_Sort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		rand.Seed(time.Now().Unix())
		ints := rand.Perm(2560000)
		floats := make([]float64, 0, cap(ints))
		for i := range ints {
			floats = append(floats, float64(ints[i]))
		}
		ints = nil
		runtime.GC()
		b.StartTimer()

		result := Sort(floats)
		b.StopTimer()
		require.NotNil(b, result)
	}
}

func Benchmark_ConcurrentSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		rand.Seed(time.Now().Unix())
		ints := rand.Perm(2560000)
		floats := make([]float64, 0, cap(ints))
		for i := range ints {
			floats = append(floats, float64(ints[i]))
		}
		ints = nil
		runtime.GC()
		b.StartTimer()

		result := ConcurrentSort(floats)
		b.StopTimer()
		require.NotNil(b, result)
	}
}
