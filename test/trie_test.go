package test

import (
	"awesome-golang/common"
	"awesome-golang/strings/trie"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	tc1 := []string{"banana", "banhana", "apple", "append", "count", "compare"}
	tc2 := []string{}
	for i := 0; i < len(common.MISSISSIPI); i++ {
		tc2 = append(tc2, common.MISSISSIPI[i:])
	}

	myTrie := trie.New()
	suffixTrie := trie.New()

	for _, s := range tc1 {
		myTrie.Insert(s)
	}
	for _, s := range tc2 {
		suffixTrie.Insert(s)
	}

	for _, s := range tc1 {
		assert.Equal(t, true, myTrie.Find(s))
		assert.Equal(t, false, suffixTrie.Find(s))
	}
	for _, s := range tc2 {
		assert.Equal(t, true, suffixTrie.Find(s))
		assert.Equal(t, false, myTrie.Find(s))
	}
}
