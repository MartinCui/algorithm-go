package search

// arr: non-empty target array
// key: key to find in array
// return: index if key was found, otherwise -1
func interpolationSearch(arr []int, key int) int{
	if len(arr) == 0{
		return -1
	}

	size := len(arr)
	low := 0
	high := size -1
	bestGuessPosition := 0

	for ;arr[low] != arr[high] && key > arr[low] && key < arr[high];{
		bestGuessPosition = low + int((float64(key - arr[low])/float64(arr[high] - arr[low])) * float64(high - low))
		if arr[bestGuessPosition] == key{
			return bestGuessPosition
		} else if arr[bestGuessPosition] > key{
			high = bestGuessPosition - 1
		} else{
			low = bestGuessPosition + 1
		}
	}

	if arr[low] == key{
		return low
	}

	if arr[high] == key{
		return high
	}

	return -1
}