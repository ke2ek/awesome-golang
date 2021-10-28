package unionfind

/*
A disjoint-set data structure is a data structure that keeps track of a set of elements partitioned into
a number of disjoint (non-overlapping) subsets.

A union-find algorithm is an algorithm that performs two useful operations on such a data structure

1. Find: Determine which subset a particular element is in.
	- This can be used for determining if two elements are in the same subset.
2. Union: Join two subsets into a single subset.

Here is the optimized solution.
*/

type DisjointSet struct {
	parent []int
	rank   []int
}

func NewDisjointSet(n int) *DisjointSet {
	ds := &DisjointSet{parent: make([]int, n), rank: make([]int, n)}
	// Each element put into each subset. i.e. subset[i] = i
	for i := 0; i < n; i++ {
		ds.parent[i] = i
	}
	return ds
}

// O(log(N)), Note that the path compression optimization has been applied.
// Return the number of root node in the tree, meaning the subset that the specific element(u) belongs to.
func (this *DisjointSet) Find(u int) int {
	if u == this.parent[u] {
		return u
	}
	this.parent[u] = this.Find(this.parent[u]) // duplicate operations will be removed.
	return this.parent[u]
}

// Merge one tree that u belongs to and another tree that v belongs to.
func (this *DisjointSet) Union(u, v int) {
	u = this.Find(u)
	v = this.Find(v)
	if u == v {
		return // They have been already in the same subset.
	}
	// Assume that the tree that `u` belongs to is lower height than the tree that `v` belongs to.
	// It will move a tree with the lower height into a subtree of the higher tree.
	if this.rank[u] > this.rank[v] {
		u, v = v, u
	}
	this.parent[u] = v
	if this.rank[u] == this.rank[v] {
		this.rank[v]++
	}
}

func (this *DisjointSet) Parent() []int {
	return this.parent
}

func (this *DisjointSet) Rank() []int {
	return this.rank
}
