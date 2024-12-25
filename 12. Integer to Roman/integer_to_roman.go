package main

import "fmt"


func intToRoman(num int) string {
	// Define the Roman numerals and corresponding values
	romanNumerals := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	result := ""

	// Iterate through the Roman numerals array
	for _, rn := range romanNumerals {
		// While the current value fits into the number
		for num >= rn.value {
			result += rn.symbol
			num -= rn.value
		}
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(intToRoman(3749)) // Output: MMMDCCXLIX
	fmt.Println(intToRoman(58))   // Output: LVIII
	fmt.Println(intToRoman(1994)) // Output: MCMXCIV
}
