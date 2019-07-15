package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	for p := head; p != nil;  {
		if p.Next != nil &&p.Val == p.Next.Val {
			p.Next = p.Next.Next
			continue
		}
		p = p.Next
	}
	return head
}

func main() {
/*	l5 := &ListNode{
		Val:3,
	}
	l4 := &ListNode{
		Val:3,
		Next:l5,
	}*/
	l3 := &ListNode{
		Val:1,
		//Next:l4,
	}
	l2 := &ListNode{
		Val:1,
		Next:l3,
	}
	l1 := &ListNode{
		Val:1,
		Next:l2,
	}
	dd := deleteDuplicates(l1)
	for p := dd; p != nil ;p=p.Next {
		fmt.Println(p.Val)
	}
}