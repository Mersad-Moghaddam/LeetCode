package main

// longestPalindrome returns the longest palindromic substring of the given string.
// Time complexity: O(n^2), where n is the length of the string.
// Space complexity: O(n^2), for the DP table.
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	n := len(s)
	start, maxLength := 0, 1

	// Create a DP table to store palindrome information
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	// Every single character is a palindrome
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	// Check for substrings of length 2
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			start = i
			maxLength = 2
		}
	}

	// Check for substrings of length > 2
	for length := 3; length <= n; length++ {
		for i := 0; i < n-length+1; i++ {
			j := i + length - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				start = i
				maxLength = length
			}
		}
	}

	return s[start : start+maxLength]
}
