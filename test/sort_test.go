package test

import (
	mySort "awesome-golang/sort"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const N = 10000

func TestBubbleSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	arr, ans := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N * 100)
		ans[i] = arr[i]
	}
	sort.Ints(ans)
	mySort.BubbleSort(&arr)
	for i := 0; i < N; i++ {
		assert.Equal(t, ans[i], arr[i])
	}
}

func TestInsertionSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	arr, ans := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N * 100)
		ans[i] = arr[i]
	}
	sort.Ints(ans)
	mySort.InsertionSort(&arr)
	for i := 0; i < N; i++ {
		assert.Equal(t, ans[i], arr[i])
	}
}

func TestQuickSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	arr, ans := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N * 100)
		ans[i] = arr[i]
	}
	sort.Ints(ans)
	mySort.QuickSort(&arr, 0, N-1)
	for i := 0; i < N; i++ {
		assert.Equal(t, ans[i], arr[i])
	}
}

func TestMergeSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	arr, ans := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N * 100)
		ans[i] = arr[i]
	}
	sort.Ints(ans)
	arr = mySort.MergeSort(&arr, 0, N-1)
	for i := 0; i < N; i++ {
		assert.Equal(t, ans[i], arr[i])
	}
}

func TestShellSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	arr, ans := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N * 100)
		ans[i] = arr[i]
	}
	sort.Ints(ans)
	mySort.ShellSort(&arr)
	for i := 0; i < N; i++ {
		assert.Equal(t, ans[i], arr[i])
	}
}
