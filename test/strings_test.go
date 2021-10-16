package test

import (
	"awesome-golang/strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	s := "This is the reverse test of a given string."
	ans := "string. given a of test reverse the is This"
	assert.Equal(t, ans, strings.ReverseWords(s))

	s = "       This is the reverse    test of      a given string.    "
	assert.Equal(t, ans, strings.ReverseWords(s))
}

func TestIsPalindrome(t *testing.T) {
	s1 := "aaabbbcccbbbaaa"
	s2 := "aaabbbccba"
	assert.Equal(t, true, strings.IsPalindrome(s1))
	assert.Equal(t, false, strings.IsPalindrome(s2))
}

func TestCountWords(t *testing.T) {
	s1 := "This is the count test."
	s2 := "   This is the    count    test.    "
	assert.Equal(t, 5, strings.CountWords(s1))
	assert.Equal(t, 5, strings.CountWords(s2))
}
