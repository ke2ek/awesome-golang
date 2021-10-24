package test

import (
	"awesome-golang/common"
	segmentTree "awesome-golang/segment-tree"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRangeMinimumQuery(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	N := 100
	arr := make([]int, N)
	ans := int(1e9)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N * 100)
		ans = common.Min(ans, arr[i])
	}
	rmq := segmentTree.NewRangeMinimumQuery(arr)
	assert.Equal(t, ans, rmq.Query(0, N-1))
	rmq.Update(0, -99)
	assert.Equal(t, -99, rmq.Query(0, N-1))
	rmq.Update(N/2, -488)
	assert.Equal(t, -488, rmq.Query(0, N-1))
	rmq.Update(N/4, -999)
	assert.Equal(t, -999, rmq.Query(0, N-1))
	assert.Equal(t, -488, rmq.Query(N/2, N-1))
}

func TestRangeSumQuery(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	N := 10000
	arr := make([]int, N)
	ans := 0
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N * 100)
		ans += arr[i]
	}
	rsq := segmentTree.NewRangeSumQuery(arr)
	assert.Equal(t, ans, rsq.Query(0, N-1))
	for i := 0; i < 1000; i++ {
		index := rand.Intn(N)
		value := rand.Intn(N * 100)
		prevSum := rsq.Query(0, index-1)
		rsq.Update(index, value)
		ans += value - arr[index]
		arr[index] = value
		assert.Equal(t, ans, rsq.Query(0, N-1))
		assert.Equal(t, prevSum, rsq.Query(0, index-1))
	}
}

func TestFenwickTree(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	N := 10
	tree := segmentTree.NewFenwickTree(N)
	arr := make([]int, N)
	ans := make([]int, N) // psum
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N * 100)
		if i == 0 {
			ans[i] = arr[i]
		} else {
			ans[i] = ans[i-1] + arr[i]
		}
		tree.Add(i, arr[i])
	}

	for i, psum := range ans {
		assert.Equal(t, psum, tree.Psum(i))
	}

	for i := 0; i < 1000; i++ {
		index := rand.Intn(N)
		value := rand.Intn(N * 100)
		psum1, psum2 := tree.Psum(index-1), tree.Psum(index)
		diff := value - arr[index]
		tree.Add(index, diff)
		arr[index] = value
		assert.Equal(t, psum1, tree.Psum(index-1))
		assert.Equal(t, psum2+diff, tree.Psum(index))
	}
}
