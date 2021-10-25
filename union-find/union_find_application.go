package unionfind

// Detect Cycle in an Undirected Graph.
// Assume that a graph will be given an adjacent list.
func HasCycle(graph [][]int) bool {
	ds := NewDisjointSet(len(graph))
	for u := 0; u < len(graph); u++ {
		for _, v := range graph[u] {
			if v < u {
				// Edges that have been already merged will be excluded. (only if a graph is an undirected graph)
				continue
			}
			setU := ds.Find(u)
			setV := ds.Find(v)
			if setU == setV {
				return true
			}
			ds.Union(setU, setV)
		}
	}
	return false
}

func CountCycle(graph [][]int) int {
	numCycles := 0
	ds := NewDisjointSet(len(graph))
	for u := 0; u < len(graph); u++ {
		for _, v := range graph[u] {
			if v < u {
				// Edges that have been already merged will be excluded. (only if a graph is an undirected graph)
				continue
			}
			setU := ds.Find(u)
			setV := ds.Find(v)
			if setU == setV {
				numCycles++
				continue
			}
			ds.Union(setU, setV)
		}
	}
	return numCycles
}
