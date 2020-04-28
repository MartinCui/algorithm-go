package datastructure

import (
	"errors"
)

type Stack interface{
	Push(item interface{})
	IsEmpty() bool
	Pop() (interface{}, error)
	Peek() (interface{}, error)
}

func NewStack() Stack{
	return &stackImpl{
		storage: nil,
		index:   -1,
	}
}

type stackImpl struct{
	storage []interface{}
	index int
}

func (s *stackImpl) IsEmpty() bool{
	return s.index < 0
}

func (s *stackImpl) Push(item interface{}){
	if s.index < len(s.storage) - 1{
		s.storage[s.index + 1] = item
	}else{
		s.storage = append(s.storage, item)
	}

	s.index++
}

func (s *stackImpl) Pop() (interface{}, error){
	current, err := s.getCurrent()
	if err == nil{
		s.storage[s.index] = nil
		s.index--
	}

	return current, err
}

func (s *stackImpl) Peek()  (interface{}, error){
	return s.getCurrent()
}

func (s *stackImpl) getCurrent() (interface{}, error){
	if s.index < 0{
		return nil, errors.New("stack is empty")
	}

	return s.storage[s.index], nil
}
