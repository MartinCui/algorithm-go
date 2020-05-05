package symboltable

import (
	common "github.com/martincui/algorithm"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

type testComparable struct {
	v int
}

func (tc testComparable) CompareTo(other interface{}) int {
	typedOther := other.(testComparable)
	return tc.v - typedOther.v
}

func TestBinarySearch(t *testing.T) {
	test(t, NewBinarySearchSt())
}

func TestBinarySearchTree(t *testing.T) {
	test(t, NewBinarySearchTreeSt())
}

func TestRedBlackTree(t *testing.T) {
	test(t, NewRedBlackTreeSt())
}

func test(t *testing.T, st SymbolTable) {
	nativeSt := SymbolTable(&nativeMapSt{m: make(map[Comparable]interface{})})
	rq := require.New(t)

	for i := 0; i < 100; i++ {
		size := 999
		intKeys := common.MakeRandomSlice(size, true)
		values := common.MakeRandomSlice(size, true)
		keys := make([]testComparable, size)
		for i, ik := range intKeys {
			keys[i] = testComparable{v: ik}
		}

		for i, key := range keys {
			st.Set(key, values[i])
			nativeSt.Set(key, values[i])
		}

		for _, key := range keys {
			rq.Equal(nativeSt.Get(key), st.Get(key))
		}
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	benchmark(b, NewBinarySearchSt())
}

func BenchmarkBinarySearchTree(b *testing.B) {
	benchmark(b, NewBinarySearchTreeSt())
}

func BenchmarkRedBlackTree(b *testing.B) {
	benchmark(b, NewRedBlackTreeSt())
}

func BenchmarkNativeMap(b *testing.B) {
	benchmark(b, &nativeMapSt{m: make(map[Comparable]interface{}, 0)})
}

type nativeMapSt struct {
	m map[Comparable]interface{}
}

func (nm *nativeMapSt) Get(key Comparable) interface{} {
	return nm.m[key]
}

func (nm *nativeMapSt) Set(key Comparable, value interface{}) {
	nm.m[key] = value
}

func (nm *nativeMapSt) Size() int { return len(nm.m) }

func benchmark(b *testing.B, st SymbolTable) {
	size := 9999
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			st.Set(testComparable{v: rd.Int()}, rd.Int())
			_ = st.Get(testComparable{v: rd.Int()})
		}
	}
}
