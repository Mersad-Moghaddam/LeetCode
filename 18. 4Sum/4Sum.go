package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	n := len(nums)

	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, n-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return result
}

func main() {
	nums1 := []int{1, 0, -1, 0, -2, 2}
	target1 := 0
	fmt.Println(fourSum(nums1, target1)) // Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

	nums2 := []int{2, 2, 2, 2, 2}
	target2 := 8
	fmt.Println(fourSum(nums2, target2)) // Output: [[2,2,2,2]]
}
