package prim

/*
Prim algorithm is a greedy algorithm in graph theory as, in each step, it adds the vertex that is the closest
to the current Minimum Spanning Tree and does not form a cycle.

It starts from an arbitrary vertex at first.

Time Complexity is O(V^2 + E) where V is the number of vertices and E is the number of edges in a graph.
*/

// Assume that graph[i][j] is the weight between i and j.
func GetMinimumCost(graph [][]int) ([][2]int, int) {
	INF := 987654321
	V := len(graph)
	isAdded := make([]bool, V)  // It will set to true if i-th vertex is added to the MST.
	minWeight := make([]int, V) // It will set to the minimum weight from i-th vertex.
	opposite := make([]int, V)  // It will set to the opposite of i-th vertex on the edge with minWeight[i].
	// 1. Initialize minWeight[0] = 0, opposite[0] = 0, isAdded[i] = false
	for i := 1; i < V; i++ {
		minWeight[i], opposite[i] = INF, -1
	}
	ret := 0
	mst := [][2]int{}
	for iter := 0; iter < V; iter++ {
		// 2. Find the next vertex which is not still contained in the current MST.
		next := -1
		for v := 0; v < V; v++ {
			if isAdded[v] {
				continue
			}
			if next != -1 && minWeight[next] < minWeight[v] {
				continue
			}
			next = v
		}
		// if the next vertex is the first vertex to be added, skip.
		if opposite[next] != next {
			mst = append(mst, [2]int{next, opposite[next]})
		}
		// 3. Add a new edge to the MST if the edge is not included.
		ret += minWeight[next]
		isAdded[next] = true
		// 4. Update minWeight and visit adjacency edges of next by choosing the edge with the minimum weight.
		for v := 0; v < V; v++ {
			if graph[next][v] > 0 && !isAdded[v] {
				if minWeight[v] > graph[next][v] {
					opposite[v] = next
					minWeight[v] = graph[next][v]
				}
			}
		}
	}
	return mst, ret
}
