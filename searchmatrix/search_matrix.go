package searchmatrix

import "fmt"

/*
matrix is ordered from left to right, and also top to bottom like below:
1,  2,  5,  7,  9
3,  6,  17, 25, 25
4,  7,  18, 26, 28

the task is to find value "v" position in this matrix (anyone is okay if there are more than 1)
*/

func findInMatrix(m [][]int, v int) *findResult {
	if len(m) == 0 || len(m[0]) == 0 {
		return &findResult{found: false}
	}

	return _findInMatrix(m, v, &position{0, 0}, &position{len(m) - 1, len(m[0]) - 1})
}

func _findInMatrix(m [][]int, v int, from *position, to *position) *findResult {
	if from.rowIndex > len(m)-1 || from.columnIndex > len(m[0])-1 {
		return &findResult{found: false}
	}

	if m[from.rowIndex][from.columnIndex] == v {
		return &findResult{thePosition: from, found: true}
	}

	if from.Equals(to) {
		return &findResult{found: false}
	}

	diagonalStep := minInt(to.columnIndex-from.columnIndex, to.rowIndex-from.rowIndex)
	diagonalEnd := &position{
		rowIndex:    from.rowIndex + diagonalStep,
		columnIndex: from.columnIndex + diagonalStep,
	}

	firstBiggerThanVDiagonalPosition, match := findFirstBiggerThanVDiagonalPosition(m, v, from, diagonalEnd)
	if match {
		return &findResult{thePosition: firstBiggerThanVDiagonalPosition, found: true}
	}
	return findInPartitions(m, v, from, to, firstBiggerThanVDiagonalPosition)
}

func findInPartitions(m [][]int, v int, from *position, to *position, firstBiggerThanVDiagonalPosition *position) *findResult {
	rightTopFrom := &position{
		rowIndex:    from.rowIndex,
		columnIndex: firstBiggerThanVDiagonalPosition.columnIndex,
	}
	rightTopTo := &position{
		rowIndex:    firstBiggerThanVDiagonalPosition.rowIndex - 1,
		columnIndex: to.columnIndex,
	}
	leftBottomFrom := &position{
		rowIndex:    firstBiggerThanVDiagonalPosition.rowIndex,
		columnIndex: from.columnIndex,
	}
	leftBottomTo := &position{
		rowIndex:    to.rowIndex,
		columnIndex: firstBiggerThanVDiagonalPosition.columnIndex - 1,
	}

	if firstBiggerThanVDiagonalPosition.rowIndex > to.rowIndex {
		if firstBiggerThanVDiagonalPosition.columnIndex > to.columnIndex {
			return &findResult{found: false}
		} else {
			return _findInMatrix(m, v, rightTopFrom, to)
		}
	} else {
		if firstBiggerThanVDiagonalPosition.columnIndex > to.columnIndex {
			return _findInMatrix(m, v, leftBottomFrom, to)
		} else {
			rightTopSearchResult := _findInMatrix(m, v, rightTopFrom, rightTopTo)
			if rightTopSearchResult.found {
				return rightTopSearchResult
			} else {
				return _findInMatrix(m, v, leftBottomFrom, leftBottomTo)
			}
		}
	}
}

func findFirstBiggerThanVDiagonalPosition(m [][]int, v int, from *position, diagonalEnd *position) (p *position, valueMatch bool) {
	for from.before(diagonalEnd) {
		mid := from.mid(diagonalEnd)
		if m[mid.rowIndex][mid.columnIndex] > v {
			diagonalEnd = mid.diagonalPrevious()
		} else if m[mid.rowIndex][mid.columnIndex] < v {
			from = mid.diagonalNext()
		} else {
			return mid, true
		}
	}

	if m[from.rowIndex][from.columnIndex] > v {
		return from, false
	} else if m[from.rowIndex][from.columnIndex] < v {
		return from.diagonalNext(), false
	} else {
		return from, true
	}
}

func minInt(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

type position struct {
	rowIndex    int
	columnIndex int
}

func (p *position) Equals(another *position) bool {
	return p.rowIndex == another.rowIndex && p.columnIndex == another.columnIndex
}

func (p *position) String() string {
	return fmt.Sprintf("row: %d, column: %d", p.rowIndex, p.columnIndex)
}

func (p *position) before(another *position) bool {
	return p.rowIndex < another.rowIndex && p.columnIndex < another.columnIndex
}

func (p *position) mid(biggerPosition *position) *position {
	return &position{
		rowIndex:    p.rowIndex + (biggerPosition.rowIndex-p.rowIndex)/2,
		columnIndex: p.columnIndex + (biggerPosition.columnIndex-p.columnIndex)/2,
	}
}

func (p *position) diagonalPrevious() *position {
	return &position{
		rowIndex:    p.rowIndex - 1,
		columnIndex: p.columnIndex - 1,
	}
}

func (p *position) diagonalNext() *position {
	return &position{
		rowIndex:    p.rowIndex + 1,
		columnIndex: p.columnIndex + 1,
	}
}

type findResult struct {
	thePosition *position
	found       bool
}

func (r *findResult) String() string {
	return fmt.Sprintf("found: %v, position: %s", r.found, r.thePosition)
}
