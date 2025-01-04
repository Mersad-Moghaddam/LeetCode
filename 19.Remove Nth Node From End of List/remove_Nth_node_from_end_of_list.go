package main

import "fmt"

// ListNode defines a node of a singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// removeNthFromEnd removes the n-th node from the end of the list and returns the head.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head} // Dummy node to handle edge cases
	first, second := dummy, dummy

	// Move the first pointer n+1 steps ahead to maintain a gap of n between first and second
	for i := 0; i <= n; i++ {
		first = first.Next
	}

	// Move both pointers until the first pointer reaches the end
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// Remove the n-th node from the end
	second.Next = second.Next.Next

	return dummy.Next
}

// Helper function to create a linked list from a slice
func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, val := range nums[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

// Helper function to convert a linked list to a slice
func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

// Test the function
func main() {
	head := createList([]int{1, 2, 3, 4, 5})
	n := 2
	newHead := removeNthFromEnd(head, n)
	fmt.Println(listToSlice(newHead)) // Output: [1, 2, 3, 5]

	head = createList([]int{1})
	n = 1
	newHead = removeNthFromEnd(head, n)
	fmt.Println(listToSlice(newHead)) // Output: []

	head = createList([]int{1, 2})
	n = 1
	newHead = removeNthFromEnd(head, n)
	fmt.Println(listToSlice(newHead)) // Output: [1]
}
