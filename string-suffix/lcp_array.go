package stringsuffix

import "fmt"

func GetLongestCommonPrefix(s string) string {
	N := len(s)
	sfa := NewSuffixArray(s).Array()
	prevSfa := make([]int, N)
	prevLcp := make([]int, N)
	lcp := make([]int, N)

	prevSfa[sfa[0]] = -1
	for i := 1; i < N; i++ {
		prevSfa[sfa[i]] = sfa[i-1]
	}

	for i, common := 0, 0; i < N; i++ {
		if prevSfa[i] == -1 {
			prevLcp[i] = 0
		} else {
			for i+common < N && prevSfa[i]+common < N {
				if s[i+common] == s[prevSfa[i]+common] {
					common++
				} else {
					break
				}
			}
			prevLcp[i] = common
			if common > 0 {
				common--
			}
		}
	}

	maxLen, start := 0, 0
	for i := 0; i < N; i++ {
		lcp[i] = prevLcp[sfa[i]]
		if lcp[i] > maxLen {
			maxLen = lcp[i]
			start = i
		}
	}

	fmt.Println(lcp)
	return s[start+1 : start+1+maxLen]
}
