/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // Dummy node to build the result list
	current := dummy
	carry := 0

	// Iterate through both linked lists and handle carry
	for l1 != nil || l2 != nil || carry != 0 {
		// Get values from the current nodes or 0 if a list is shorter
		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		// Calculate the sum and carry
		sum := val1 + val2 + carry
		carry = sum / 10

		// Append the digit to the result list
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return dummy.Next
}
