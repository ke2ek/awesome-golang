package tree

import "awesome-golang/common"

/*
AVL Tree is a self-balancing Binary Search Tree (BST)
where the difference between heights of left and right subtrees cannot be more than one for all nodes.
*/

type AVLNode struct {
	Key    int
	Value  interface{}
	left   *AVLNode
	right  *AVLNode
	height int
}

// Return balance factor of the current node.
func (this *AVLNode) getBalance() int {
	return this.left.height - this.right.height
}

func (this *AVLNode) adjustHeight() {
	this.height = common.Max(this.left.height, this.right.height) + 1
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
	1. Insert a new key.
	2. Get the diff of height between the left subtree and the right subtree.
	3. if diff (= left height - right height) > 1, left-left case / left-right case
		3-1. if key < this.left.Key, then left-left case -> rotateRight(the current node).
			It menas that the height of the left subtree > the height of the right subtree in the left subtree.
		3-2. if key > this.left.Key, then left-right case -> rotateLeft(the left child of the current node)
														-> rotateRight(the current node)
			It menas that the height of the left subtree < the height of the right subtree in the left subtree.
	4. if diff (= left height - right height) < -1, right-right case / right-left case
		4-1. if key > this.right.Key, then right-right case -> rotateLeft(the current node)
			It means that the height of the left subtree < the height of the right subtree in the right subtree.
		4-2. if key < this.right.Key, then right-left case -> rotateRight(the right child of the current node)
														-> rotateLeft(the current node)
			It menas that the height of the left subtree > the height of the right subtree in the right subtree.
*/
func (this *AVLNode) Add(key int, value interface{}) *AVLNode {
	// 1. Perform the normal BST insertion.
	if key < this.Key {
		if this.left != nil {
			this.left = this.left.Add(key, value)
		} else {
			this.left = &AVLNode{Key: key, Value: value, height: 1}
			return this
		}
	} else if key > this.Key {
		if this.right != nil {
			this.right = this.right.Add(key, value)
		} else {
			this.right = &AVLNode{Key: key, Value: value, height: 1}
			return this
		}
	} else if key == this.Key {
		this.Value = value
		return this
	}

	// 2. Update height of this ancestor node.
	this.adjustHeight()

	// 3. Get the balance factor of this ancestor node to check whether this node became unbalanced.
	balance := this.getBalance()

	// 4. If this node became unbalanced, deal with the following 4 cases.
	if balance > 1 {
		if key > this.left.Key { // Left-Right case
			this.left = this.left.rotateLeft()
		}
		// if key < this.left.Key, then Left-Left case, so this.rotateRight()
		return this.rotateRight()
	}

	if balance < -1 {
		if key < this.right.Key { // Right-Left case
			this.right = this.right.rotateRight()
		}
		// if key > this.right.Key, then Right-Right case, so this.rotateLeft()
		return this.rotateLeft()
	}

	return this
}

func (this *AVLNode) Remove(key int) *AVLNode {
	return nil
}

func (this *AVLNode) Search(key int) interface{} {
	if key == this.Key {
		return this.Value
	} else if key < this.Key && this.left != nil {
		return this.left.Search(key)
	} else if key > this.Key && this.right != nil {
		return this.right.Search(key)
	}
	return nil
}
