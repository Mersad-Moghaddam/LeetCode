package main

import (
	"fmt"
)

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
