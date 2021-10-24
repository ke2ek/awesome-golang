package unionfind

// Detect Cycle in an Undirected Graph.
// Assume that a graph will be given an adjacent list.
func HasCycle(graph [][]int) bool {
	ds := NewDisjointSet(len(graph))
	for u := 0; u < len(graph); u++ {
		for _, v := range graph[u] {
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
			setU := ds.Find(u)
			setV := ds.Find(v)
			if setU != setV {
				ds.Union(setU, setV)
			}
		}
	}
	for u := 0; u < len(graph); u++ {
		for _, v := range graph[u] {
			setU := ds.Find(u)
			setV := ds.Find(v)
			if setU != setV {
				numCycles++
			}
		}
	}
	return numCycles
}
