package sort

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"sort"
	"testing"
	"time"
)

const (
	testArraySize  = 999
	benchArraySize = 9999
	testStoreTimes = 100
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

func TestShell(t *testing.T){
	testSort(t, shellSort)
}

func TestMerge(t *testing.T){
	testSort(t, mergeSort)
}

func TestQuick(t *testing.T){
	testSort(t, quickSort)
}

func makeRandomSlice(size int) []int {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, size)
	for i := 0; i < len(arr); i++ {
		arr[i] = rnd.Int()
	}
	return arr
}

func testSort(t *testing.T, sortFun func([]int)) {
	rq := require.New(t)
	for i := 0; i < testStoreTimes; i++ {
		arr := makeRandomSlice(testArraySize)

		correctResult := make([]int, len(arr))
		copy(correctResult, arr)

		sortFun(arr)
		sort.Ints(correctResult)
		rq.Equal(correctResult, arr)
	}
}

func benchmarkSort(b *testing.B, sortFun func([]int)) {
	randomArrays := make([][]int, b.N)
	for i := 0; i < len(randomArrays); i++ {
		randomArrays[i] = makeRandomSlice(benchArraySize)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortFun(randomArrays[i])
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

func BenchmarkShell(b *testing.B) {
	benchmarkSort(b, shellSort)
}

func BenchmarkMerge(b *testing.B) {
	benchmarkSort(b, mergeSort)
}

func BenchmarkQuick(b *testing.B) {
	benchmarkSort(b, quickSort)
}

