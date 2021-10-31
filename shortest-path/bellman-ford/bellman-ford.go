package bellmanford

import "awesome-golang/common"

/*
Bellman-Ford algorithm is to find the shortest path like Dijkstra's algorithm, but it can handle the graph with negative weights.
Also, using this algorithm, we can know if a given graph has any cycle with negative weights.

How it works: Reduce the difference between the upper bound of the shortest path and the actual cost/distance/etc
of the shortest path by predicting the proper upper bound.

What is relaxing: it means that we are gonna update upper[v] if upper[v] <= upper[u] + weight(u, v),
where u is the source vertex and v is the destination vertex.

Process:
1. the array upper[] contains each length of the actual shortest path.
2. Initialize upper[s] = 0, where s is the start vertex, and upper[v] = INF, where v is vertices other than s in the graph.
3. Approximate the path to the shortest path by relaxing along the edge (u, v).
	- It will success all the time regardless of negative-weights because the shortest path between u and v, satisfy:
		- dist[v] <= dist[u] + weight(u, v)
	- It works only in a weighted graph without a cycle.
4. Repeat relaxing on every edges at most (V-1) times, where V is the number of vertices.

To check if a given graph has a cycle with negative weights, it iwll repeat relaxing at most V times.
If relaxing fails at the V-th iteration, it happens to be the graph with a cycle.
(Notice that a cycle with negative weights reduces the toal cost of the shortest path whenever we calculate.)

Time Complexity is O(V*E) where V is the number of vertices and E is the number of edges in a graph.
*/

// Assume that graph is an adjacency list. e.g., graph[i][0] = (next vertex, the cost of the path)
// Return the shortest path from the source to each vertex. (the returned list has 0 at the position of source)
func GetShortestPath(graph [][][2]int, source int) []int {
	V := len(graph)
	upper := make([]int, V)
	for i := 0; i < V; i++ {
		upper[i] = common.INF
	}
	upper[source] = 0
	isRelaxing := false
	for iter := 0; iter < V; iter++ {
		isRelaxing = false
		for here := 0; here < V; here++ {
			for _, nbr := range graph[here] {
				there, cost := nbr[0], nbr[1]
				if upper[there] > upper[here]+cost {
					upper[there] = upper[here] + cost
					isRelaxing = true
				}
			}
		}
		// If relaxing fail on every edges, it doesn't have to run the remaining iterations.
		if !isRelaxing {
			break
		}
	}
	// If relaxing at the V-th iteration succeeded, it must have a cycle so it returns an empty list.
	if isRelaxing {
		upper = []int{}
	}
	return upper
}
