package test

import (
	unionFind "awesome-golang/union-find"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnionFind(t *testing.T) {
	naiveDs := unionFind.NewNaiveDisjointSet(4)
	naiveDs.Union(0, 1)
	naiveDs.Union(1, 2)
	naiveDs.Union(2, 3)
	// Test that it is a biased tree in one direction.
	for i := 0; i < 4; i++ {
		assert.Equal(t, 3, naiveDs.Find(i))
	}

	ds := unionFind.NewDisjointSet(4)
	ds.Union(0, 1)
	ds.Union(1, 2)
	ds.Union(2, 3)
	// Test that it is a balanced tree.
	for i := 0; i < 4; i++ {
		assert.Equal(t, 1, ds.Find(i))
	}

	/*
		Test that it can be the compressed tree.
		initial data:
			rank = [0,0,0,0,0,0,0,0,0,0,0,0,0]
			parent = [0,1,2,3,4,5,6,7,8,9,10,11,12]
	*/

	ds2 := unionFind.NewDisjointSet(13)
	ds2.Union(1, 0) //  parent[1] = 0, rank[0] = 1
	// rank = [1,0,0,0,0,0,0,0,0,0,0,0,0]
	// parent = [0,0,2,3,4,5,6,7,8,9,10,11,12]

	ds2.Union(2, 0) // parent[2] = 0, rank[0] = 1
	// rank = [1,0,0,0,0,0,0,0,0,0,0,0,0]
	// parent = [0,0,0,3,4,5,6,7,8,9,10,11,12]

	ds2.Union(3, 1) // parent[3] = 0, rank[0] = 1
	// rank = [1,0,0,0,0,0,0,0,0,0,0,0,0]
	// parent = [0,0,0,0,4,5,6,7,8,9,10,11,12]

	ds2.Union(5, 4) // parent[5] = 4, rank[4] = 1
	// rank = [1,0,0,0,1,0,0,0,0,0,0,0,0]
	// parent = [0,0,0,0,4,4,6,7,8,9,10,11,12]

	ds2.Union(6, 4) // parent[6] = 4, rank[4] = 1
	// rank = [1,0,0,0,1,0,0,0,0,0,0,0,0]
	// parent = [0,0,0,0,4,4,4,7,8,9,10,11,12]

	ds2.Union(4, 1) // parent[4] = 0, rank[0] = 2
	// rank = [2,0,0,0,1,0,0,0,0,0,0,0,0]
	// parent = [0,0,0,0,0,4,4,7,8,9,10,11,12]

	ds2.Union(10, 9) // parent[10] = 9, rank[9] = 1
	// rank = [2,0,0,0,1,0,0,0,0,1,0,0,0]
	// parent = [0,0,0,0,0,4,4,7,8,9,9,11,12]

	ds2.Union(11, 9) // parent[11] = 9, rank[9] = 1
	// rank = [2,0,0,0,1,0,0,0,0,1,0,0,0]
	// parent = [0,0,0,0,0,4,4,7,8,9,9,9,12]

	ds2.Union(12, 8) // parent[12] = 8, rank[8] = 1
	// rank = [2,0,0,0,1,0,0,0,1,1,0,0,0]
	// parent = [0,0,0,0,0,4,4,7,8,9,9,9,8]

	ds2.Union(9, 7) // parent[7] = 9, rank[9] = 1
	// rank = [2,0,0,0,1,0,0,0,1,1,0,1,0]
	// parent = [0,0,0,0,0,4,4,9,8,9,9,9,8]

	ds2.Union(12, 7) // parent[8] = 9, rank[9] = 2
	// rank = [2,0,0,0,1,0,0,0,1,2,0,1,0]
	// parent = [0,0,0,0,0,4,4,9,9,9,9,9,8]

	ds2.Union(7, 0) // parent[9] = 0, rank[0] = 3
	// rank = [3,0,0,0,1,0,0,0,1,2,0,1,0]
	// parent = [0,0,0,0,0,4,4,9,9,0,9,9,8]

	parent := []int{0, 0, 0, 0, 0, 4, 4, 9, 9, 0, 9, 9, 8}
	rank := []int{3, 0, 0, 0, 1, 0, 0, 0, 1, 2, 0, 0, 0}
	assert.Equal(t, parent, ds2.Parent())
	assert.Equal(t, rank, ds2.Rank())
}

func TestDetectCycle(t *testing.T) {
	graph1 := [][]int{
		{1, 2}, {0, 2}, {0, 1},
	}
	graph2 := [][]int{
		{1, 2}, {0}, {0}, // no-cycle
	}
	graph3 := [][]int{
		{1, 2}, {0, 2}, {0, 1},
		{4, 5}, {3, 5}, {3, 4},
		{7, 8}, {6, 8}, {6, 7},
	}
	assert.Equal(t, true, unionFind.HasCycle(graph1))
	assert.Equal(t, false, unionFind.HasCycle(graph2))
	assert.Equal(t, true, unionFind.HasCycle(graph3))
	assert.Equal(t, 1, unionFind.CountCycle(graph1))
	assert.Equal(t, 0, unionFind.CountCycle(graph2))
	assert.Equal(t, 3, unionFind.CountCycle(graph3))
}
