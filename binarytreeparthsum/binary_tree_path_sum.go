package binarytreeparthsum

type node struct {
	value int
	left  *node
	right *node
}

func countPathsWithSum(root *node, targetValue int) int {
	totalSumToCountDic := make(map[int]int)
	return _countPathsWithSum(root, 0, targetValue, totalSumToCountDic)
}

func _countPathsWithSum(n *node, parentTotalSum, targetValue int, totalSumToCountDic map[int]int) int {
	if n == nil {
		return 0
	}

	thisTotalSum := n.value + parentTotalSum
	pathCountFromThisNodeUp := 0
	if thisTotalSum == targetValue {
		pathCountFromThisNodeUp++
	}

	matchingPathCountFromThisUp, exists := totalSumToCountDic[thisTotalSum-targetValue]
	if exists {
		pathCountFromThisNodeUp += matchingPathCountFromThisUp
	}

	addTotalSumToCountDic(totalSumToCountDic, thisTotalSum, 1)
	pathCountFromLeftDescendantUp := _countPathsWithSum(n.left, thisTotalSum, targetValue, totalSumToCountDic)
	pathCountFromRightDescendantUp := _countPathsWithSum(n.right, thisTotalSum, targetValue, totalSumToCountDic)
	addTotalSumToCountDic(totalSumToCountDic, thisTotalSum, -1)

	return pathCountFromThisNodeUp + pathCountFromLeftDescendantUp + pathCountFromRightDescendantUp
}

func addTotalSumToCountDic(totalSumToCountDic map[int]int, key, delta int) {
	currentValue, exists := totalSumToCountDic[key]
	if exists {
		totalSumToCountDic[key] = currentValue + delta
	} else {
		totalSumToCountDic[key] = delta
	}
}
