package main

import "math"

// reverse reverses the digits of a 32-bit signed integer.
// If the reversed integer overflows, it returns 0.
// The function handles both positive and negative integers.
func reverse(x int) int {
	// Define the range of 32-bit signed integers
	const intMin = math.MinInt32 // -2^31
	const intMax = math.MaxInt32 // 2^31 - 1

	result := 0
	for x != 0 {
		// Extract the last digit
		pop := x % 10
		x /= 10

		// Check for overflow/underflow before updating the result
		if result > (intMax-pop)/10 || result < (intMin-pop)/10 {
			return 0
		}

		// Add the digit to the reversed number
		result = result*10 + pop
	}

	return result
}
