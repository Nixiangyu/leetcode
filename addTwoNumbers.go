package main

import (
	"fmt"
	"text/scanner"
	"math/big"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var cur  = head
	var each_num int
	for {
		if l1 == nil && l2 == nil && each_num == 0 {
			break
		}

		if l1 != nil {
			each_num = l1.Val + each_num
			l1 = l1.Next
		}

		if l2 != nil {
			each_num = l2.Val + each_num

			l2 = l2.Next
		}

		var temp = &ListNode{
			Val: each_num % 10,
		}

		cur.Next = temp
		cur = temp
		each_num = each_num / 10
	}


	return head.Next
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode, t int) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	var r int
	var t1, t2 *ListNode
	if l1 != nil {
		r = l1.Val + r
		t1 = l1.Next
	}

	if l2 != nil {
		r = l2.Val + r
		t2 = l2.Next
	}

	var temp = &ListNode{
		Val: (r + t) % 10,
	}

	res := addTwoNumbers1(t1, t2, (r+t)/10)
	temp.Next = res

	return temp

}

func main() {
}
