package sort

func MergeSort(arr *[]int, l, r int) []int {
	// base case: arr has been already sorted because there is only one element.
	if l == r {
		return []int{(*arr)[l]}
	}
	m := (l + r) / 2
	left := MergeSort(arr, l, m)
	right := MergeSort(arr, m+1, r)
	sorted := make([]int, r-l+1)
	for i, j, k := 0, 0, 0; i < len(left) || j < len(right); k++ {
		if j == len(right) {
			sorted[k] = left[i]
			i++
		} else if i == len(left) {
			sorted[k] = right[j]
			j++
		} else {
			if left[i] < right[j] {
				sorted[k] = left[i]
				i++
			} else {
				sorted[k] = right[j]
				j++
			}
		}
	}
	return sorted
}
