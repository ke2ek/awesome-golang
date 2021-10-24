package unionfind

/*
A disjoint-set data structure is a data structure that keeps track of a set of elements partitioned into
a number of disjoint (non-overlapping) subsets.

A union-find algorithm is an algorithm that performs two useful operations on such a data structure

1. Find: Determine which subset a particular element is in.
	- This can be used for determining if two elements are in the same subset.
2. Union: Join two subsets into a single subset.
*/

type NaiveDisjointSet struct {
	parent []int
}

func NewNaiveDisjointSet(n int) *NaiveDisjointSet {
	ds := &NaiveDisjointSet{parent: make([]int, n)}
	// Each element put into each subset. i.e. subset[i] = i
	for i := 0; i < n; i++ {
		ds.parent[i] = i
	}
	return ds
}

// O(h), where h is the height of a tree.
// Return the number of root node in the tree, meaning the subset that the specific element(u) belongs to.
func (this *NaiveDisjointSet) Find(u int) int {
	if u == this.parent[u] {
		return u
	}
	return this.Find(this.parent[u])
}

// Merge one tree that u belongs to and another tree that v belongs to.
func (this *NaiveDisjointSet) Union(u, v int) {
	u = this.Find(u)
	v = this.Find(v)
	if u == v {
		return // They have been already in the same subset.
	}
	this.parent[u] = v // Notice that there can be a biased tree in one direction.
}
