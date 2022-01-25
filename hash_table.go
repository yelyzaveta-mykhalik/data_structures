package main

import "fmt"

const ARRAYSIZE = 7

//HashTable structure
type HashTable struct {
	array [ARRAYSIZE]*bucket
}

//bucket structure
type bucket struct {
	head *bucketNode
}

//bucketNode
type bucketNode struct {
	key  string
	next *bucketNode
}

//Insert takes a key and add it to the hash array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

//Search
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

//Delete
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

//insert
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exists")
	}
}

//search
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

//delete
func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

//hashFunction
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ARRAYSIZE
}

//init creates a bucket in each slot of hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	hashTable := Init()
	list := []string{
		"ERIC",
		"TOLIK",
		"KYLIE",
		"TOKEN",
		"JIMMY",
		"ANDREW",
	}

	for _, v := range list {
		hashTable.Insert(v)
		fmt.Println(hashTable)
	}

}
