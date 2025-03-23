package arraysandstrings

import "strings"

func ReverseWords(s string) string {
	wordSlice := strings.Fields(s)
	start, end := 0, len(wordSlice)-1

	for start < end {
		wordSlice[start], wordSlice[end] = wordSlice[end], wordSlice[start]
		start++
		end--
	}

	return strings.Join(wordSlice, " ")
}
