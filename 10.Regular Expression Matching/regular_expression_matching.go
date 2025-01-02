package main

import "fmt"

// isMatch returns true if the string `s` matches the given pattern `p`.
// The pattern can contain two special characters: '.' and '*'. '.' matches any single character,
// and '*' matches zero or more of the preceding element.
//
// The time complexity is O(mn), where m and n are the lengths of string `s` and pattern `p` respectively.
// The space complexity is also O(mn).
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	// dp[i][j] means if s[0..i-1] matches p[0..j-1]
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// Base case: empty string matches empty pattern
	dp[0][0] = true

	// Handle patterns like a*, a*b*, a*b*c* that can match empty string
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2] // '*' acts as zero of the preceding element
		}
	}

	// Fill the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '.' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// '*' can match zero of the preceding element
				dp[i][j] = dp[i][j-2] || (dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.'))
			}
		}
	}

	return dp[m][n]
}

// main demonstrates the usage of the isMatch function with various examples.
// It prints the results of matching strings with patterns containing '.' and '*'
// as special characters. The expected outputs are provided in the comments.

func main() {
	fmt.Println(isMatch("aa", "a"))                   // Output: false
	fmt.Println(isMatch("aa", "a*"))                  // Output: true
	fmt.Println(isMatch("ab", ".*"))                  // Output: true
	fmt.Println(isMatch("aab", "c*a*b"))              // Output: true
	fmt.Println(isMatch("mississippi", "mis*is*p*.")) // Output: false
}
