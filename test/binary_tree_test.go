package test

import (
	"awesome-golang/common"
	"awesome-golang/tree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeTree() *tree.BTNode {
	node64 := tree.BTNode{Value: 64}
	node1 := tree.BTNode{Value: 1}
	node56 := tree.BTNode{Value: 56}
	node102 := tree.BTNode{Value: 102}

	node3 := tree.BTNode{Value: 3, Left: &node64, Right: &node1}
	node7 := tree.BTNode{Value: 7, Right: &node56}
	node12 := tree.BTNode{Value: 12, Right: &node102}
	node86 := tree.BTNode{Value: 86}

	node4 := tree.BTNode{Value: 4, Left: &node3, Right: &node7}
	node8 := tree.BTNode{Value: 8, Left: &node86, Right: &node12}

	node5 := tree.BTNode{Value: 5, Left: &node4, Right: &node8}

	return &node5
}

func TestBinaryTree(t *testing.T) {
	root := makeTree()
	ans := [][]int{
		{5, 4, 3, 64, 1, 7, 56, 8, 86, 12, 102}, // pre-order
		{64, 3, 1, 4, 7, 56, 5, 86, 8, 12, 102}, // in-order
		{64, 1, 3, 56, 7, 4, 86, 102, 12, 8, 5}, // post-order
		{5, 4, 8, 3, 7, 86, 12, 64, 1, 56, 102}, // bfs
	}
	nodes := make([][]*tree.BTNode, 3)
	root.PreOrder(&nodes[0])
	root.InOrder(&nodes[1])
	root.PostOrder(&nodes[2])
	for i := 0; i < 3; i++ {
		for j, num := range ans[i] {
			assert.Equal(t, nodes[i][j].Value.(int), num, common.ERROR_MSG)
		}
	}

	nodes = [][]*tree.BTNode{}
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
