package main

import "fmt"

type stack struct {
	arr []int
}

func (s *stack) Push(num int) {
	s.arr = append(s.arr, num)
}

func (s *stack) Top() int {
	size := len(s.arr)
	if size == 0 {
		return -1
	}
	element := s.arr[size-1]
	s.arr = s.arr[:size-1]
	return element
}

func main() {
	myStack := &stack{}
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	fmt.Println(myStack.Top())
	fmt.Println(myStack.Top())
	fmt.Println(myStack.Top())
	fmt.Println(myStack.Top())
}
