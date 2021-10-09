package tree

import (
	"awesome-golang/common"
	"strconv"
)

/*
AVL Tree is a self-balancing Binary Search Tree (BST)
where the difference between heights of left and right subtrees cannot be more than one for all nodes.
*/

type AVLNode struct {
	key    int
	value  interface{}
	left   *AVLNode
	right  *AVLNode
	height int
}

func NewAVLNode(key int, value interface{}) *AVLNode {
	return &AVLNode{key: key, value: value, height: 1}
}

func (this *AVLNode) Key() int           { return this.key }
func (this *AVLNode) Value() interface{} { return this.value }
func (this *AVLNode) Left() interface{}  { return this.left }
func (this *AVLNode) Right() interface{} { return this.right }

// Return balance factor of the current node.
func (this *AVLNode) getBalance() int {
	leftH, rightH := 0, 0
	if !common.IsNil(this.left) {
		leftH = this.left.height
	}
	if !common.IsNil(this.right) {
		rightH = this.right.height
	}
	return leftH - rightH
}

func (this *AVLNode) adjustHeight() {
	if common.IsNil(this.left) && common.IsNil(this.right) {
		this.height = 1
	} else if common.IsNil(this.left) {
		this.height = this.right.height + 1
	} else if common.IsNil(this.right) {
		this.height = this.left.height + 1
	} else {
		this.height = common.Max(this.left.height, this.right.height) + 1
	}
}

func (this *AVLNode) rotateLeft() (r *AVLNode) {
	r = this.right
	this.right = r.left
	r.left = this
	this.adjustHeight()
	r.adjustHeight()
	return
}

func (this *AVLNode) rotateRight() (l *AVLNode) {
	l = this.left
	this.left = l.right
	l.right = this
	this.adjustHeight()
	l.adjustHeight()
	return
}

/*
	1. Get the diff of height between the left subtree and the right subtree.
	2. if diff (= left height - right height) > 1, left-left case / left-right case
		3-1. if key < this.left.key, then left-left case -> rotateRight(the current node).
			It menas that the height of the left subtree > the height of the right subtree in the left subtree.
		3-2. if key > this.left.key, then left-right case -> rotateLeft(the left child of the current node)
														-> rotateRight(the current node)
			It menas that the height of the left subtree < the height of the right subtree in the left subtree.
	3. if diff (= left height - right height) < -1, right-right case / right-left case
		4-1. if key > this.right.key, then right-right case -> rotateLeft(the current node)
			It means that the height of the left subtree < the height of the right subtree in the right subtree.
		4-2. if key < this.right.key, then right-left case -> rotateRight(the right child of the current node)
														-> rotateLeft(the current node)
			It menas that the height of the left subtree > the height of the right subtree in the right subtree.
*/
func (this *AVLNode) rotate(key int) *AVLNode {
	// Update height of this ancestor node.
	this.adjustHeight()

	// Get the balance factor of this ancestor node to check whether this node became unbalanced.
	balance := this.getBalance()

	// If this node became unbalanced, deal with the following 4 cases.
	if balance > 1 {
		if key > this.left.key { // Left-Right case
			this.left = this.left.rotateLeft()
		}
		// if key < this.left.key, then Left-Left case, so this.rotateRight()
		return this.rotateRight()
	}

	if balance < -1 {
		if key < this.right.key { // Right-Left case
			this.right = this.right.rotateRight()
		}
		// if key > this.right.key, then Right-Right case, so this.rotateLeft()
		return this.rotateLeft()
	}
	return this
}

func (this *AVLNode) Add(key int, value interface{}) interface{} {
	// Perform the normal BST insertion.
	if key < this.key {
		if this.left != nil {
			this.left = this.left.Add(key, value).(*AVLNode)
		} else {
			this.left = NewAVLNode(key, value)
			return this
		}
	} else if key > this.key {
		if this.right != nil {
			this.right = this.right.Add(key, value).(*AVLNode)
		} else {
			this.right = NewAVLNode(key, value)
			return this
		}
	} else if key == this.key {
		this.value = value
		return this
	}
	return this.rotate(key)
}

/*
	1. Remove the key.
	2. Same as Add().
*/
func (this *AVLNode) Remove(key int) interface{} {
	// Perform the normal BST deletion.
	if key < this.key {
		this.left = this.left.Remove(key).(*AVLNode)
	} else if key > this.key {
		this.right = this.right.Remove(key).(*AVLNode)
	} else {
		if key != this.key {
			panic("Remove(): it has no key the same as " + strconv.Itoa(key))
		}
		// Remove the key here.
		if this.left != nil && this.right != nil {
			parent, child := GetSuccessor(this)
			successorParent, successor := parent.(*AVLNode), child.(*AVLNode)
			if successorParent != this {
				successorParent.left = successor.Remove(successor.key).(*AVLNode)
			} else {
				successorParent.right = successor.Remove(successor.key).(*AVLNode)
			}
			this.key, this.value = successor.key, successor.value
		} else if this.left != nil {
			return this.left
		} else if this.right != nil {
			return this.right
		} else {
			return (*AVLNode)(nil)
		}
	}
	return this.rotate(key)
}

type AVLTree struct {
	root *AVLNode
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

func (this *AVLTree) AddNode(key int, value interface{}) {
	if this.root == nil {
		this.root = NewAVLNode(key, value)
	} else {
		this.root = this.root.Add(key, value).(*AVLNode)
	}
}

func (this *AVLTree) RemoveNode(target int) {
	this.root = this.root.Remove(target).(*AVLNode)
}

func (this *AVLTree) Search(target int) interface{} {
	return Search(this.root, target)
}

func (this *AVLTree) GetInOrder() []interface{} {
	nodes := make([]interface{}, 0)
	InOrder(this.root, &nodes)
	return nodes
}

func (this *AVLTree) Validate(min, max int) bool {
	return IsBST(this.root, min, max)
}
