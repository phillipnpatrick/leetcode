package Solutions

// https://leetcode.com/problems/remove-linked-list-elements/description/
// Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new head.

// Example 1:
// Input: head = [1,2,6,3,4,5,6], val = 6
// Output: [1,2,3,4,5]

// Example 2:
// Input: head = [], val = 1
// Output: []

// Example 3:
// Input: head = [7,7,7,7], val = 7
// Output: []

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	slow := head

	for slow != nil && slow.Val == val {
		if slow.Val == val {
			head = head.Next
		}
		slow = slow.Next
	}

	if head != nil && head.Next != nil {
		fast := slow.Next

		for fast != nil {
			if fast.Val == val {
				slow.Next = fast.Next
			} else {
				slow = slow.Next
			}
			fast = fast.Next
		}
	}

	return head
}
