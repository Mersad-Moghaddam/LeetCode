package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Start with the first string as the prefix candidate
	prefix := strs[0]

	// Compare the prefix with each string in the array
	for _, str := range strs[1:] {
		// Trim the prefix until it matches the start of the current string
		for len(prefix) > 0 && !startsWith(str, prefix) {
			prefix = prefix[:len(prefix)-1]
		}
		// If the prefix becomes empty, no common prefix exists
		if prefix == "" {
			return ""
		}
	}

	return prefix
}

// Helper function to check if a string starts with a given prefix
func startsWith(s, prefix string) bool {
	if len(prefix) > len(s) {
		return false
	}
	return s[:len(prefix)] == prefix
}

func main() {
	// Test cases
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // Output: "fl"
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))    // Output: ""
}
