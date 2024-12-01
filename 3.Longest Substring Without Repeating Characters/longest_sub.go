package main

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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
