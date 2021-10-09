package tree

import (
	"awesome-golang/common"
	"strconv"
)

type BST interface {
	I
	Add(key int, value interface{}) interface{}
	Remove(key int) interface{}
}

type BSTNode struct {
	key   int
	value interface{}
	left  *BSTNode
	right *BSTNode
}

func NewBSTNode(key int, value interface{}, left, right *BSTNode) *BSTNode {
	return &BSTNode{key: key, value: value, left: left, right: right}
}

func (this *BSTNode) Key() int           { return this.key }
func (this *BSTNode) Value() interface{} { return this.value }
func (this *BSTNode) Left() interface{}  { return this.left }
func (this *BSTNode) Right() interface{} { return this.right }

func (this *BSTNode) Add(key int, value interface{}) interface{} {
	if key == this.key {
		this.value = value
	} else if key < this.key {
		if this.left != nil {
			this.left.Add(key, value)
		} else {
			this.left = &BSTNode{key: key, value: value}
		}
	} else if key > this.key {
		if this.right != nil {
			this.right.Add(key, value)
		} else {
			this.right = &BSTNode{key: key, value: value}
		}
	}
	return (*BSTNode)(nil)
}

func (this *BSTNode) Remove(key int) interface{} {
	if key < this.key {
		this.left = this.left.Remove(key).(*BSTNode)
	} else if key > this.key {
		this.right = this.right.Remove(key).(*BSTNode)
	} else {
		if key != this.key {
			panic("Remove(): it has no key = " + strconv.Itoa(key))
		}
		// Remove the key here.
		if this.left != nil && this.right != nil {
			// Find the successor.
			successorParent, successor := this, this.right
			for successor.left != nil {
				successorParent = successor
				successor = successor.left
			}
			// Move the successor to this.
			this.key, this.value = successor.key, successor.value
			if successorParent != this {
				successorParent.left = successor.Remove(successor.key).(*BSTNode)
			} else {
				successorParent.right = successor.Remove(successor.key).(*BSTNode)
			}
		} else if this.left != nil {
			return this.left
		} else if this.right != nil {
			return this.right
		} else { // If there is no child
			return (*BSTNode)(nil)
		}
	}
	return this
}

// Return the root of the trimmed tree in which all the keys are in [low, high].
func TrimBST(node *BSTNode, low, high int) *BSTNode {
	if node == nil { // base case
		return nil
	}
	if node.key < low {
		// Don't have to consider the current node and the left subtree because all the keys in them are less than low.
		// But, some keys in the right subtree can be between low and high because of the BST property.
		return TrimBST(node.right, low, high)
	} else if node.key > high {
		// Likewise, the current node and the right subtree have keys greater than high.
		// But, some keys in the left subtree can be between low and high because of the BST property.
		return TrimBST(node.left, low, high)
	} else { // if the current key is between low and high,
		// Alter the pointers of both subtree into the root of each trimmed tree.
		node.left = TrimBST(node.left, low, high)
		node.right = TrimBST(node.right, low, high)
	}
	return node
}

// Common
func Search(root BST, target int) interface{} {
	if common.IsNil(root) {
		return nil
	}
	if target == root.Key() {
		return root.Value()
	} else if target < root.Key() {
		return Search(root.Left().(BST), target)
	} else if target > root.Key() {
		return Search(root.Right().(BST), target)
	}
	return nil
}

func SearchIteratively(node BST, target int) interface{} {
	for !common.IsNil(node) {
		if target < node.Key() {
			node = node.Left().(BST)
		} else if target > node.Key() {
			node = node.Right().(BST)
		} else {
			return node.Value()
		}
	}
	return nil
}

func IsBST(root BST, minKey, maxKey int) bool {
	if common.IsNil(root) {
		return true
	}
	if root.Key() < minKey || root.Key() > maxKey {
		return false
	}
	return IsBST(root.Left().(BST), minKey, root.Key()-1) && IsBST(root.Right().(BST), root.Key()+1, maxKey)
}
