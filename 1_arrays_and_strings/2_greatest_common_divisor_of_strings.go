package arraysandstrings

// GCDOfStrings finds the greatest common divisor (GCD) of two strings.
// If there is no common divisor string, it returns an empty string.
//
// A string X is a divisor of string Y if Y can be formed by repeating X multiple times.
//
// Example:
//
//	GCDOfStrings("ABCABC", "ABC") -> "ABC"
//	GCDOfStrings("ABABAB", "ABAB") -> "AB"
//	GCDOfStrings("LEET", "CODE") -> ""
func GCDOfStrings(str1, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	len1 := len(str1)
	len2 := len(str2)
	for len2 != 0 {
		temp := len2
		len2 = len1 % len2
		len1 = temp
	}

	return str1[:len1]
}
