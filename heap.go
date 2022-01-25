package main

import "fmt"

//MaxHeap struct hat a slice that holds an array
type MaxHeap struct {
	array []int
}

//Insert adds an element
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

//Extract returns the largest key, and removes it from heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	last := len(h.array) - 1

	//when the array is empty
	if len(h.array) == 0 {
		fmt.Println("Cannot extract from an empty array")
		return -1
	}

	//take out the last index and put it in the root
	h.array[0] = h.array[last]
	h.array = h.array[:last]

	h.maxHeapifyDown(0)

	return extracted
}

//maxHeapifyUp will heapify from bottom top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyDown will heapify top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	for l <= lastIndex {
		if l == lastIndex { //when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { //when left child is larger
			childToCompare = l
		} else { //when right child is larger
			childToCompare = r
		}
		// compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

//get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

//get the left child index
func left(i int) int {
	return 2*i + 1
}

//get the right child index
func right(i int) int {
	return 2*i + 2
}

//swap keys in array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	m := &MaxHeap{}
	buildHeap := []int{10, 20, 30, 3, 6, 7, 11, 23, 89, 17}
	for _, value := range buildHeap {
		m.Insert(value)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
