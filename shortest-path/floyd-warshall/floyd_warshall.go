package floydwarshall

import (
	"awesome-golang/common"
)

/*
Floyd-Warshall algorithm is to find the shortest path for all pairs of vertices in an weighted graph withour any cycle
by discovering the shortest path between each neighbor and the current vertex.

How it works:
Assume that DistS(u, v) is the shortest distance between u and v in the set of verteices, S, and
DistS-x(u, v) is the shortest distance between u and v, not via the vertex, x.

We can know that: DistS(u, v) = min(DistS(u, x) + DistS(x, v), DistS-x(u, v))

Notice that every edges in the shortest path should satisfy above expression.
*/

func convertAdjacencyToMatrix(graph [][][2]int) [][]int {
	V := len(graph)
	matrix := make([][]int, V)
	for i := 0; i < V; i++ {
		matrix[i] = make([]int, V)
		for j := 0; j < V; j++ {
			matrix[i][j] = common.INF
		}
	}
	for u := 0; u < V; u++ {
		matrix[u][u] = 0
		for _, nbr := range graph[u] {
			v, cost := nbr[0], nbr[1]
			matrix[u][v] = cost
		}
	}
	return matrix
}

// O(V^3) where V is the number of vertices.
// Assume that graph is an adjacency list. e.g., graph[i][0] = (next vertex, the cost of the path)
func GetShortestPathNaive(graph [][][2]int) [][][]int {
	V := len(graph)
	// Initialize
	dist := make([][][]int, V)
	matrix := convertAdjacencyToMatrix(graph)
	for i := 0; i < V; i++ {
		dist[i] = make([][]int, V)
		for j := 0; j < V; j++ {
			dist[i][j] = make([]int, V)
		}
	}

	for i := 0; i < V; i++ {
		for j := 0; j < V; j++ {
			if i != j {
				dist[0][i][j] = common.Min(matrix[i][j], matrix[i][0]+matrix[0][j])
			}
		}
	}

	// Calculate the shortest path.
	for k := 1; k < V; k++ {
		for i := 0; i < V; i++ {
			for j := 0; j < V; j++ {
				dist[k][i][j] = common.Min(dist[k-1][i][j], dist[k-1][i][k]+dist[k-1][k][j])
			}
		}
	}
	return dist
}

// O(V^2) where V is the number of vertices.
// Advanced Implementation with Sliding Window in DP.
func GetShortestPath(graph [][][2]int) (matrix, via [][]int) {
	V := len(graph)
	matrix = convertAdjacencyToMatrix(graph)
	// It is used to reconstruct the shortest path.
	via = make([][]int, V)
	for i := 0; i < V; i++ {
		via[i] = make([]int, V)
		for j := 0; j < V; j++ {
			via[i][j] = -1
		}
	}

	// Calculate the shortest path.
	for k := 0; k < V; k++ {
		for u := 0; u < V; u++ {
			if matrix[u][k] == common.INF {
				continue
			}
			for v := 0; v < V; v++ {
				//matrix[u][v] = common.Min(matrix[u][v], matrix[u][k]+matrix[k][v])
				if matrix[u][v] > matrix[u][k]+matrix[k][v] {
					via[u][v] = k
					matrix[u][v] = matrix[u][k] + matrix[k][v]
				}
			}
		}
	}

	return
}

// Reconstruct the shortest path from u to v.
func Reconstruct(u, v int, matrix, via [][]int, path *[]int) {
	if via[u][v] == -1 {
		*path = append(*path, u)
		if u != v {
			*path = append(*path, v)
		}
	} else {
		k := via[u][v]
		Reconstruct(u, k, matrix, via, path)
		if len(*path) > 0 {
			*path = (*path)[:len(*path)-1] // Prevent duplicating k.
		}
		Reconstruct(k, v, matrix, via, path)
	}
}
