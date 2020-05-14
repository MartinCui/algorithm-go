package binarytreeparthsum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPathSum(t *testing.T) {
	rq := require.New(t)
	n := &node{
		value: 10,
		left: &node{
			value: 5,
			left: &node{
				value: 3,
				left:  &node{value: 3},
				right: &node{value: -2},
			},
			right: &node{
				value: 2,
				right: &node{value: 1},
			},
		},
		right: &node{
			value: -3,
			right: &node{value: 11},
		},
	}
	rq.Equal(2, countPathsWithSum(n, 6))
	rq.Equal(2, countPathsWithSum(n, 7))
	rq.Equal(3, countPathsWithSum(n, 8))
}
