package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var res, r, l int
	if root == nil {
		return 0
	}

	if root.Right != nil {
		r = maxDepth(root.Right)
	}

	if root.Left != nil {
		l = maxDepth(root.Left)
	}

	if r > l {
		res = r + 1
	} else {
		res = l + 1
	}

	return res
}

func maxDepth1(root *TreeNode) int {
	var res, sum int
	if root == nil {
		return 0
	}

	var temp []*TreeNode
	var num map[int]int
	num = make(map[int]int)
	var p = root
	for {
		if len(temp) == 0 && p == nil {
			break
		}

		if p != nil {
			temp = append(temp, p)
			sum++
			num[p.Val] = sum
			p = p.Left

		} else {
			if res < sum {
				res = sum
			}
			p = temp[len(temp) - 1]
			sum = num[p.Val]
			temp = temp[:len(temp) - 1]
			p = p.Right
		}
	}

	return res
}

func main() {
	v5 := &TreeNode{
		Val:40,
	}
	v4 := &TreeNode{
		Val:19,
	}
	v3 := &TreeNode{
		Val:20,
		Left:v4,
		Right:v5,
	}
	v2 := &TreeNode{
		Val:9,
	}
	v1 := &TreeNode{
		Val:3,
		Left:v2,
		Right:v3,
	}

	res := maxDepth1(v1)
	fmt.Println(res)
}
