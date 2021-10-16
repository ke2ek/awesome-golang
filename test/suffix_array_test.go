package test

import (
	"awesome-golang/common"
	suffixarray "awesome-golang/strings/suffix-array"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
"banana"
0: banana
1: anana
2: nana
3: ana
4: na
5: a

suffix array:
5: a
3: ana
1: anana
0: banana
4: na
2: nana
*/

func TestSuffixArray(t *testing.T) {
	sfa := suffixarray.NewSuffixArrayNaive(common.BANANA).Array()
	sfa2 := suffixarray.NewSuffixArray(common.BANANA).Array()
	ans := []string{"a", "ana", "anana", "banana", "na", "nana"}
	for i, idx := range sfa {
		assert.Equal(t, ans[i], common.BANANA[idx:])
	}

	for i, idx := range sfa2 {
		assert.Equal(t, ans[i], common.BANANA[idx:])
	}

	sfa3 := suffixarray.NewSuffixArrayNaive(common.MISSISSIPI).Array()
	sfa4 := suffixarray.NewSuffixArray(common.MISSISSIPI).Array()
	ans = []string{"i", "ipi", "issipi", "ississipi", "mississipi", "pi", "sipi", "sissipi", "ssipi", "ssissipi"}
	for i, idx := range sfa3 {
		assert.Equal(t, ans[i], common.MISSISSIPI[idx:])
	}

	for i, idx := range sfa4 {
		assert.Equal(t, ans[i], common.MISSISSIPI[idx:])
	}
}

func TestLCPArray(t *testing.T) {
	var lcp string
	lcp = suffixarray.GetLongestCommonPrefix(common.BANANA)
	assert.Equal(t, "ana", lcp)
	lcp = suffixarray.GetLongestCommonPrefix(common.MISSISSIPI)
	assert.Equal(t, "issi", lcp)
	lcp = suffixarray.GetLongestCommonPrefix("aa")
	assert.Equal(t, "a", lcp)
}

/*
Initial: 0 1 2 3 4 5 6 7 8 9
Suffix Array:
Group[0] = 12: mississipi
Group[1] = 8: ississipi
Group[2] = 18: ssissipi
Group[3] = 18: sissipi
Group[4] = 8: issipi
Group[5] = 18: ssipi
Group[6] = 18: sipi
Group[7] = 8: ipi
Group[8] = 15: pi
Group[9] = 8: i

t = 16: 9 7 4 1 0 8 6 3 5 2
Suffix Array:
Group[9] = 0: i
Group[7] = 1: ipi
Group[4] = 2: issipi
Group[1] = 3: ississipi
Group[0] = 4: mississipi
Group[8] = 5: pi
Group[6] = 6: sipi
Group[3] = 7: sissipi
Group[5] = 8: ssipi
Group[2] = 9: ssissipi
*/
