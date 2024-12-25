package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	// Sort the array
	sort.Ints(nums)
	n := len(nums)
	result := [][]int{}

	// Iterate through the array
	for i := 0; i < n-2; i++ {
		// Skip duplicate values for the first number
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1

		// Use two pointers to find valid triplets
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicate values for the second and third numbers
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// Move pointers inward
				left++
				right--
			} else if sum < 0 {
				left++ // Move the left pointer to increase the sum
			} else {
				right-- // Move the right pointer to decrease the sum
			}
		}
	}

	return result
}

func main() {
	// Example test cases
	nums1 := []int{-1, 0, 1, 2, -1, -4}
	nums2 := []int{0, 1, 1}
	nums3 := []int{0, 0, 0}

	fmt.Println("Example 1:", threeSum(nums1)) // Output: [[-1,-1,2],[-1,0,1]]
	fmt.Println("Example 2:", threeSum(nums2)) // Output: []
	fmt.Println("Example 3:", threeSum(nums3)) // Output: [[0,0,0]]
}
