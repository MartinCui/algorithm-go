package datastructure

import (
	common "github.com/martincui/algorithm"
)

type heapPriorityQueue struct {
	binaryHeap []int
	size       int
}

func newHeapPriorityQueue() *heapPriorityQueue {
	return &heapPriorityQueue{
		binaryHeap: make([]int, 1),
		size:       0,
	}
}

func (pq *heapPriorityQueue) add(v int) {
	if pq.size == len(pq.binaryHeap)-1 {
		pq.binaryHeap = append(pq.binaryHeap, v)
	} else {
		pq.binaryHeap[pq.size+1] = v
	}

	pq.size++
	pq.swimUp(pq.size)
}

func (pq *heapPriorityQueue) popMax() int {
	if pq.size <= 0 {
		panic("popping up empty PQ")
	}

	max := pq.binaryHeap[1]
	pq.binaryHeap[1] = pq.binaryHeap[pq.size]
	pq.binaryHeap[pq.size] = 0
	pq.size--
	pq.sinkDown(1)

	return max
}

func (pq *heapPriorityQueue) sinkDown(position int) {
	for currentPosition := position; currentPosition < pq.size; {
		currentNodeValue := pq.binaryHeap[currentPosition]
		leftChildPosition := 2 * currentPosition
		rightChildPosition := leftChildPosition + 1
		if leftChildPosition > pq.size{
			break
		}
		if rightChildPosition > pq.size{
			if currentNodeValue < pq.binaryHeap[leftChildPosition]{
				common.Exchange(pq.binaryHeap, currentPosition, leftChildPosition)
			}

			break
		}

		maxChildPosition := leftChildPosition
		if pq.binaryHeap[rightChildPosition] > pq.binaryHeap[leftChildPosition]{
			maxChildPosition = rightChildPosition
		}
		if currentNodeValue >= pq.binaryHeap[maxChildPosition]{
			break
		}

		common.Exchange(pq.binaryHeap, currentPosition, maxChildPosition)
		currentPosition = maxChildPosition
	}
}

func (pq *heapPriorityQueue) swimUp(position int) {
	for currentPosition := position; ; {
		abovePosition := currentPosition / 2
		if abovePosition > 0 && pq.binaryHeap[abovePosition] < pq.binaryHeap[currentPosition] {
			common.Exchange(pq.binaryHeap, currentPosition, abovePosition)
			currentPosition = abovePosition
		} else {
			break
		}
	}
}
