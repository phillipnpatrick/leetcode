package Solutions

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
// Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

// Example 1:
// Input: head = [1,1,2]
// Output: [1,2]

// Example 2:
// Input: head = [1,1,2,3,3]
// Output: [1,2,3]

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
        return head
    }
	dummy := head.Next
	previous := head

	for dummy != nil {
		if dummy.Val == previous.Val {
			previous.Next = dummy.Next
		} else {
			previous = previous.Next
		}		
		dummy = previous.Next
	}

    return head
}