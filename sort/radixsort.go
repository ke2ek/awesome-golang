package sort

/*
Radix Sort is a sorting algorithm designed to work on items
where the key of each item is an ordered set of integers in the range 0 to (N-1) inclusive both ends,
or can be transformed into such an ordered set.

Time Complexity: O(N*D)
where D is the maximum size of the bucket.
*/

func RadixSort(arr []int, maxSize int) {
	buckets := make([][]int, 10)
	exp := 1
	for i := 1; i <= maxSize; i++ {
		for _, num := range arr {
			key := int(num/exp) % 10
			buckets[key] = append(buckets[key], num)
		}
		exp *= 10
		j := 0
		for k := 0; k < 10; k++ {
			for _, num := range buckets[k] {
				arr[j] = num
				j++
			}
			// Finish to sort eariy if maxSize is greater than the number of digits
			// of the maximum number from arr.
			if len(buckets[k]) == len(arr) {
				return
			}
			buckets[k] = []int{}
		}
	}
}
