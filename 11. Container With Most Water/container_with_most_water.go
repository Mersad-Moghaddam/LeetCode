package main

import (
	"fmt"
)




func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0

	for left < right {
		// Calculate the current area
		width := right - left
		h := 0
		if height[left] < height[right] {
			h = height[left]
			left++
		} else {
			h = height[right]
			right--
		}
		currentArea := width * h

		// Update the maximum area if the current one is larger
		if currentArea > maxWater {
			maxWater = currentArea
		}
	}

	return maxWater
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("Maximum water that can be contained:", maxArea(height))
}
