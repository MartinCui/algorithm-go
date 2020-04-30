package sort

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minPosition := i
		min := arr[i]
		for j := i; j < len(arr); j++ {
			if arr[j] < min {
				minPosition = j
				min = arr[j]
			}
		}
		arr[minPosition] = arr[i]
		arr[i] = min
	}

	return arr
}
