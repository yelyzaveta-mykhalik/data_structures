package main

import fmt "fmt"

type Node struct {
	data int
	next *Node // address of the next node
}

type linkedList struct {
	head   *Node
	length int
}

func (l *linkedList) prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	myList := linkedList{}
	node1 := &Node{data: 36}
	node2 := &Node{data: 17}
	node3 := &Node{data: 44}
	node4 := &Node{data: 98}
	node5 := &Node{data: 56}
	node6 := &Node{data: 11}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)
	myList.printListData()

	myList.deleteWithValue(100)
	myList.deleteWithValue(11)
	myList.printListData()

	emptyLinkedList := linkedList{}
	emptyLinkedList.deleteWithValue(34)
}
