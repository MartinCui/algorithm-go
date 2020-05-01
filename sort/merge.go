package sort

func mergeSort(arr []int) {
	aux := make([]int, len(arr))
	_sort(arr, aux, 0, len(arr) - 1)
}

func _sort(arr, aux []int, low, high int) {
	if low == high{
		return
	}

	if high - low == 1 && arr[low] > arr[high]{
		tmp := arr[high]
		arr[high] = arr[low]
		arr[low] = tmp
	}else{
		middle := low + ((high - low) / 2)
		_sort(arr, aux, low, middle)
		_sort(arr, aux, middle + 1, high)
		_merge(arr, aux, low, middle, high)
	}
}

func _merge(arr, aux []int, low, mid, high int) {
	for i := low; i <= high; i++ {
		aux[i] = arr[i]
	}

	leftIndex := low
	rightIndex := mid + 1
	for i := low; i <= high; i++ {
		if leftIndex > mid {
			arr[i] = aux[rightIndex]
			rightIndex++
		} else if rightIndex > high {
			arr[i] = aux[leftIndex]
			leftIndex++
		} else if aux[leftIndex] < aux[rightIndex] {
			arr[i] = aux[leftIndex]
			leftIndex++
		} else {
			arr[i] = aux[rightIndex]
			rightIndex++
		}
	}
}
