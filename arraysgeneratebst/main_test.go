package arraysgeneratebst

import "testing"

func TestArraysBst(t *testing.T) {
	node213 := &node{
		value: 2,
		left:  &node{value: 1},
		right: &node{value: 3},
	}
	result := arraysThatGenerateBst(node213)
	t.Log(result)

	node526 := &node{
		value: 5,
		left:  node213,
		right: &node{value: 6},
	}
	result = arraysThatGenerateBst(node526)
	t.Log(result)
}
