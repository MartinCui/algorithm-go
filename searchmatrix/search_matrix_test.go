package searchmatrix

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	rq := require.New(t)
	m := [][]int{
		{1, 2, 5, 7, 9},
		{3, 6, 17, 25, 25},
		{4, 7, 18, 26, 28},
	}
	rq.Equal(&findResult{thePosition: &position{rowIndex: 0, columnIndex: 0}, found: true}, findInMatrix(m, 1))
	rq.Equal(&findResult{thePosition: &position{rowIndex: 0, columnIndex: 1}, found: true}, findInMatrix(m, 2))
	rq.Equal(&findResult{thePosition: &position{rowIndex: 2, columnIndex: 3}, found: true}, findInMatrix(m, 26))
	rq.Equal(&findResult{thePosition: &position{rowIndex: 1, columnIndex: 2}, found: true}, findInMatrix(m, 17))
	rq.Equal(&findResult{thePosition: &position{rowIndex: 2, columnIndex: 2}, found: true}, findInMatrix(m, 18))
	rq.Equal(&findResult{thePosition: &position{rowIndex: 2, columnIndex: 4}, found: true}, findInMatrix(m, 28))
	rq.Equal(&findResult{thePosition: &position{rowIndex: 0, columnIndex: 3}, found: true}, findInMatrix(m, 7))
	rq.Equal(&findResult{thePosition: &position{rowIndex: 1, columnIndex: 4}, found: true}, findInMatrix(m, 25))
	rq.Equal(&findResult{found: false}, findInMatrix(m, 30))
	rq.Equal(&findResult{found: false}, findInMatrix(m, 8))
}
