package sort

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		valueToInsert := arr[i]
		valueInserted := false
		for j := i; j > 0; j-- {
			if arr[j-1] > valueToInsert {
				arr[j] = arr[j-1]
			} else {
				valueInserted = true
				arr[j] = valueToInsert
				break
			}
		}

		if !valueInserted {
			arr[0] = valueToInsert
		}
	}

	return arr
}
