package strings

// O(N)
func ReverseWords(s string) string {
	reversedString := ""
	start, end := len(s)-1, len(s)-1
	for end >= 0 {
		// Skip whitespaces.
		for s[start] == ' ' {
			start--
			if start < 0 {
				break
			}
		}
		end = start
		// March down to the index before the beginning of the word.
		if start >= 0 {
			for s[start] != ' ' {
				start--
				if start < 0 {
					break
				}
			}
		}
		// start = (start index of each word - 1), end = end index of each word
		reversedString += s[start+1 : end+1]
		// If this isn't the last word, add one whitespace after this in the buffer.
		if start > 0 {
			reversedString += " "
		}
		end = start - 1
		start = end
	}
	// Check if we have added one more whitespaces to the buffer.
	if reversedString[len(reversedString)-1] == ' ' {
		reversedString = reversedString[:len(reversedString)-1]
	}
	return reversedString
}
