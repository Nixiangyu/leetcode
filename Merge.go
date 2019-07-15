package main

import (
	"fmt"
	"sort"
)

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

// Input: 1->2->4, 1->3->4
// Output: 1->1->2->3->4->4
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l2 == nil && l1 != nil {
		fmt.Println("dd")
		return l1
	}
	p1 := l1
	p2 := l2
	var temp map[int]int
	temp = make(map[int]int)

	for {
		if p1 == nil && p2 == nil {
			break
		}

		if p1 != nil {
			temp[p1.Val]++
			p1 = p1.Next
		}

		if p2 != nil {
			temp[p2.Val]++
			p2 = p2.Next
		}
	}

	var keys []int
	for k := range temp {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for key, value := range temp {
		fmt.Println(key, value)
	}

	var curNode *ListNode
	var Head *ListNode
	var flag int
	for _, value := range keys {
		for i := 1; i <= temp[value]; i++ {
			if flag == 0 {
				Head = &ListNode{
					Val:value,
				}
				curNode = Head
				flag = 1
			} else {
				temp := &ListNode{
					Val:value,
				}

				curNode.Next = temp
				curNode = temp
			}
		}
	}

	return Head
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l2 == nil && l1 != nil {
		fmt.Println("dd")
		return l1
	}

	if l1 == nil && l2 != nil {
		fmt.Println("dd")
		return l2
	}

	var Head = l2
	for p := l1; p != nil; p = p.Next {
		var prvNode *ListNode
		for p1 := Head; p1 != nil; p1 = p1.Next {
			if p.Val <= p1.Val {
				if prvNode == nil {
					fmt.Println("==>>",p.Val, p1.Val)
					newHead := &ListNode{
						Val:p.Val,
					}
					newHead.Next = p1
					Head = newHead
				}else{
					fmt.Println("=dd=",p.Val, p1.Val)
					temp := &ListNode{
						Val:p.Val,
					}
					temp.Next = p1
					prvNode.Next = temp
				}
				break
			}

			prvNode = p1
		}
		if prvNode != nil && prvNode.Next == nil {
			fmt.Println("==", p.Val, p.Next, prvNode.Next, prvNode.Val)
			temp := &ListNode{
				Val:p.Val,
			}

			prvNode.Next = temp

		}
	}

	return Head
}

func main() {
	var v3 = &ListNode{
		Val:4,
	}
	var v2 = &ListNode{
		Val:3,
		Next:v3,
	}

	var v1 = &ListNode{
		Val:1,
		Next:v2,
	}
	var v6 = &ListNode{
		Val:4,
	}
	var v5 = &ListNode{
		Val:2,
		Next:v6,
	}

	var v4 = &ListNode{
		Val:1,
		Next:v5,
	}

	dd := mergeTwoLists1(v1, v4)
	for p := dd; p != nil; p = p.Next {
		fmt.Println("==>>", p.Val)
	}
}
