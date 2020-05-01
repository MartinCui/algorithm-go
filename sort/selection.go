package sort

func selectionSort(arr []int){
	for i := 0; i < len(arr); i++ {
		minPosition := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minPosition] {
				minPosition = j
			}
		}

		if minPosition != i {
			min := arr[minPosition]
			arr[minPosition] = arr[i]
			arr[i] = min
		}
	}
}
