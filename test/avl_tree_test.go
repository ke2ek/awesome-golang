package test

import (
	tree "awesome-golang/binary-tree"
	"awesome-golang/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAVLTree(t *testing.T) {
	data := common.MakeData()
	avl := tree.NewAVLTree()
	insertedOrder := []int{1, 2, 3, 5, 7, 9, 12, 6, 14, 8, 4, 23, 27}
	for _, key := range insertedOrder {
		avl.AddNode(key, data[key])
	}

	nodes := avl.GetInOrder()
	ans := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 14, 23, 27}
	for i, node := range nodes {
		assert.Equal(t, data[ans[i]], node.(*common.Person))
		result := avl.Search(ans[i])
		assert.Equal(t, data[ans[i]].Name, result.(*common.Person).Name)
	}

	removed := []int{9, 2, 8, 5}
	for _, target := range removed {
		avl.RemoveNode(target)
		assert.Equal(t, nil, avl.Search(target))
	}

	nodes = avl.GetInOrder()
	ans = []int{1, 3, 4, 6, 7, 12, 14, 23, 27}
	for i, node := range nodes {
		assert.Equal(t, data[ans[i]], node.(*common.Person))
		result := avl.Search(ans[i])
		assert.Equal(t, data[ans[i]].Name, result.(*common.Person).Name)
	}

	assert.Equal(t, true, avl.Validate(1, 27))
	assert.Equal(t, true, avl.Validate(1, 30))
	assert.Equal(t, false, avl.Validate(1, 26))
	assert.Equal(t, false, avl.Validate(4, 27))
}
