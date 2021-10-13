package sort

// Return the index of pivot after partitioning
// This function is gonna divide the given array into two groups.
// The numbers coming to the left of the pivot are less than pivot.
// The numbers coming to the right of the pivot are greater than pivot.
func partition(arr *[]int, l, r int) int {
	pivot := (*arr)[r]
	i := l - 1
	for j := l; j < r; j++ {
		if (*arr)[j] < pivot {
			i++
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		}
	}
	i++
	(*arr)[i], (*arr)[r] = (*arr)[r], (*arr)[i]
	return i
}

/*
Time Complexity
Worst Case: [1,2,3,4,5,6,7,8,9] -> it has been already sorted.
			[2,2,2,2,2,2,2,2,2] -> it has had the same number.

QuickSort(arr, 0, 10) -> QuickSort(arr, 0, 9) -> ... -> QuickSort(arr, 0, 1)
Iteration in partition: N -> N-1 -> ... -> 1
Thus, it's gonna spend O(N^2) = N*(N+1)/2 = N + (N-1) + ... + 1

Best Case: Every time it choose the proper index of the pivot.
It will deal with the half of a given array every recursion to partition.

QuickSort(arr, 0, 10) -> QuickSort(arr, 0, 5) -> ... -> QuickSort(arr, 0, 1)
Iteration in partition: N -> N/2 * 2 = N -> N/4 * 4 = N -> ... -> 1 * N = N
Thus, it's gonna spend O(NlogN) because the number of recursive calls is gonna be logN.
*/
func QuickSort(arr *[]int, l, r int) {
	// base case
	if l >= r {
		return
	}
	// recursive case
	p := partition(arr, l, r)
	QuickSort(arr, l, p-1)
	QuickSort(arr, p+1, r)
}
