package test

import (
	mySort "awesome-golang/sort"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	N := int(1e5)
	arr, ans := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N)
		ans[i] = arr[i]
	}
	sort.Ints(ans)
	mySort.BubbleSort(&arr)
	for i := 0; i < N; i++ {
		assert.Equal(t, ans[i], arr[i])
	}
}

func TestQuickSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	N := int(1e5)
	arr, ans := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N)
		ans[i] = arr[i]
	}
	sort.Ints(ans)
	mySort.QuickSort(&arr, 0, N-1)
	for i := 0; i < N; i++ {
		assert.Equal(t, ans[i], arr[i])
	}
}
