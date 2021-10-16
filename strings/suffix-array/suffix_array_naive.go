package suffixarray

import (
	"sort"
)

/*
Naive Method
1. Get the start index of each suffix.
2. Sort by suffix, not index.
*/

type SuffixArrayNaive struct {
	s   string
	arr []int
}

func (this *SuffixArrayNaive) Len() int {
	return len(this.arr)
}

func (this *SuffixArrayNaive) Less(i, j int) bool {
	return this.s[this.arr[i]:] < this.s[this.arr[j]:]
}

func (this *SuffixArrayNaive) Swap(i, j int) {
	this.arr[i], this.arr[j] = this.arr[j], this.arr[i]
}

func (this *SuffixArrayNaive) Array() []int {
	return this.arr
}

// O(N^2 * log(N))
func NewSuffixArrayNaive(s string) *SuffixArrayNaive {
	sfa := &SuffixArrayNaive{s: s, arr: make([]int, len(s))}
	for i := 0; i < sfa.Len(); i++ {
		sfa.arr[i] = i
	}
	sort.Sort(sfa)
	return sfa
}
