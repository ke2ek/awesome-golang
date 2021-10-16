package trie

/*
Trie is a tree data structure used for locating specific keys from within a set.
These keys are most often strings, with links between nodes defined not by the entire key, but by individual characters.

The trie is a tree of nodes which supports Find and Insert operations.
Find returns the value for a key string, and Insert inserts a string (the key) and a value into the trie.
Both Insert and Find run in O(M) time, where M is the length of the key.
*/

type trieNode struct {
	terminal bool
	children map[int]*trieNode
}

func (this *trieNode) insert(s *string, pos int) {
	if pos == len(*s) {
		this.terminal = true
		return
	}
	key := int((*s)[pos])
	if this.children[key] == nil {
		this.children[key] = &trieNode{children: map[int]*trieNode{}}
	}
	this.children[key].insert(s, pos+1)
}

func (this *trieNode) find(s *string, pos int) bool {
	if pos == len(*s) {
		return this.terminal
	}
	key := int((*s)[pos])
	if this.children[key] == nil {
		return false
	}
	return this.children[key].find(s, pos+1)
}

type Trie struct {
	root *trieNode
}

func New() *Trie {
	return &Trie{root: &trieNode{children: map[int]*trieNode{}}}
}

func (this *Trie) Insert(s string) {
	this.root.insert(&s, 0)
}

func (this *Trie) Find(s string) bool {
	return this.root.find(&s, 0)
}
