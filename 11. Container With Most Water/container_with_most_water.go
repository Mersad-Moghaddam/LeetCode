package main

import (
	"fmt"
)

// maxArea returns the maximum area that can be contained in a container given the heights of its edges.
//
// The function takes a slice of integers representing the heights of the container's edges as input and returns
// the maximum area that can be contained. The heights are given in the order of the leftmost edge to the
// rightmost edge of the container.
//
// The algorithm uses a two-pointer approach, starting from the leftmost and rightmost edges and moving
// towards the center. It calculates the current area and updates the maximum area if the current one is
// larger. The time complexity is O(n), where n is the length of the input slice.
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
