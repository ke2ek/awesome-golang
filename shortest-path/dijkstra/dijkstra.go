package dijkstra

import (
	"awesome-golang/common"
	priorityQueue "awesome-golang/priority-queue"
)

/*
Dijkstra's algorithm is to find the shortest path by discovering the next closest vertex from the current vertex in each step.
It is not finding neighbors unlike BFS. But it is also using Priority Queue.
Basically, a priority queue sorts elements in a decreasing order.
So, we are gonna multiply the element to be inserted by -1 to get the minimum value.

Time Complexity is O(E * log(V)) where E is the number of edges and V is the number of nodes.
	- The number of vertices to be added into the queue is at most O(|E|).
	- It will cost O(log(E)) to insert or delete an element into the queue.
	- In gnenral, the number of edges is smaller than V^2 so O(log(E)) <= O(log(V^2)) = O(2*log(V)) = O(log(V)).
*/

// Assume that graph is an adjacency list. e.g., graph[i][0] = (next vertex, the cost of the path)
// Return the shortest path from the source to each vertex. (the returned list has 0 at the position of source)
func GetShortestPath(graph [][][2]int, source int) []int {
	V := len(graph)
	dist := make([]int, V)
	for i := 0; i < V; i++ {
		dist[i] = common.INF
	}
	dist[source] = 0
	pq := priorityQueue.New()
	pq.Push(0, source) // (length fo path, vertex)
	for !pq.Empty() {
		node := pq.Pop()
		cost := node.Key
		here := node.Value.(int)

		// If the current vertex has been already chosen by the shorter distance, ignore.
		if dist[here] < cost {
			continue
		}

		// Visit all the neighbors.
		for _, nbr := range graph[here] {
			there := nbr[0]
			nextDist := cost + nbr[1]
			if dist[there] > nextDist {
				dist[there] = nextDist
				pq.Push(-nextDist, there)
			}
		}
	}
	return dist
}
