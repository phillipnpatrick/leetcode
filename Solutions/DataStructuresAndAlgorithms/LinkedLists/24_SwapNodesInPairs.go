package Solutions

// https://leetcode.com/problems/swap-nodes-in-pairs/description/
// Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

// Example 1:
// Input: head = [1,2,3,4]
// Output: [2,1,4,3]

// Example 2:
// Input: head = []
// Output: []

// Example 3:
// Input: head = [1]
// Output: [1]

// Example 4:
// Input: head = [1,2,3]
// Output: [2,1,3]

// A -> B -> C -> D	-> E	==> B -> A -> D -> C -> E

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := head.Next
	var previous *ListNode

	for head != nil && head.Next != nil {
		if previous != nil {
			previous.Next = head.Next
		}
		previous = head

		nextNode := head.Next.Next
		head.Next.Next = head

		head.Next = nextNode
		head = nextNode
	}

	return dummy
}
