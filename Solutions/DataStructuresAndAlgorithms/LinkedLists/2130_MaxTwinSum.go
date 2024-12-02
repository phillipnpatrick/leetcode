package Solutions

import "math"

// https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/description/
// In a linked list of size n, where n is even, the ith node (0-indexed) of the linked list is known as the twin of the (n-1-i)th node, if 0 <= i <= (n / 2) - 1.
// For example, if n = 4, then node 0 is the twin of node 3, and node 1 is the twin of node 2. These are the only nodes with twins for n = 4.
// The twin sum is defined as the sum of a node and its twin.
// Given the head of a linked list with even length, return the maximum twin sum of the linked list.

// Example 1:
// Input: head = [5,4,2,1]
// Output: 6
// Explanation:
// Nodes 0 and 1 are the twins of nodes 3 and 2, respectively. All have twin sum = 6.
// There are no other nodes with twins in the linked list.
// Thus, the maximum twin sum of the linked list is 6.

// Example 2:
// Input: head = [4,2,2,3]
// Output: 7
// Explanation:
// The nodes with twins present in this linked list are:
// - Node 0 is the twin of node 3 having a twin sum of 4 + 3 = 7.
// - Node 1 is the twin of node 2 having a twin sum of 2 + 2 = 4.
// Thus, the maximum twin sum of the linked list is max(7, 4) = 7.

// Example 3:
// Input: head = [1,100000]
// Output: 100001
// Explanation:
// There is only one node with a twin in the linked list having twin sum of 1 + 100000 = 100001.

func pairSum(head *ListNode) int {
	sum := 0
	if head == nil || head.Next == nil {
		return sum
	}

	// 1. Find the middle of the linked list using the fast and slow pointer technique from the previous article.
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 2. Once at the middle of the linked list, perform a reversal. Basically, reverse only the second half of the list.
	var previous *ListNode
	curr := slow
	for curr != nil {
		nextNode := curr.Next // first, make sure we don't lose the next node
		curr.Next = previous  // reverse the direction of the pointer
		previous = curr       // set the current node to prev for the next node
		curr = nextNode       // move on
	}

	// 3. After reversing the second half, every node is spaced n / 2 apart from its pair node, where n is the number of nodes in the list which we can find from step 1.
	// 4. With that in mind, create another fast pointer n / 2 ahead of slow. Now, just iterate n / 2 times from head to find every pair sum slow.val + fast.val.
	slow = head
	fast = previous
	for fast != nil && fast.Next != nil {
		sum = int(math.Max(float64(sum), float64(slow.Val+fast.Val)))
		slow = slow.Next
		fast = fast.Next
	}

	return sum
}