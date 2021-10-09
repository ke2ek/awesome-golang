package test

import (
	"awesome-golang/skiplist"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkiplist(t *testing.T) {
	sl := skiplist.New(5, 0.5)
	data := []int{10, 4, 6, 8, 19, 395, 432, 300, 9, 32,
		17, 28, 35, 88, 53, 36, 13, 45, 47, 66,
		33, 25, 90, 18, 7, 66, 2, 300, 88, 245}
	// total count : 30
	for _, num := range data {
		sl.Add(num)
		sl.Print()
	}

	removed := []int{245, 18, 432, 36, 13}
	for _, num := range removed {
		assert.Equal(t, true, sl.Search(num))
		assert.Equal(t, true, sl.Erase(num))
		sl.Print()
		assert.Equal(t, false, sl.Search(num))
	}

	removed = []int{300, 88, 66}
	for _, num := range removed {
		assert.Equal(t, true, sl.Search(num))
		assert.Equal(t, true, sl.Erase(num))
		assert.Equal(t, true, sl.Search(num))
		assert.Equal(t, true, sl.Erase(num))
		assert.Equal(t, false, sl.Search(num))
	}
}
