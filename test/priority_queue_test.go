package test

import (
	priorityqueue "awesome-golang/priority-queue"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Student struct {
	sno   int
	name  string
	age   int
	major string
}

func TestPriorityQueue(t *testing.T) {
	pq := priorityqueue.New()
	data := map[string]Student{
		"Park": Student{sno: 1234, name: "Park", age: 21, major: "CS"},
		"Kim":  Student{sno: 456, name: "Kim", age: 22, major: "Economics"},
		"Kwon": Student{sno: 42, name: "Kwon", age: 24, major: "CS"},
		"Lee":  Student{sno: 899, name: "Lee", age: 25, major: "Physics"},
		"Kim2": Student{sno: 3, name: "Kim", age: 22, major: "Mathmatics"},
	}
	for _, value := range data {
		pq.Push(value.sno, value)
	}

	ans := []int{3, 42, 456, 899, 1234}
	for i, num := range ans {
		x := pq.Pop()
		fmt.Println(i, num, x.Key)
		assert.Equal(t, num, x.Key)
		assert.Equal(t, len(ans)-1-i, pq.Size())
	}
	assert.Equal(t, true, pq.Empty())
}
