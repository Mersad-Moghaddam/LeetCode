package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

// twoSum finds two numbers in the given slice that add up to the target.
// It returns the indices of the two numbers. If no such numbers exist, it returns an empty slice.
// The function uses a hashmap to track the indices of the numbers visited so far.
// Time complexity: O(n), where n is the number of elements in the nums slice.
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if index, found := seen[complement]; found {
			return []int{index, i}
		}

		seen[num] = i
	}

	return []int{}
}
