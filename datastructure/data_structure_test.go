package datastructure

import (
	common "github.com/martincui/algorithm"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestStack(t *testing.T) {
	rq := require.New(t)

	s := NewStack()
	rq.True(s.IsEmpty())

	valuesToPush := []int{2, 3, 56, 7, 5, 23, 345, 677, 2}
	for _, v := range valuesToPush {
		s.Push(v)
	}

	for i := len(valuesToPush) - 1; i >= 0; i-- {
		peekValue, err := s.Peek()
		rq.NoError(err)
		rq.Equal(valuesToPush[i], peekValue.(int))

		popValue, err := s.Pop()
		rq.NoError(err)
		rq.Equal(valuesToPush[i], popValue.(int))
	}

	_, err := s.Peek()
	rq.Error(err)
	_, err = s.Pop()
	rq.Error(err)
	rq.True(s.IsEmpty())
}

func TestQueue(t *testing.T) {
	rq := require.New(t)

	q := NewQueue()
	rq.True(q.IsEmpty())

	valuesToAdd := []int{2, 3, 56, 7, 5, 23, 345, 677, 2}
	for _, v := range valuesToAdd {
		q.Add(v)
	}

	for _, v := range valuesToAdd {
		peekValue, err := q.Peek()
		rq.NoError(err)
		rq.Equal(v, peekValue.(int))

		popValue, err := q.Pop()
		rq.NoError(err)
		rq.Equal(v, popValue.(int))
	}

	_, err := q.Peek()
	rq.Error(err)
	_, err = q.Pop()
	rq.Error(err)
	rq.True(q.IsEmpty())
}

func TestEvaluateExpression(t *testing.T) {
	rq := require.New(t)

	requireExpressionResult(rq, "3 + 5 - 7", 1)
	requireExpressionResult(rq, "3 * 5 - (7 + 3) / 5", 9)
	requireExpressionResult(rq, "((3 + 5) - 7) / 5", 0.2)
	requireExpressionResult(rq, "(3 + (((5 - 7) * 3) / 5))", 1.8)
	requireExpressionResult(rq, "(3.3 * 2.5) /1.1", 7.5)
}

func requireExpressionResult(rq *require.Assertions, expression string, expectedResult float64) {
	result, err := EvaluateExpression(expression)
	rq.NoError(err)
	rq.LessOrEqual(math.Abs(expectedResult-result), 0.00000001)
}

func TestBinaryHeap(t *testing.T) {
	rq := require.New(t)
	heap := newHeapPriorityQueue()
	randomSlice := common.MakeRandomSlice(10, true)
	for _, v := range randomSlice {
		heap.add(v)
	}
	randomSlice2 := common.MakeRandomSlice(10, true)
	for i := 0; i < len(randomSlice)/2; i++ {
		heap.popMax()
		heap.add(randomSlice2[i])
		heap.add(randomSlice2[i] + len(randomSlice)/2)
	}

	previousMax := math.MaxInt32
	for count := 0; count < int(float32(len(randomSlice))*1.5); count++ {
		currentMax := heap.popMax()
		rq.GreaterOrEqual(previousMax, currentMax)
		previousMax = currentMax
	}

	rq.Equal(0, heap.size)
}

func TestTrie(t *testing.T) {
	rq := require.New(t)
	trie := NewTrie(5)
	trie.Set([]int{2, 3, 4}, "234")
	trie.Set([]int{0, 1}, "01")
	trie.Set([]int{0, 1, 2}, "012")
	trie.Set([]int{0, 1, 2, 4}, "0124")
	trie.Set([]int{2, 4}, "24")
	rq.Equal("234", trie.Get([]int{2, 3, 4}))
	rq.Equal("01", trie.Get([]int{0, 1}))
	rq.Equal("012", trie.Get([]int{0, 1, 2}))
	rq.Equal("0124", trie.Get([]int{0, 1, 2, 4}))
	rq.Equal("24", trie.Get([]int{2, 4}))
	rq.Nil(trie.Get([]int{0}))
	rq.Nil(trie.Get([]int{3, 4}))
}
