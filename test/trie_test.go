package test

import (
	"awesome-golang/common"
	"awesome-golang/trie"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	tc1 := []string{"banana", "banhana", "apple", "append", "count", "compare"}
	tc2 := []string{}
	tc3 := []string{}
	for i := 0; i < len(common.MISSISSIPI); i++ {
		tc2 = append(tc2, common.MISSISSIPI[i:])
	}
	for i := 0; i < len(common.BANANA); i++ {
		tc3 = append(tc3, common.BANANA[i:])
	}

	myTrie := trie.New()
	suffixTrie := trie.New()
	iterTrie := trie.New()

	for _, s := range tc1 {
		myTrie.Insert(s)
	}
	for _, s := range tc2 {
		suffixTrie.Insert(s)
	}
	for _, s := range tc3 {
		iterTrie.InsertIterative(s)
	}

	for _, s := range tc1 {
		assert.Equal(t, true, myTrie.Find(s))
		assert.Equal(t, false, suffixTrie.Find(s))
	}
	for _, s := range tc2 {
		assert.Equal(t, true, suffixTrie.Find(s))
		assert.Equal(t, false, myTrie.Find(s))
	}
	for _, s := range tc3 {
		assert.Equal(t, true, iterTrie.FindIterative(s))
		assert.Equal(t, false, suffixTrie.Find(s))
	}
}

func TestSuffixTree(t *testing.T) {
	tree := trie.NewSuffixTree("abbc")
	assert.Equal(t, 6, tree.Count())
	assert.Equal(t, 4, tree.CountLeaf())
	tree.FreeSuffixTreeByPostOrder(tree.Root())

	tree2 := trie.NewSuffixTree("abcabxabcd$")
	tree2.PrintPretty(tree2.Root(), 0)
	assert.Equal(t, true, tree2.HasSubString("bxa"))
	assert.Equal(t, true, tree2.HasSubString("abcd"))
	assert.Equal(t, false, tree2.HasSubString("bxad"))
	assert.Equal(t, 17, tree2.Count())
	assert.Equal(t, 11, tree2.CountLeaf())
	tree2.FreeSuffixTreeByPostOrder(tree2.Root())

	tree3 := trie.NewSuffixTree("GEEKSFORGEEKS$")
	assert.Equal(t, 21, tree3.Count())
	tree3.FreeSuffixTreeByPostOrder(tree3.Root())
}

func TestLongestRepeatedSubstring(t *testing.T) {
	testcases := []string{
		"GEEKSFORGEEKS$",
		"AAAAAAAAAA$",
		"ABCDEFG$",
		"ABABABA$",
		"ATCGATCGA$",
		"banana$",
		"abcpqrabpqpq$",
		"pqrpqpqabab$",
	}
	ans := []string{"GEEKS", "AAAAAAAAA", "", "ABABA", "ATCGA", "ana", "ab", "ab"}
	for i, tc := range testcases {
		assert.Equal(t, ans[i], trie.LongestRepeatedSubstring(tc))
	}
}

func TestMakeSuffixArray(t *testing.T) {
	testcases := []string{
		"GEEKSFORGEEKS$",
		"banana$",
		"AAAAAAAAAA$",
		"ABCDEFG$",
		"ABABABA$",
		"abcabxabcd$",
		"CCAAACCCGATTA$",
	}
	ans := [][]int{
		{9, 1, 10, 2, 5, 8, 0, 11, 3, 6, 7, 12, 4},
		{5, 3, 1, 0, 4, 2},
		{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		{0, 1, 2, 3, 4, 5, 6},
		{6, 4, 2, 0, 5, 3, 1},
		{0, 6, 3, 1, 7, 4, 2, 8, 9, 5},
		{12, 2, 3, 4, 9, 1, 0, 5, 6, 7, 8, 11, 10},
	}
	for i, tc := range testcases {
		tree := trie.NewSuffixTree(tc)
		assert.Equal(t, ans[i], tree.SuffixArray())
		tree.PrintPretty(tree.Root(), 0)
		tree.FreeSuffixTreeByPostOrder(tree.Root())
	}
}

func TestLongestCommonSubstring(t *testing.T) {
	testcases := [][]string{
		{"xabxac", "abcabxabcd"},
		{"xabxaabxa", "babxba"},
		{"GeeksforGeeks", "GeeksQuiz"},
		{"abcde", "fghie"},
		{"pqrst", "uvwxyz"},
		{"xabcyz", "abcdeabcde"},
	}
	ans := []string{"abxa", "abx", "Geeks", "e", "", "abc"}
	for i, tc := range testcases {
		assert.Equal(t, ans[i], trie.LongestCommonSubstring(tc[0], tc[1]))
	}
}

func TestLongestPalindromicSubstring(t *testing.T) {
	testcases := []string{
		"cabbaabb", "forgeeksskeegfor",
		"abcde", "abcdae", "abacd",
		"abcdc", "abacdfgdcaba", "xabax",
		"xyabacdfgdcaba", "xababayz",
	}
	ans := []string{"bbaabb", "geeksskeeg", "a", "a", "aba", "cdc", "aba", "xabax", "aba", "ababa"}
	for i, tc := range testcases {
		assert.Equal(t, ans[i], trie.LongestPalindromicSubstring(tc))
	}
}
