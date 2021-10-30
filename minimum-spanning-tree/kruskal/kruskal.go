package kruskal

import (
	unionFind "awesome-golang/union-find"
	"sort"
)

/*
Kruskal algorithm is a greedy algorithm in graph theory as, in each step, it adds the next lowest-weight edge
that won't form a cycle to the current Minimum Spanning Tree.

Time Complexity is O(E*log2(E)) where E is the number of edges in a graph.
*/

// Assume that graph[i][j] is the weight between i and j.
func GetMinimumCost(graph [][]int) ([][2]int, int) {
	// 1. Sort every edges in an increasing order of weights.
	edges := [][3]int{} // (cost, i, j)
	for i := 0; i < len(graph); i++ {
		for j := i + 1; j < len(graph); j++ {
			if graph[i][j] > 0 {
				edges = append(edges, [3]int{graph[i][j], i, j})
			}
		}
	}
	sort.Slice(edges[:], func(i, j int) bool {
		if edges[i][0] == edges[j][0] {
			if edges[i][1] == edges[j][1] {
				return edges[i][2] < edges[j][2]
			}
			return edges[i][1] < edges[j][1]
		}
		return edges[i][0] < edges[j][0]
	})
	// 2. Add an edge to the MST by checking with Disjoint Set if it forms a cycle.
	// 3. Keep adding edges until all of the vertices are contained in the MST.
	ret := 0
	mst := [][2]int{}
	ds := unionFind.NewDisjointSet(len(graph))
	for _, edge := range edges {
		cost, from, to := edge[0], edge[1], edge[2]
		u := ds.Find(from)
		v := ds.Find(to)
		if u != v {
			ds.Union(u, v)
			ret += cost
			mst = append(mst, [2]int{u, v})
		}
	}
	return mst, ret
}
