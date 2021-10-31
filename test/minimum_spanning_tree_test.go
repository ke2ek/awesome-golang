package test

import (
	"awesome-golang/minimum-spanning-tree/kruskal"
	"awesome-golang/minimum-spanning-tree/prim"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testcase = [][]int{
		{0, 3, 7, 0, 0, 0, 0, 0},
		{3, 0, 12, 9, 0, 0, 0, 0},
		{7, 12, 0, 10, 0, 0, 0, 0},
		{0, 9, 10, 0, 8, 9, 4, 0},
		{0, 0, 0, 8, 0, 0, 0, 2},
		{0, 0, 0, 9, 0, 0, 0, 6},
		{0, 0, 0, 4, 0, 0, 0, 5},
		{0, 0, 0, 0, 2, 6, 5, 0},
	}
	ans           = 36
	kruskalAnswer = [][2]int{{4, 7}, {0, 1}, {3, 6}, {6, 7}, {5, 7}, {1, 2}, {1, 7}}
	primAnswer    = [][2]int{{1, 0}, {2, 0}, {3, 1}, {6, 3}, {7, 6}, {4, 7}, {5, 7}}
)

func TestKruskal(t *testing.T) {
	mst, minCost := kruskal.GetMinimumCost(testcase)
	assert.Equal(t, ans, minCost)
	assert.Equal(t, kruskalAnswer, mst)
}

func TestPrim(t *testing.T) {
	mst, minCost := prim.GetMinimumCost(testcase)
	assert.Equal(t, ans, minCost)
	assert.Equal(t, primAnswer, mst)
}
