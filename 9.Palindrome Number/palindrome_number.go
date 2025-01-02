package main

import "fmt"

// isPalindrome determines if the given integer is a palindrome.
// A palindrome is a number or a text phrase that reads the same backwards as forwards.
// A single digit is always a palindrome.
// Examples: 12321, 10001, 121, 123454321, 1, 5
func isPalindrome(x int) bool {
	// Negative numbers and numbers ending with 0 (except 0 itself) are not palindromes
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversedHalf := 0
	for x > reversedHalf {
		reversedHalf = reversedHalf*10 + x%10
		x /= 10
	}

	// Check if the original half matches the reversed half
	return x == reversedHalf || x == reversedHalf/10
}

// Examples
func main() {
	fmt.Println(isPalindrome(121))  // Output: true
	fmt.Println(isPalindrome(-121)) // Output: false
	fmt.Println(isPalindrome(10))   // Output: false
}
