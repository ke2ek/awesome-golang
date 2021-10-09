package tree

type BSTNode struct {
	Key    int
	Value  interface{}
	left   *BSTNode
	right  *BSTNode
	parent *BSTNode
}

func (this *BSTNode) getSuccessor() *BSTNode {
	node := this
	for node.left != nil {
		node = node.left
	}
	return node
}

func (this *BSTNode) replaceInParent(node *BSTNode) {
	if this.parent != nil {
		if this == this.parent.left {
			this.parent.left = node
		} else {
			this.parent.right = node
		}
	}
	if node != nil {
		node.parent = this.parent
	}
}

func (this *BSTNode) Remove(key int) {
	if key < this.Key {
		this.left.Remove(key)
	} else if key > this.Key {
		this.right.Remove(key)
	} else {
		// Remove the key here.
		if this.left != nil && this.right != nil {
			successor := this.getSuccessor()
			this.Key, this.Value = successor.Key, successor.Value
			successor.Remove(successor.Key)
		} else if this.left != nil {
			this.replaceInParent(this.left)
		} else if this.right != nil {
			this.replaceInParent(this.right)
		} else { // If there is no child
			this.replaceInParent(nil)
		}
	}
}

func (this *BSTNode) Add(key int, value interface{}) {
	if key == this.Key {
		this.Value = value
	} else if key < this.Key {
		if this.left != nil {
			this.left.Add(key, value)
		} else {
			this.left = &BSTNode{Key: key, Value: value}
		}
	} else if key > this.Key {
		if this.right != nil {
			this.right.Add(key, value)
		} else {
			this.right = &BSTNode{Key: key, Value: value}
		}
	}
}

func (this *BSTNode) Search(key int) interface{} {
	if key == this.Key {
		return this.Value
	} else if key < this.Key && this.left != nil {
		return this.left.Search(key)
	} else if key > this.Key && this.right != nil {
		return this.right.Search(key)
	}
	return nil
}

// Common
func TraverseInOrder(node *BSTNode, nodes *([]*BSTNode)) {
	if node == nil {
		return
	}
	TraverseInOrder(node.left, nodes)
	*nodes = append(*nodes, node)
	TraverseInOrder(node.right, nodes)
}

func SearchIteratively(node *BSTNode, key int) interface{} {
	for node != nil {
		if key < node.Key {
			node = node.left
		} else if key > node.Key {
			node = node.right
		} else {
			return node.Value
		}
	}
	return nil
}

func IsBST(node *BSTNode, minKey, maxKey int) bool {
	if node == nil {
		return true
	}
	if node.Key < minKey || node.Key > maxKey {
		return false
	}
	return IsBST(node.left, minKey, node.Key-1) && IsBST(node.right, node.Key+1, maxKey)
}

// Return the root of the trimmed tree in which all the keys are in [low, high].
func Trim(node *BSTNode, low, high int) *BSTNode {
	if node == nil { // base case
		return nil
	}
	if node.Key < low {
		// Don't have to consider the current node and the left subtree because all the keys in them are less than low.
		// But, some keys in the right subtree can be between low and high because of the BST property.
		return Trim(node.right, low, high)
	} else if node.Key > high {
		// Likewise, the current node and the right subtree have keys greater than high.
		// But, some keys in the left subtree can be between low and high because of the BST property.
		return Trim(node.left, low, high)
	} else { // if the current key is between low and high,
		// Alter the pointers of both subtree into the root of each trimmed tree.
		node.left = Trim(node.left, low, high)
		node.right = Trim(node.right, low, high)
	}
	return node
}
