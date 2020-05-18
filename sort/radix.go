package sort

import (
	"math"
)

func radixSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	min := math.MaxInt32
	max := math.MinInt32
	for _, v := range arr {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	base := 10
	previousDivision := 1
	currentDivision := 10
	for {
		lastRound := false
		if max < currentDivision {
			lastRound = true
		}

		counterSpace := make([][]int, base)
		for _, v := range arr {
			sortValue := 0
			if v >= previousDivision {
				sortValue = (v % currentDivision) / previousDivision
			}
			counterSpace[sortValue] = append(counterSpace[sortValue], v)
		}

		arrIndex := 0
		for _, values := range counterSpace {
			for _, v := range values {
				arr[arrIndex] = v
				arrIndex++
			}
		}

		if lastRound {
			break
		}

		previousDivision = currentDivision
		currentDivision = currentDivision * base
	}
}
