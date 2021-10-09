package test

import (
	"awesome-golang/tree"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	No   int
	Name string
	Age  int
}

func makeData() map[int]*Person {
	return map[int]*Person{
		8:  &Person{No: 8, Name: "Jack", Age: 12},
		4:  &Person{No: 4, Name: "Dan", Age: 18},
		12: &Person{No: 12, Name: "Lisa", Age: 24},
		2:  &Person{No: 2, Name: "Kim", Age: 9},
		1:  &Person{No: 1, Name: "Lee", Age: 14},
		3:  &Person{No: 3, Name: "Park", Age: 16},
		5:  &Person{No: 5, Name: "Moana", Age: 23},
		6:  &Person{No: 6, Name: "Raven", Age: 35},
		7:  &Person{No: 7, Name: "Liu", Age: 28},
		9:  &Person{No: 9, Name: "Jinny", Age: 27},
		23: &Person{No: 23, Name: "Jane", Age: 22},
		14: &Person{No: 14, Name: "Jim", Age: 30},
		27: &Person{No: 27, Name: "Ken", Age: 29},
	}
}

func TestBinarySearchTree(t *testing.T) {
	data := makeData()
	root := &tree.BSTNode{Key: 8}
	for key, value := range data {
		root.Add(key, value)
	}

	ans := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 14, 23, 27}
	nodes := []*tree.BSTNode{}
	tree.TraverseInOrder(root, &nodes)
	for i, node := range nodes {
		assert.Equal(t, ans[i], node.Key)
		person := root.Search(ans[i]).(*Person)
		assert.Equal(t, data[ans[i]].Name, person.Name)
	}

	removed := []int{6, 27, 23, 4}
	for _, target := range removed {
		assert.Equal(t, true, root.Remove(target))
		assert.Equal(t, nil, root.Search(target))
	}
	/*
		nodes = []*tree.BSTNode{}
		ans = []int{1, 2, 3, 5, 7, 8, 9, 12, 14}
		tree.TraverseInOrder(root, &nodes)
		for i, node := range nodes {
			assert.Equal(t, node.Key, ans[i])
		}
	*/
}
