package search

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestInterpolationSearch(t *testing.T) {
	testSearch(t, 100, interpolationSearch)
}

func TestBinarySearch(t *testing.T) {
	testSearch(t, 100, binarySearch)
}

func testSearch(t *testing.T, n int, searchFunc func(a []int, x int) int) {
	for i := 0; i < n; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		key := r.Int()
		arr := generateSortedArray(10000, key)
		correctAnswer := sort.SearchInts(arr, key)
		searchAnswer := searchFunc(arr, key)
		require.Equal(t, arr[correctAnswer], arr[searchAnswer])
	}
}

func BenchmarkInterpolationSearch(b *testing.B) {
	benchmarkSearch(b, interpolationSearch)
}

func BenchmarkGolangSearch(b *testing.B) {
	benchmarkSearch(b, sort.SearchInts)
}

func BenchmarkBinarySearch(b *testing.B) {
	benchmarkSearch(b, binarySearch)
}

func benchmarkSearch(b *testing.B, searchFunc func(a []int, x int) int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	key := r.Intn(9999999)
	arr := generateSortedArray(9999999, key)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		searchFunc(arr, key)
	}
}

func generateSortedArray(size, includeValue int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rawArray := make([]int, size)
	for i := 0; i < size-1; i++ {
		rawArray[i] = r.Int()
	}
	rawArray[size-1] = includeValue

	sort.Ints(rawArray)
	return rawArray
}
