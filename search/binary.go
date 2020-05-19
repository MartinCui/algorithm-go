package search

// arr: sorted target array
// key: key to search in arr
// returns: index in arr if key is found in array, otherwise -1
func binarySearch(arr []int, key int) int {
	if len(arr) == 0 {
		return -1
	}

	return _binarySearch(arr, key, 0, len(arr) - 1)
}

func _binarySearch(arr []int, key int, low, high int) int{
	for low <= high && key >= arr[low] && key <= arr[high] {
		middle := low + ((high - low) / 2)
		if arr[middle] > key {
			high = middle - 1
		} else if arr[middle] < key {
			low = middle + 1
		} else {
			return middle
		}
	}

	return -1
}