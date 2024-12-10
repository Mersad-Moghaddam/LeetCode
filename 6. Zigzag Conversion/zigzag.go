package main

import "strings"

// convert rearranges a string into a zigzag pattern on a given number of rows
// and reads it row by row.
//
// The function takes a string `s` and an integer `numRows` as inputs. It places
// characters from `s` into `numRows` rows in a zigzag pattern, then concatenates
// and returns the characters row by row.
//
// If `numRows` is 1 or the length of `s` is less than or equal to `numRows`,
// the original string is returned unchanged.
func convert(s string, numRows int) string {
	// If numRows is 1 or the string length is less than the number of rows, return the string as it is
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	// Create a slice to hold rows
	rows := make([]string, numRows)
	curRow := 0
	goingDown := false

	// Iterate through each character in the string
	for _, char := range s {
		rows[curRow] += string(char)
		// Change direction at the top and bottom rows
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		// Move up or down the rows
		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	// Join all rows to create the final result
	return strings.Join(rows, "")
}
