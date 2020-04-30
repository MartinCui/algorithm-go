package datastructure

import (
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