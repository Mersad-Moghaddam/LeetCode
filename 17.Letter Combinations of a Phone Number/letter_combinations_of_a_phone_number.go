package main

import (
	"fmt"
)

// letterCombinations generates all possible letter combinations that a given
// digit string could represent on a traditional telephone keypad.
//
// The function takes a string `digits` as input, where each character is a digit
// from '2' to '9'. Each digit maps to a set of letters, similar to the mapping
// on a telephone keypad. For example, '2' maps to "abc", '3' maps to "def", etc.
// The function uses a backtracking approach to explore all possible combinations
// of letters for the given digit sequence.
//
// If the input string is empty, the function returns an empty slice.
//
// Example:
// Input: digits = "23"
// Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	// Mapping of digits to corresponding letters
	digitToLetters := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var result []string
	var combination func(index int, current string)

	// Backtracking function to generate combinations
	combination = func(index int, current string) {
		if index == len(digits) {
			result = append(result, current)
			return
		}

		letters := digitToLetters[digits[index]]
		for i := 0; i < len(letters); i++ {
			combination(index+1, current+string(letters[i]))
		}
	}

	combination(0, "")
	return result
}

func main() {
	// Test cases
	fmt.Println(letterCombinations("23")) // Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
	fmt.Println(letterCombinations("89")) // Output: []
	fmt.Println(letterCombinations("2"))  // Output: ["a","b","c"]
}
