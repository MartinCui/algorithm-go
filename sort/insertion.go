package sort

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			tmp := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = tmp
		}
	}

	return arr
}
