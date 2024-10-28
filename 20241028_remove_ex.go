package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result := &ListNode{}
	result.Next = head
	cur := result
	var pre *ListNode
	i := 1
	for head != nil {
		if i >= n {
			pre = cur
			cur = cur.Next
		}
		head = head.Next
		i++
	}
	pre.Next = pre.Next.Next
	return result.Next
}
