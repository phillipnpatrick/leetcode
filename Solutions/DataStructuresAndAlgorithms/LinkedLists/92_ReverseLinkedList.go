package Solutions

// https://leetcode.com/problems/reverse-linked-list-ii/description/
// Given the head of a singly linked list and two integers left and right where left <= right, reverse the nodes of the list from position left to position right,
// and return the reversed list.

// Example 1:
// Input: head = [1,2,3,4,5], left = 2, right = 4
// Output: [1,4,3,2,5]

// Example 2:
// Input: head = [5], left = 1, right = 1
// Output: [5]

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil || right - left < 1 {
		return head
	}

	index := 1
	var list, leftNode *ListNode

	for head != nil && index < left {
		if leftNode == nil {
			leftNode = &ListNode{Val: head.Val}
			list = leftNode
		} else {
			leftNode.Val = head.Val
		}
		
		if index + 1 < left {
			leftNode.Next = &ListNode{}
			leftNode = leftNode.Next
		}
		head = head.Next
		index++
	}

	var middleNode *ListNode
	for middleNode == nil || index <= right {
		nextNode := head.Next // first, make sure we don't lose the next node
		head.Next = middleNode  // reverse the direction of the pointer
		middleNode = head       // set the current node to prev for the next node
		head = nextNode       // move on
		index++
	}

	if list == nil {
		list = middleNode
	} else {
		leftNode.Next = middleNode
	}
	
	for middleNode.Next != nil {
		middleNode = middleNode.Next
	}
	middleNode.Next = head

	return list
}
