package main

import (
	"math"
)

// findMedianSortedArrays returns the median of two sorted arrays of integers.
// The time complexity is O(log(min(x, y))), where x and y are the lengths of the input arrays.
// The space complexity is O(1).
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Ensure nums1 is the smaller array
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	x, y := len(nums1), len(nums2)
	low, high := 0, x

	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (x+y+1)/2 - partitionX

		// Handle edge cases for left and right partitions
		maxLeftX := math.MinInt64
		if partitionX != 0 {
			maxLeftX = nums1[partitionX-1]
		}

		minRightX := math.MaxInt64
		if partitionX != x {
			minRightX = nums1[partitionX]
		}

		maxLeftY := math.MinInt64
		if partitionY != 0 {
			maxLeftY = nums2[partitionY-1]
		}

		minRightY := math.MaxInt64
		if partitionY != y {
			minRightY = nums2[partitionY]
		}

		// Check if we found the correct partitions
		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			// Check if the total number of elements is odd or even
			if (x+y)%2 == 0 {
				return float64(max(maxLeftX, maxLeftY)+min(minRightX, minRightY)) / 2
			} else {
				return float64(max(maxLeftX, maxLeftY))
			}
		} else if maxLeftX > minRightY {
			// Move towards the left in nums1
			high = partitionX - 1
		} else {
			// Move towards the right in nums1
			low = partitionX + 1
		}
	}

	// Input arrays are not valid, return 0
	return 0
}

// max returns the larger of two integers a and b.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min returns the minimum of two integers a and b.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
