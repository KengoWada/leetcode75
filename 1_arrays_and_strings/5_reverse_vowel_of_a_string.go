package arraysandstrings

import "strings"

func ReverseVowels(s string) string {
	// var allVowels = []rune{'a', 'e', 'i', 'o', 'u'}

	// var vowels = []rune{}
	// for _, char := range s {
	// 	if slices.Contains(allVowels, unicode.ToLower(char)) {
	// 		vowels = append(vowels, char)
	// 	}
	// }

	// result := ""
	// for _, char := range s {
	// 	if slices.Contains(allVowels, unicode.ToLower(char)) {
	// 		lastIndex := len(vowels) - 1
	// 		char = vowels[lastIndex]
	// 		vowels = vowels[:lastIndex]
	// 	}

	// 	result += string(char)
	// }

	// return result

	word := []rune(s)
	start := 0
	end := len(word) - 1
	vowels := "aeiouAEIOU"

	for start < end {
		for start < end && !strings.ContainsRune(vowels, word[start]) {
			start++
		}
		for start < end && !strings.ContainsRune(vowels, word[end]) {
			end--
		}

		word[start], word[end] = word[end], word[start]
		start++
		end--
	}

	return string(word)
}
