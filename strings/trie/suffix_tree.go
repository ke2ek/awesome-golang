package trie

import (
	"fmt"
	"sort"
)

/*
Suffix tree is a compressed trie containing all the suffixes of the given text as their keys and indices in the text as their values.

The construction of such a tree for the string S takes time and space linear in the length of S.
Suffix trees also provide one of the first linear-time solutions for the longest common substring problem.

The suffix tree for the string S of length n is defined as a tree such that:
1. The tree has exactly n leaves numbered from 1 to n.
2. Except for the root, every internal node has at least two children.
3. Each edge is labelled with a non-empty substring of S.
4. No two edges starting out of a node can have string-labels beginning with the same character.
5. The string obtained by concatenating all the string-labels found on the path from the root to leaf i spells out suffix S[i..n] for i from 1 to n.

# of leaf nodes = M = the number of suffixes of a given string
# of non-leaf nodes <= M-1 (because an internal node of a suffix tree has at least 2 children)
# of total nodes <= 2M-1
So, space time complexity will be O(M).

Ideas:
1. This data structure starts from the idea that it reduce some paths without branching from a node to a leaf node
by each path which has only one edge.
	- Each label of edges, that are compressed, equals the concatenated edge labels from the node to the leaf node.
2. Labling edges as suffixes is not efficient because of storing strings,
so it will be better to label edges as (offset, length) pair which means the start index and the length of each suffix.
3. Each leaf node stores the start index of the suffix that the path represents.
4. Each leaf node means the end of substring(=$) because the tree needs to distinguish overlapping substrings.
	e.g., "abcd" and "ab" are on the same edge if there is no use of $.
	(Before)
	root - abcd$
	(After)
	root - ab - cd$
		  - $

How to build -> Ukkonen’s Algorithm
1. Whenever adding a suffix to the tree, check if there is another suffix,
the prefix of which is the same as the suffix to be inserted.
2. If it exists, branch that node into two nodes, the node with $ (end of string) and the node with other letters,
where other letters are from the rest of the existing suffix except the prefix(=suffix to be inserted).
3. Suffix Links
e.g., one edge with path-label xA, where x is a character and A denotes substring, connects root node and ndoe W via node V
and another edge with path-label A connects node X and Y.
In this case, the edge going from node V to node X is so-called a suffix link.
	root --x-->  V --A--> W
	             | <<<<<<<<<<<<< this is a suffix link
	      ...  - X --A--> Y

If A is empty string, then a suffix link of that node will go to root node.
4. Edge-label Compression

*/

type suffixTreeNode struct {
	start       int
	end         *int
	suffixIndex int
	suffixLink  *suffixTreeNode
	children    map[int]*suffixTreeNode
}

type SuffixTree struct {
	/* lastNewNode will point to newly created internal node, waiting for it's suffix link to be set, which might get
	a new suffix link (other than root) in next extension of same phase. lastNewNode will be set to NULL when last
	newly created internal node (if there is any) got it's suffix link reset to new internal node created in next
	extension of same phase. */
	root                 *suffixTreeNode
	lastNewNode          *suffixTreeNode
	activeNode           *suffixTreeNode
	count                int
	activeEdge           int // activeEdge is represented as input string character index (not the character itself)
	activeLength         int
	remainingSuffixCount int // remainingSuffixCount tells how many suffixes yet to be added in tree
	leafEnd              int
	rootEnd              int
	splitEnd             int
	size                 int // Length of input string
	text                 string
}

func (this *suffixTreeNode) getSortedKeys() []int {
	keys := make([]int, 0, len(this.children))
	for k := range this.children {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func (this *suffixTreeNode) edgeLength() int {
	return (*this.end) - this.start + 1
}

func (this *SuffixTree) newNode(start int, end *int) *suffixTreeNode {
	this.count++
	node := &suffixTreeNode{
		start:       start,
		end:         end,
		suffixIndex: -1,
		suffixLink:  this.root,
		children:    map[int]*suffixTreeNode{},
	}
	return node
}

/* activePoint change for walk down (APCFWD) using Skip/Count Trick (Trick 1).
If activeLength is greater than current edge length, set next internal node as activeNode
and adjust activeEdge and activeLength accordingly to represent same activePoint */
func (this *SuffixTree) walkDown(current *suffixTreeNode) bool {
	edgeLen := current.edgeLength()
	if this.activeLength >= edgeLen {
		this.activeEdge = int(this.text[this.activeEdge+edgeLen] - ' ')
		this.activeLength -= edgeLen
		this.activeNode = current
		return true
	}
	return false
}

func (this *SuffixTree) extendSuffixTree(pos int) {
	// Extension Rule 1, this takes care of extending all leaves created so far in tree.
	this.leafEnd = pos
	// Indicates that a new suffix added to the list of suffixes yet to be added in tree.
	this.remainingSuffixCount++
	// While starting a new phase, indicates there is no internal node waiting for it's suffix link reset in current phase.
	this.lastNewNode = nil

	// Add all suffixes (yet to be added) one by one in tree.
	for this.remainingSuffixCount > 0 {
		if this.activeLength == 0 {
			this.activeEdge = int(this.text[pos] - ' ')
		}
		if this.activeNode.children[this.activeEdge] == nil {
			// If there is no outgoing edge starting with activeEdge from activeNode,
			// Extension Rule 2 (A new leaf edge gets created)
			this.activeNode.children[this.activeEdge] = this.newNode(pos, &this.leafEnd)

			// A new leaf edge is created in above line starting from an existing node (the current activeNode),
			// and if there is any internal node waiting for it's suffix link get reset,
			// point the suffix link from that last internal node to current activeNode.
			// Then set lastNewNode to NULL indicating no more node waiting for suffix link reset.
			if this.lastNewNode != nil {
				this.lastNewNode.suffixLink = this.activeNode
				this.lastNewNode = nil
			}
		} else {
			// If there is an outgoing edge starting with activeEdge from activeNode,
			// Get the next node at the end of edge starting with activeEdge.
			next := this.activeNode.children[this.activeEdge]
			if this.walkDown(next) {
				// If it is walked down, start from next node (the new activeNode).
				continue
			}

			// Extension Rule 3 (current character being processed is already on the edge)
			if this.text[next.start+this.activeLength] == this.text[pos] {
				// If a newly created node waiting for it's suffix link to be set,
				// then set suffix link of that waiting node to current active node.
				if this.lastNewNode != nil && this.activeNode != this.root {
					this.lastNewNode.suffixLink = this.activeNode
					this.lastNewNode = nil
				}
				this.activeLength++
				// Stop all further processing in this phase and move on to next phase.
				break
			}

			// Here is when activePoint is in middle of the edge being traversed
			// and current character being processed is not on the edge (we fall off the tree).
			// In this case, it has to add a new internal node and a new leaf edge going out of that new node.
			// This is Extension Rule 2, where a new leaf edge and a new internal node get created.
			this.splitEnd = next.start + this.activeLength - 1
			node := this.newNode(next.start, &this.splitEnd)
			this.activeNode.children[this.activeEdge] = node
			// New leaf coming out of new internal node
			node.children[int(this.text[pos])] = this.newNode(pos, &this.leafEnd)
			next.start += this.activeLength
			node.children[this.activeEdge] = next

			// If there is any internal node created in last extensions of same phase
			// which is still waiting for it's suffix link reset, do it now.
			if this.lastNewNode != nil {
				this.lastNewNode.suffixLink = node
			}

			// Make the current newly created internal node waiting for it's suffix link reset
			// (which is pointing to root at present).
			// If we come across any other internal node (existing or newly created) in next extension of same phase
			// when a new leaf edge gets added (i.e. when Extension Rule 2 applies is any of the next extension of
			// same phase) at that point, suffixLink of this node will point to that internal node.
			this.lastNewNode = node
		}

		// One suffix got added in tree.
		this.remainingSuffixCount--
		if this.activeNode == this.root && this.activeLength > 0 {
			this.activeLength--
			this.activeEdge = int(this.text[pos-this.remainingSuffixCount+1] - ' ')
		} else if this.activeNode != this.root {
			this.activeNode = this.activeNode.suffixLink
		}
	}
}

/* Print the suffix tree as well along with setting suffix index, so tree will be printed in DFS manner.
Each edge along with it's suffix index will be printed. */
func (this *SuffixTree) setSuffixIndexByDFS(node *suffixTreeNode, labelHeight int) {
	if node == nil {
		return
	}
	if node.start != -1 {
		fmt.Printf("text[%d:%d]=", node.start, *node.end)
		fmt.Print(this.text[node.start : *node.end+1])
	}
	isLeaf := true
	keys := node.getSortedKeys()
	for _, key := range keys {
		if node.children[key] != nil {
			if isLeaf && node.start != -1 {
				fmt.Printf(" [%d]\n", node.suffixIndex)
			}
			// Current node is not a leaf as it has outgoing edges from it.
			isLeaf = false
			this.setSuffixIndexByDFS(node.children[key], labelHeight+node.children[key].edgeLength())
		}
	}
	if isLeaf {
		node.suffixIndex = this.size - labelHeight
		fmt.Printf(" [%d]\n", node.suffixIndex)
	}
}

func (this *SuffixTree) freeSuffixTreeByPostOrder(node *suffixTreeNode) {
	if node == nil {
		return
	}
	keys := node.getSortedKeys()
	for _, key := range keys {
		if node.children[key] != nil {
			this.freeSuffixTreeByPostOrder(node.children[key])
			node.children[key] = nil
		}
	}
}

func (this *SuffixTree) Count() int {
	return this.count
}

func NewSuffixTree(s string) *SuffixTree {
	tree := &SuffixTree{activeEdge: -1, leafEnd: -1, size: len(s), text: s}
	tree.rootEnd = -1
	tree.root = tree.newNode(-1, &tree.rootEnd)
	tree.activeNode = tree.root
	for i := 0; i < tree.size; i++ {
		tree.extendSuffixTree(i)
	}
	tree.setSuffixIndexByDFS(tree.root, 0)
	tree.freeSuffixTreeByPostOrder(tree.root)
	return tree
}
