package test

import (
	bellmanFord "awesome-golang/shortest-path/bellman-ford"
	"awesome-golang/shortest-path/dijkstra"
	floydWarshall "awesome-golang/shortest-path/floyd-warshall"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	undiretedGraph = [][][2]int{
		{{1, 4}, {7, 8}},
		{{0, 4}, {7, 11}, {2, 8}},
		{{1, 8}, {3, 7}, {5, 4}, {8, 2}},
		{{2, 7}, {4, 9}, {5, 14}},
		{{3, 9}, {5, 10}},
		{{2, 4}, {3, 14}, {4, 10}, {6, 2}},
		{{5, 2}, {7, 1}, {8, 6}},
		{{0, 8}, {1, 11}, {6, 1}, {8, 7}},
		{{2, 2}, {6, 6}, {7, 7}},
	}
	directedGraph = [][][2]int{
		{{1, -1}, {2, 4}},
		{{2, 3}, {3, 2}, {4, 2}},
		{},
		{{1, 1}, {2, 5}},
		{{3, -3}},
	}

	undiretedGraphAnswer = []int{0, 4, 12, 19, 21, 11, 9, 8, 14}
	directedGraphAnswer  = []int{0, -1, 2, -2, 1}
)

func TestDijkstra(t *testing.T) {
	result := dijkstra.GetShortestPath(undiretedGraph, 0)
	assert.Equal(t, undiretedGraphAnswer, result)
}

func TestBellmanFord(t *testing.T) {
	result := bellmanFord.GetShortestPath(undiretedGraph, 0)
	assert.Equal(t, undiretedGraphAnswer, result)

	negCycleGraph := [][][2]int{
		{{1, -4}, {2, 8}},
		{{0, -4}, {2, 11}},
		{{0, 8}, {1, 11}},
	}
	result2 := bellmanFord.GetShortestPath(negCycleGraph, 0)
	assert.Equal(t, []int{}, result2)

	result3 := bellmanFord.GetShortestPath(directedGraph, 0)
	assert.Equal(t, []int{0, -1, 2, -2, 1}, result3)
}

func TestFloydWarshall(t *testing.T) {
	V1 := len(undiretedGraph)
	result := floydWarshall.GetShortestPathNaive(undiretedGraph)
	matrix, _ := floydWarshall.GetShortestPath(undiretedGraph)
	for i := 0; i < V1; i++ {
		ans1 := dijkstra.GetShortestPath(undiretedGraph, i)
		assert.Equal(t, ans1, result[V1-1][i])
		assert.Equal(t, ans1, matrix[i])

		ans2 := bellmanFord.GetShortestPath(undiretedGraph, i)
		assert.Equal(t, ans2, result[V1-1][i])
		assert.Equal(t, ans2, matrix[i])
	}

	graph := [][][2]int{
		{{1, 5}, {3, 10}},
		{{2, 3}},
		{{3, 1}},
		{},
	}
	V2 := len(graph)
	result2 := floydWarshall.GetShortestPathNaive(graph)
	matrix2, via2 := floydWarshall.GetShortestPath(graph)
	for i := 0; i < V2; i++ {
		ans1 := dijkstra.GetShortestPath(graph, i)
		assert.Equal(t, ans1, result2[V2-1][i])
		assert.Equal(t, ans1, matrix2[i])

		ans2 := bellmanFord.GetShortestPath(graph, i)
		assert.Equal(t, ans2, result2[V2-1][i])
		assert.Equal(t, ans2, matrix2[i])
	}
	path := []int{}
	floydWarshall.Reconstruct(0, V2-1, matrix2, via2, &path)
	assert.Equal(t, []int{0, 1, 2, 3}, path)
}
