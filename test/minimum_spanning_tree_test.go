package test

import (
	"awesome-golang/minimum-spanning-tree/kruskal"
	"awesome-golang/minimum-spanning-tree/prim"
	"fmt"
	"testing"
)

func TestKruskal(t *testing.T) {
	graph := [][]int{}
	mst, ret := kruskal.GetMinimumCost(graph)
	fmt.Println(mst, ret)
}

func TestPrim(t *testing.T) {
	graph := [][]int{}
	mst, ret := prim.GetMinimumCost(graph)
	fmt.Println(mst, ret)
}
