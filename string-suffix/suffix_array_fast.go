package stringsuffix

import (
	"sort"
)

/*
Manber-Myers Algorithm

Requirements:
	group[] := group number that each suffix belongs to. (where index means the start index of each suffix)
	arr[] := the start index of each suffix (by lexicographical order) = suffix array
	t := the number of first letters to be compared.

1. Set group[] as the first letter of each suffix and arr[] as the start index of each suffix.
	1.1. group[i] = str[i]
	1.2. arr[i] = i
	1.3. t = 1
2. whenever t increases, it will compare two suffixes as like below:
	2.1. If both suffixes belong to the same group, consider t letters.
	2.2. If both suffixes belong to different groups, compare only group numbers
		which represents the group of only the first t letters.
	2.3. after sorting, update group[] as like below:
		update group[arr[i]] as group[arr[i-1]]+1 if arr[i-1] is in front of arr[i] in the current group state.
		otherwise, update group[arr[i]] as group[arr[i-1]]
	2.4. t will be multiplied by 2.

The important thing in this algorithm is that
	- it reuses the lexicographical order information updated in previous steps by using group[].
	- it will iterate log(N) times in the first iteration because t continue to increase by 2 times until t equals to the length,
		- where N is the length of a given string.
	- In that iteration, sort algorithm will spend O(NlogN) time.
		- Totally, the time complexity is gonna be O(N * log(N)^2)
*/

type SuffixArray struct {
	s     string
	arr   []int
	group []int
	t     int
}

func (this *SuffixArray) Len() int {
	return len(this.arr)
}

func (this *SuffixArray) Less(i, j int) bool {
	if this.group[this.arr[i]] != this.group[this.arr[j]] {
		return this.group[this.arr[i]] < this.group[this.arr[j]]
	}
	return this.group[this.arr[i]+this.t] < this.group[this.arr[j]+this.t]
}

func (this *SuffixArray) Swap(i, j int) {
	this.arr[i], this.arr[j] = this.arr[j], this.arr[i]
}

func (this *SuffixArray) Array() []int {
	return this.arr
}

// O(N * (log(N))^2)
func NewSuffixArray(s string) *SuffixArray {
	sfa := &SuffixArray{
		s:     s,
		arr:   make([]int, len(s)),
		group: make([]int, len(s)*2+1),
		t:     1,
	}
	for i := 0; i < sfa.Len(); i++ {
		sfa.arr[i] = i
		sfa.group[i] = int(s[i])
	}

	for sfa.t < len(s) {
		sort.Sort(sfa)
		updateGroup := make([]int, len(sfa.group))
		updateGroup[sfa.arr[0]] = 1
		for i := 1; i < len(s); i++ {
			updateGroup[sfa.arr[i]] = updateGroup[sfa.arr[i-1]]
			if sfa.Less(i-1, i) {
				updateGroup[sfa.arr[i]]++
			}
		}
		sfa.group = updateGroup
		sfa.t *= 2
	}

	return sfa
}
