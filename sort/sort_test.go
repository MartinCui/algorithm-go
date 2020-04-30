package sort

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"sort"
	"testing"
	"time"
)

const (
	testArraySize = 999
)

func TestSelection(t *testing.T) {
	testSort(t, selectionSort)
}

func TestPopUp(t *testing.T) {
	testSort(t, popUpSort)
}

func TestInsertion(t *testing.T) {
	testSort(t, insertionSort)
}

func TestBinarySearchInsertion(t *testing.T) {
	testSort(t, binarySearchInsertionSort)
}

func makeRandomSlice(size int) []int {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, size)
	for i := 0; i < len(arr); i++ {
		arr[i] = rnd.Int()
	}
	return arr
}

func testSort(t *testing.T, sortFun func([]int) []int) {
	rq := require.New(t)
	for i := 0; i < 1000; i++ {
		arr := makeRandomSlice(testArraySize)

		correctResult := make([]int, len(arr))
		copy(correctResult, arr)

		result := sortFun(arr)
		sort.Ints(correctResult)
		rq.Equal(correctResult, result)
	}
}

func benchmarkSort(b *testing.B, sortFun func([]int) []int) {
	arr := makeRandomSlice(9999)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sortFun(arr)
	}
}

func BenchmarkSelection(b *testing.B) {
	benchmarkSort(b, selectionSort)
}

func BenchmarkPopUp(b *testing.B) {
	benchmarkSort(b, popUpSort)
}

func BenchmarkInsertion(b *testing.B) {
	benchmarkSort(b, insertionSort)
}

func BenchmarkBinarySearchInsertion(b *testing.B) {
	benchmarkSort(b, binarySearchInsertionSort)
}
