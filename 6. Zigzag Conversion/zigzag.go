package main

import "strings"

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
