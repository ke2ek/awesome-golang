package sort

/*
Shellsort is an optimization of insertion sort that allows the exchange of items that are far apart.
The method starts by sorting pairs of elements far apart from each other,
then progressively reducing the gap between elements to be compared.

By starting with far apart elements, it can move some out-of-place elements into position faster than a simple nearest neighbor exchange.

Time Complexity: O(N*log(N)^2)
*/

func ShellSort(arr []int) {
	for L := len(arr) / 2; L > 0; L /= 2 {
		for start := L; start < len(arr); start++ {
			hold := arr[start]
			current := start - L
			for current >= 0 && arr[current] > hold {
				arr[current+L] = arr[current]
				current -= L
			}
			arr[current+L] = hold
		}
	}
}
