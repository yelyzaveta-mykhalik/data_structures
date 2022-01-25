package main

import "fmt"

//Node
type node struct {
	Key   int
	Left  *node
	Right *node
}

//Insert adds a node to the tree
func (n *node) Insert(k int) {
	if n.Key < k {
		//move right
		if n.Right == nil {
			n.Right = &node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		//move left
		if n.Left == nil {
			n.Left = &node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

//Search takes a key value and returns true if there is a node with the same value
func (n *node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key < k {
		//move right
		return n.Right.Search(k)
	} else if n.Key > k {
		// move left
		return n.Left.Search(k)
	}
	return true
}

func main() {
	tree := &node{Key: 99}
	tree.Insert(200)
	tree.Insert(1)
	tree.Insert(17)
	tree.Insert(43)
	tree.Insert(234)
	tree.Insert(98)
	tree.Insert(115)
	tree.Insert(179)
	tree.Insert(12)
	tree.Insert(8)
	fmt.Println(tree.Search(12))
}
