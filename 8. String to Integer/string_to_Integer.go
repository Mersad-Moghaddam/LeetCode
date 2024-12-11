package main

import (
	"math"
	"unicode"
)

// myAtoi converts a string to a 32-bit signed integer.
//
// The function follows these steps:
// 1. It ignores any leading whitespace.
// 2. It checks for an optional '+' or '-' sign to determine the integer's sign.
// 3. It processes consecutive digits, converting them to an integer.
// If the integer overflows, it returns the maximum or minimum 32-bit signed integer value.
// The function returns 0 if no valid conversion can be performed.
func myAtoi(s string) int {
	var i, sign int
	var result int64

	// Step 1: Ignore leading whitespace
	for i < len(s) && s[i] == ' ' {
		i++
	}

	// Step 2: Check for '+' or '-' sign
	if i < len(s) && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		} else {
			sign = 1
		}
		i++
	} else {
		sign = 1
	}

	// Step 3: Convert digits to integer
	for i < len(s) && unicode.IsDigit(rune(s[i])) {
		result = result*10 + int64(s[i]-'0')
		if result*int64(sign) < math.MinInt32 {
			return math.MinInt32
		}
		if result*int64(sign) > math.MaxInt32 {
			return math.MaxInt32
		}
		i++
	}

	// Return the result with the correct sign
	return int(result) * sign
}
