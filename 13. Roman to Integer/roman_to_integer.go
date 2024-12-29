package main

import (
	"fmt"
)

// romanToInt converts a Roman numeral string to an integer.
//
// The function takes a Roman numeral string as an argument and returns the
// corresponding integer value.
//
// The function works by iterating through the string and adding the values
// of the Roman numerals to a running total. If the current value is less
// than the next value, it means that the current value should actually be
// subtracted from the total (because in Roman numerals, a smaller number in
// front of a larger one means subtraction), so it is subtracted instead.
func romanToInt(s string) int {
	// Map of Roman numeral values
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	n := len(s)

	// Iterate through the string
	for i := 0; i < n; i++ {
		// If the current value is less than the next value, subtract it
		if i < n-1 && romanMap[s[i]] < romanMap[s[i+1]] {
			total -= romanMap[s[i]]
		} else {
			// Otherwise, add it
			total += romanMap[s[i]]
		}
	}

	return total
}

func main() {
	// Example cases
	fmt.Println(romanToInt("III"))     // Output: 3
	fmt.Println(romanToInt("LVIII"))   // Output: 58
	fmt.Println(romanToInt("MCMXCIV")) // Output: 1994
}
