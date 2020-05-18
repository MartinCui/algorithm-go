package sort

import "math"

func counterSort(arr []int) {
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

	if max-min > 10240{
		panic("counter sort is not designed for sorting values with big diff")
	}

	counterSpace := make([]int, max-min+1)
	for _, v := range arr {
		counterSpace[v-min]++
	}

	arrIndex := 0
	for i, v := range counterSpace {
		if v == 0 {
			continue
		}

		valueToAdd := min + i
		for j := 0; j < v; j++ {
			arr[arrIndex] = valueToAdd
			arrIndex++
		}
	}
}
