package main

import "fmt"

const ALPHABETSIZE = 26

// represents each node in the trie and holds an array
type Nodes struct {
	children [ALPHABETSIZE]*Nodes
	isEnd    bool
}

type Trie struct {
	root *Nodes
}

//creates a new trie
func InitTrie() *Trie {
	result := &Trie{root: &Nodes{}}
	return result
}

//Insert takes a word and add it to the trie
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Nodes{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd {
		return true
	}

	return false
}

func main() {
	myTrie := InitTrie()
	toAdd := []string{
		"aragon",
		"aragorn",
		"argon",
		"oregon",
		"oregano",
		"oreo",
	}

	for _, v := range toAdd {
		myTrie.Insert(v)
	}

	fmt.Println(myTrie.Search("hello"))
}
