package test

import (
	tree "awesome-golang/binary-tree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeTree() *tree.BTNode {
	node64 := tree.NewBTNode(64, nil, nil)
	node1 := tree.NewBTNode(1, nil, nil)
	node56 := tree.NewBTNode(56, nil, nil)
	node102 := tree.NewBTNode(102, nil, nil)

	node3 := tree.NewBTNode(3, node64, node1)
	node7 := tree.NewBTNode(7, nil, node56)
	node86 := tree.NewBTNode(86, nil, nil)
	node12 := tree.NewBTNode(12, nil, node102)

	node4 := tree.NewBTNode(4, node3, node7)
	node8 := tree.NewBTNode(8, node86, node12)

	node5 := tree.NewBTNode(5, node4, node8)

	return node5
}

func TestBinaryTree(t *testing.T) {
	root := makeTree()
	ans := [][]int{
		{5, 4, 3, 64, 1, 7, 56, 8, 86, 12, 102}, // pre-order
		{64, 3, 1, 4, 7, 56, 5, 86, 8, 12, 102}, // in-order
		{64, 1, 3, 56, 7, 4, 86, 102, 12, 8, 5}, // post-order
		{5, 4, 8, 3, 7, 86, 12, 64, 1, 56, 102}, // bfs
	}
	nodes := make([][]interface{}, 3)
	tree.PreOrder(root, &nodes[0])
	tree.InOrder(root, &nodes[1])
	tree.PostOrder(root, &nodes[2])
	for i := 0; i < 3; i++ {
		for j, num := range ans[i] {
			assert.Equal(t, num, nodes[i][j].(int))
		}
	}

	result := *(tree.BFS(root))
	i, j := 0, 0
	for i < len(ans[3]) && j < len(result) {
		for _, value := range result[j] {
			assert.Equal(t, ans[3][i], value.(int))
			i++
		}
		j++
	}
}
