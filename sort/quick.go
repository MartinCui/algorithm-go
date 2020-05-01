package sort

func quickSort(arr []int) {
	_quickSortSort(arr, 0, len(arr)-1)
}

func _quickSortSort(arr []int, low, high int) {
	if low >= high {
		return
	}

	middle := _quickSortPartition(arr, low, high)
	_quickSortSort(arr, low, middle-1)
	_quickSortSort(arr, middle+1, high)
}

func _quickSortPartition(arr []int, low, high int) int {
	key := arr[low]
	leftIndex := low
	rightIndex := high + 1
	for ; ; {
		for leftIndex++; leftIndex < high && arr[leftIndex] < key; leftIndex++ {
		}
		for rightIndex--; rightIndex > low && arr[rightIndex] > key; rightIndex-- {
		}
		if leftIndex >= rightIndex{
			break
		}
		exchange(arr, leftIndex, rightIndex)
	}

	exchange(arr, rightIndex, low)
	return rightIndex
}
