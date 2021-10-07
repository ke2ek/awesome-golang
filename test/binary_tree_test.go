package test

import (
	"awesome-golang/common"
	"awesome-golang/tree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeTree() *tree.Node {
	node64 := tree.NewNode(64, nil, nil)
	node1 := tree.NewNode(1, nil, nil)
	node56 := tree.NewNode(56, nil, nil)
	node102 := tree.NewNode(102, nil, nil)

	node3 := tree.NewNode(3, node64, node1)
	node7 := tree.NewNode(7, nil, node56)
	node12 := tree.NewNode(12, nil, node102)
	node86 := tree.NewNode(86, nil, nil)

	node4 := tree.NewNode(4, node3, node7)
	node8 := tree.NewNode(8, node86, node12)

	node5 := tree.NewNode(5, node4, node8)

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
	nodes := make([][]*tree.Node, 3)
	root.PreOrder(&nodes[0])
	root.InOrder(&nodes[1])
	root.PostOrder(&nodes[2])
	for i := 0; i < 3; i++ {
		for j, num := range ans[i] {
			assert.Equal(t, nodes[i][j].Value.(int), num, common.ERROR_MSG)
		}
	}

	nodes = [][]*tree.Node{}
	root.BFS(&nodes)
	i, j := 0, 0
	for i < len(ans[3]) && j < len(nodes) {
		for _, node := range nodes[j] {
			assert.Equal(t, node.Value.(int), ans[3][i], common.ERROR_MSG)
			i++
		}
		j++
	}
}
