package main

// lengthOfLongestSubstring returns the length of the longest substring
// without repeating characters of the given string.
//
// Time complexity: O(n), where n is the length of the string.
// Space complexity: O(min(n, m)), where m is the size of the character set.
func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int)
	maxLength := 0
	start := 0

	for i := 0; i < len(s); i++ {
		if index, found := charIndex[s[i]]; found && index >= start {
			start = index + 1
		}
		charIndex[s[i]] = i
		maxLength = max(maxLength, i-start+1)
	}

	return maxLength
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
