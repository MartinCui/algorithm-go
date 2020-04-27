package search

// arr: sorted target array
// key: key to search in arr
// returns: index in arr if key is found in array, otherwise -1
func binarySearch(arr []int, key int) int {
	//21:03
	if len(arr) == 0 {
		return -1
	}

	low := 0
	high := len(arr) - 1
	for low < high && key >= arr[low] && key <= arr[high] {
		binaryIndex := low + ((high - low) / 2)
		if arr[binaryIndex] == key {
			return binaryIndex
		} else if arr[binaryIndex] > key {
			high = binaryIndex - 1
		} else {
			low = binaryIndex + 1
		}
	}

	if arr[low] == key {
		return low
	}

	return -1
}
