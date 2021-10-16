package strings

// O(N)
func CountWords(s string) int {
	numWords := 0
	start, end := 0, 0
	for start < len(s) {
		// Skip whitespaces.
		for s[end] == ' ' {
			end++
			if end == len(s) {
				break
			}
		}
		start = end
		// March up the index before the ending of the word.
		if end < len(s) {
			for s[end] != ' ' {
				end++
				if end == len(s) {
					break
				}
			}
			// start = the start index of each word
			// end = the end index of each word + 1
			numWords++
		}
		start = end + 1
		end = start
	}
	return numWords
}
