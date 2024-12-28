package main

import (
	"fmt"
)

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
