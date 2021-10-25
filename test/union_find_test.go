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
