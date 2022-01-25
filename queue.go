package main

import "fmt"

//representation of queue that holds a slice
type Queue struct {
	items []int
}

//Enqueue adds a value to the end
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

//Dequeue removes a value at the front and returns it
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myQueue := Queue{}
	myQueue.Enqueue(6)
	myQueue.Enqueue(98)
	myQueue.Enqueue(0)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}
