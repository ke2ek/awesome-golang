package stringsuffix

import "fmt"

func GetLongestCommonPrefix(s string) string {
	N := len(s)
	// sfa[i] := the start index of each suffix by ascending order.
	sfa := NewSuffixArray(s).Array()

	for i, idx := range sfa {
		fmt.Printf("sfa[%d] = %d : %s\n", i, idx, s[idx:])
	}

	// prevSfa[i] := the index of the previous suffix on the suffix starting from i.
	prevSfa := make([]int, N)
	// lcp[i] := the length of lcp on the suffix starting from i which means s[i:].
	lcp := make([]int, N)

	prevSfa[sfa[0]] = -1
	for i := 1; i < N; i++ {
		prevSfa[sfa[i]] = sfa[i-1]
	}

	for i, common := 0, 0; i < N; i++ {
		if prevSfa[i] == -1 {
			lcp[i] = 0
		} else {
			for i+common < N && prevSfa[i]+common < N {
				if s[i+common] == s[prevSfa[i]+common] {
					common++
				} else {
					break
				}
			}
			lcp[i] = common
			if common > 0 {
				common--
			}
		}
	}

	maxLen, sIdx := 0, 0
	for i := 0; i < N; i++ {
		if lcp[i] > maxLen {
			maxLen = lcp[i]
			sIdx = i
		}
	}
	return s[sIdx : sIdx+maxLen]
}
