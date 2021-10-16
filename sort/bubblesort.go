package sort

func BubbleSort(arr []int) {
	N := len(arr)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
