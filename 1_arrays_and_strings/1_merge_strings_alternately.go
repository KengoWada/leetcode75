package arraysandstrings

// MergeAlternately merges two strings by alternating characters from each string.
// If one string is longer, the remaining characters are appended at the end.
//
// Example:
//
//	MergeAlternately("abc", "pqr") -> "apbqcr"
//	MergeAlternately("ab", "pqrs") -> "apbqrs"
func MergeAlternately(word1, word2 string) string {
	len1 := len(word1)
	len2 := len(word2)
	result := ""

	size := min(len1, len2)
	for i := range size {
		result = result + string(word1[i]) + string(word2[i])
	}

	if len1 > size {
		result = result + word1[size:]
	}

	if len2 > size {
		result = result + word2[size:]
	}

	return result
}
