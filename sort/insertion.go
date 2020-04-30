package sort

func binarySearchInsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		v := arr[i]
		insertPosition := binarySearchInsertPosition(arr, i, v)
		if insertPosition == i {
			continue
		}

		for j := i; j > insertPosition; j-- {
			arr[j] = arr[j-1]
		}

		arr[insertPosition] = v
	}

	return arr
}

func binarySearchInsertPosition(arr []int, elementCount int, v int) int {
	low := 0
	high := elementCount - 1

	for ; low < high && arr[low] <= v && arr[high] >= v; {
		middlePosition := low + int(float32(high-low)/2)
		middleValue := arr[middlePosition]
		if middleValue > v {
			high = middlePosition - 1
		} else if middleValue < v {
			low = middlePosition + 1
		} else {
			return middlePosition + 1
		}
	}

	if arr[low] > v {
		return low
	} else if arr[low] == v {
		return low + 1
	} else if arr[high] < v {
		return high + 1
	} else {
		return high
	}
}
