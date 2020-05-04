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

func test(t *testing.T, st SymbolTable) {
	rq := require.New(t)

	for i := 0; i < 100; i++ {
		size := 999
		intKeys := common.MakeRandomSlice(size, true)
		values := common.MakeRandomSlice(size, true)
		keys := make([]testComparable, size)
		for i, ik := range intKeys {
			keys[i] = testComparable{v: ik}
		}

		correctResult := make(map[testComparable]int)
		for i, key := range keys {
			st.Set(key, values[i])
			correctResult[key] = values[i]
		}

		for _, key := range keys {
			rq.Equal(correctResult[key], st.Get(key))
		}
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	benchmark(b, NewBinarySearchSt())
}

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
