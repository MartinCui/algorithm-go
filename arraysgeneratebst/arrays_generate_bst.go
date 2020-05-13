package arraysgeneratebst

type node struct {
	value int
	left  *node
	right *node
}

func arraysThatGenerateBst(root *node) [][]int {
	if root == nil {
		return nil
	}

	if root.left == nil && root.right == nil {
		return [][]int{{root.value}}
	}

	leftArrays := arraysThatGenerateBst(root.left)
	rightArrays := arraysThatGenerateBst(root.right)
	allChildrenArrays := sequentialMergeArrays(leftArrays, rightArrays)
	var result [][]int
	for _, childArray := range allChildrenArrays {
		result = append(result, append([]int{root.value}, childArray...))
	}
	return result
}

func sequentialMergeArrays(left, right [][]int) [][]int {
	if len(left) == 0 {
		return right
	}

	if len(right) == 0 {
		return left
	}

	var result [][]int
	for _, leftArray := range left {
		for _, rightArray := range right {
			result = append(result, sequentialMergeArray(leftArray, rightArray)...)
		}
	}

	return result
}

func sequentialMergeArray(left, right []int) [][]int {
	totalLength := len(left) + len(right)
	leftPositions := calculateSetPositions(left, 0, totalLength)
	var result [][]int
	for _, leftPosition := range leftPositions {
		arrangement := make([]int, totalLength)
		currentRightIndex := 0
		currentLeftIndex := 0
		for i := 0; i < totalLength; i++ {
			if currentLeftIndex < len(left) && i == leftPosition[currentLeftIndex] {
				arrangement[i] = left[currentLeftIndex]
				currentLeftIndex++
			} else {
				arrangement[i] = right[currentRightIndex]
				currentRightIndex++
			}
		}
		result = append(result, arrangement)
	}

	return result
}

func calculateSetPositions(arr []int, fromPosition, totalLength int) (setPositions [][]int) {
	var result [][]int
	if len(arr) == 1 {
		for position := fromPosition; position < totalLength; position++ {
			result = append(result, []int{position})
		}
		return result
	}

	for position := fromPosition; position < totalLength-len(arr)+1; position++ {
		childPositions := calculateSetPositions(arr[1:], position+1, totalLength)
		for _, childPosition := range childPositions {
			result = append(result, append([]int{position}, childPosition...))
		}
	}
	return result
}
