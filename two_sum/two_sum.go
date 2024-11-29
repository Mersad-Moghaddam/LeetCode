package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
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
