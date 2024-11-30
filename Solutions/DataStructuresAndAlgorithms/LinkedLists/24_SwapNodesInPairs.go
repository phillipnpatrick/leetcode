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

	var previous *ListNode
	current := head

	for current != nil && current.Next != nil {
		nextNode := current.Next
		nextNextNode := current.Next.Next
		if previous == nil {
			head = nextNode
			previous = nextNode
			previous.Next = current
		} else {
			previous.Next = nextNode
			nextNode.Next = current
		}
		current.Next = nextNextNode
		previous = current
		current = nextNextNode
	}

	return head
}