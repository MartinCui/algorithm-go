package search

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestInterpolationSearch(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	key := r.Intn(999)
	arr := generateSortedArray(99999, key)
	correctAnswer := sort.SearchInts(arr, key)
	interpolationAnswer := interpolationSearch(arr, key)
	require.Equal(t, arr[correctAnswer], arr[interpolationAnswer])
}

func BenchmarkInterpolationSearch(b *testing.B) {
	benchmarkSearch(b, interpolationSearch)
}

func BenchmarkGolangSearch(b *testing.B) {
	benchmarkSearch(b, sort.SearchInts)
}

func benchmarkSearch(b *testing.B, searchFunc func (a []int, x int) int ) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	key := r.Intn(9999999)
	arr := generateSortedArray(9999999, key)

	b.ResetTimer()
	for i:=0; i<b.N; i++{
		searchFunc(arr, key)
	}
}

func generateSortedArray(size, includeValue int) []int{
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rawArray := make([]int, size)
	for i:=0; i < size-1; i++{
		rawArray[i] = r.Intn(99999)
	}
	rawArray[size-1] = includeValue

	sort.Ints(rawArray)
	return rawArray
}