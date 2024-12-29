package main

import (
	"fmt"
	"math"
	"sort"
)

// threeSumClosest finds the sum of three numbers in the given slice that is closest to the target.
// The function uses a two-pointer approach to find the closest sum and returns it.
// Time complexity: O(n^2), where n is the length of the nums slice.
// Space complexity: O(1).
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums) // Sort the array
	n := len(nums)
	closestSum := nums[0] + nums[1] + nums[2]

	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1
		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]

			// Update closestSum if the currentSum is closer to the target
			if math.Abs(float64(target-closestSum)) > math.Abs(float64(target-currentSum)) {
				closestSum = currentSum
			}

			// Move pointers based on comparison with the target
			if currentSum < target {
				left++
			} else if currentSum > target {
				right--
			} else {
				// If the currentSum is exactly equal to the target
				return currentSum
			}
		}
	}
	return closestSum
}

func main() {
	// Test cases
	nums1 := []int{-1, 2, 1, -4}
	target1 := 1
	fmt.Println(threeSumClosest(nums1, target1)) // Output: 2

	nums2 := []int{0, 0, 0}
	target2 := 1
	fmt.Println(threeSumClosest(nums2, target2)) // Output: 0
}
