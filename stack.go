package main

import "fmt"

//represents Stack that hold a slice
type Stack struct {
	items []int
}

//Push add a value at the back
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

//Pop remove a value at the end and returns it
func (s *Stack) Pop() int {
	last := len(s.items) - 1
	toRemove := s.items[last]
	s.items = s.items[:last]
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(12)
	myStack.Push(78)
	myStack.Push(99)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}
