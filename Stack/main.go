package main

import (
	"fmt"
	"sync"
)

func main() {
	stack := NewStack()
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)
	stack.push(5)

	size := len(stack.values)
	for i := 0; i < size; i++ {
		pop := stack.pop()
		if i%2 == 0 {
			fmt.Println(pop)
		}
	}
}

type StackImplementation interface {
	push(value int)
	pop() int
}

type Stack struct {
	values []int
	mutex  sync.Mutex
}

func NewStack() *Stack {
	return &Stack{
		values: nil,
		mutex:  sync.Mutex{},
	}
}

func (s *Stack) push(value int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.values = append(s.values, value)
}

func (s *Stack) pop() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if len(s.values) == 0 {
		return 0
	}

	top := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return top
}
