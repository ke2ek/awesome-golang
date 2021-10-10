package test

import (
	"awesome-golang/sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuicksort(t *testing.T) {
	arr := []int{4, 5, 7, 8, 9, 2, 1, 6, 3}
	sort.Quicksort(&arr, 0, len(arr)-1)
	for i := 1; i < 10; i++ {
		assert.Equal(t, i, arr[i-1])
	}
}
