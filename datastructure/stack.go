package datastructure

import (
	"errors"
)

type Queue interface{
	Add(item interface{})
	IsEmpty() bool
	Pop() (interface{}, error)
	Peek() (interface{}, error)
}

func NewQueue() Queue{
	return &queueImpl{head: nil, tail:nil}
}

type queueImpl struct{
	head *node
	tail *node
}

type node struct{
	value interface{}
	next *node
}

func (q *queueImpl) IsEmpty() bool{
	return q.head == nil
}

func (q *queueImpl) Add(item interface{}){
	nodeToAdd := &node{
		value: item,
		next: nil,
	}
	if q.head == nil{
		q.head = nodeToAdd
	}else{
		q.tail.next = nodeToAdd
	}

	q.tail = nodeToAdd
}

func (q *queueImpl) Pop() (interface{}, error){
	current, err := q.getCurrent()
	if err == nil{
		q.head = q.head.next
		if q.head == nil{
			q.tail = nil
		}
	}

	return current, err
}

func (q *queueImpl) Peek()  (interface{}, error){
	return q.getCurrent()
}

func (q *queueImpl) getCurrent() (interface{}, error){
	if q.head == nil{
		return nil, errors.New("queue is empty")
	}

	return q.head.value, nil
}
