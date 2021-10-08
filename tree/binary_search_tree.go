package tree

type BSTNode struct {
	Key   int
	Value interface{}
	Left  *BSTNode
	Right *BSTNode
}

func (this *BSTNode) getSuccessor() *BSTNode {
	node := this
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func (this *BSTNode) Remove(key int) *BSTNode {
	if key < this.Key {
		this.Left = this.Left.Remove(key)
	} else if key > this.Key {
		this.Right = this.Right.Remove(key)
	} else {
		// Remove the key here.
		if this.Left != nil && this.Right != nil {
			successor := this.getSuccessor()
			this.Key, this.Value = successor.Key, successor.Value
			successor.Remove(successor.Key)
		} else if this.Left != nil {
			return this.Left
		} else if this.Right != nil {
			return this.Right
		} else { // If there is no child
			return nil
		}
	}
	return this
}

func (this *BSTNode) Add(key int, value interface{}) {
	if key == this.Key {
		this.Value = value
	} else if key < this.Key {
		if this.Left != nil {
			this.Left.Add(key, value)
		} else {
			this.Left = &BSTNode{Key: key, Value: value}
		}
	} else if key > this.Key {
		if this.Right != nil {
			this.Right.Add(key, value)
		} else {
			this.Right = &BSTNode{Key: key, Value: value}
		}
	}
}

func (this *BSTNode) Search(key int) interface{} {
	if key == this.Key {
		return this.Value
	} else if key < this.Key && this.Left != nil {
		return this.Left.Search(key)
	} else if key > this.Key && this.Right != nil {
		return this.Right.Search(key)
	}
	return nil
}

// Common
func TraverseInOrder(node *BSTNode, nodes *([]*BSTNode)) {
	if node == nil {
		return
	}
	TraverseInOrder(node.Left, nodes)
	*nodes = append(*nodes, node)
	TraverseInOrder(node.Right, nodes)
}

func SearchIteratively(node *BSTNode, key int) interface{} {
	for node != nil {
		if key < node.Key {
			node = node.Left
		} else if key > node.Key {
			node = node.Right
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
	return IsBST(node.Left, minKey, node.Key-1) && IsBST(node.Right, node.Key+1, maxKey)
}

// Return the root of the trimmed tree in which all the keys are in [low, high].
func Trim(node *BSTNode, low, high int) *BSTNode {
	if node == nil { // base case
		return nil
	}
	if node.Key < low {
		// Don't have to consider the current node and the left subtree because all the keys in them are less than low.
		// But, some keys in the right subtree can be between low and high because of the BST property.
		return Trim(node.Right, low, high)
	} else if node.Key > high {
		// Likewise, the current node and the right subtree have keys greater than high.
		// But, some keys in the left subtree can be between low and high because of the BST property.
		return Trim(node.Left, low, high)
	} else { // if the current key is between low and high,
		// Alter the pointers of both subtree into the root of each trimmed tree.
		node.Left = Trim(node.Left, low, high)
		node.Right = Trim(node.Right, low, high)
	}
	return node
}
