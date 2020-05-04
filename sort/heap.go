package sort

import common "github.com/martincui/algorithm"

func heapSort(arr []int) {
	_heapSortConvertToHeap(arr)

	for toSortLength := len(arr); toSortLength > 1; toSortLength-- {
		endPosition := toSortLength - 1
		common.Exchange(arr, endPosition, 0)
		_heapSortSink(arr, 0, endPosition-1)
	}
}

func _heapSortConvertToHeap(arr []int) {
	for i := len(arr) / 2; i >= 0; i-- {
		_heapSortSink(arr, i, len(arr) - 1)
	}
}

func _heapSortSink(arr []int, position int, bottomPosition int) {
	for currentPosition := position; (currentPosition * 2) < bottomPosition; {
		leftChildPosition := (currentPosition * 2) + 1
		rightChildPosition := leftChildPosition + 1
		if rightChildPosition > bottomPosition {
			if arr[currentPosition] < arr[leftChildPosition] {
				common.Exchange(arr, currentPosition, leftChildPosition)
			}
			break
		}

		maxChildPosition := leftChildPosition
		if arr[rightChildPosition] > arr[leftChildPosition] {
			maxChildPosition = rightChildPosition
		}

		if arr[currentPosition] >= arr[maxChildPosition] {
			break
		}

		if arr[currentPosition] < arr[maxChildPosition] {
			common.Exchange(arr, currentPosition, maxChildPosition)
			currentPosition = maxChildPosition
		}
	}
}
