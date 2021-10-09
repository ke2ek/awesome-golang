package test

import (
	tree "awesome-golang/binary-tree"
	"awesome-golang/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchTree(t *testing.T) {
	data := common.MakeData()
	root := tree.NewBSTNode(8, nil, nil, nil)
	insertedOrder := []int{8, 4, 12, 2, 1, 3, 5, 6, 7, 9, 23, 14, 27}
	for _, key := range insertedOrder {
		root.Add(key, data[key])
	}

	nodes := make([]interface{}, 0)
	tree.InOrder(root, &nodes)
	ans := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 14, 23, 27}
	for i, node := range nodes {
		assert.Equal(t, data[ans[i]], node.(*common.Person))
		result := tree.Search(root, ans[i])
		assert.Equal(t, data[ans[i]].Name, result.(*common.Person).Name)
	}

	removed := []int{6, 27, 23, 4}
	for _, target := range removed {
		root = root.Remove(target).(*tree.BSTNode)
		assert.Equal(t, nil, tree.Search(root, target))
	}

	nodes = make([]interface{}, 0)
	tree.InOrder(root, &nodes)
	ans = []int{1, 2, 3, 5, 7, 8, 9, 12, 14}
	for i, node := range nodes {
		assert.Equal(t, ans[i], node.(*common.Person).No)
	}

	// Remove a root.
	root = root.Remove(8).(*tree.BSTNode)
	nodes = make([]interface{}, 0)
	tree.InOrder(root, &nodes)
	ans = []int{1, 2, 3, 5, 7, 9, 12, 14}
	for i, node := range nodes {
		assert.Equal(t, ans[i], node.(*common.Person).No)
	}

	// Check if it is a binary search tree.
	assert.Equal(t, true, tree.IsBST(root, 1, 14))
	assert.Equal(t, false, tree.IsBST(root, 2, 13))

	// Search interatively.
	ans = []int{3, 5, 14}
	for _, key := range ans {
		result := tree.SearchIteratively(root, key)
		assert.Equal(t, data[key].Name, result.(*common.Person).Name)
	}
	assert.Equal(t, nil, tree.SearchIteratively(root, 8))

	// Trim nodes out of [3, 10].
	root.Add(10, &common.Person{No: 10, Name: "Jay", Age: 24})
	root.Add(11, &common.Person{No: 11, Name: "Ray", Age: 31})
	root = tree.TrimBST(root, 3, 10)
	nodes = make([]interface{}, 0)
	tree.InOrder(root, &nodes)
	ans = []int{3, 5, 7, 9, 10}
	for i, node := range nodes {
		assert.Equal(t, ans[i], node.(*common.Person).No)
	}
}
