package Solutions

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func PrintListNode(head *ListNode) {
    for head != nil {
        fmt.Printf("%d", head.Val)

        if head.Next != nil {
            fmt.Printf(" -> ")
        }
        head = head.Next
    }
}