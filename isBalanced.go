package main

import "math"

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

func nodeLevel(root *TreeNode) int {
	if root == nil {
		return 0
	}
	r := nodeLevel(root.Right)
	l := nodeLevel(root.Left)
	if  r > l  {
		return r +1
	}

	return l+1
}
func isBalanced(root *TreeNode) bool {
	// 求叶子节点的深度差，超过1的就返回false
	r := nodeLevel(root.Right)
	l := nodeLevel(root.Left)
	if math.Abs(r - l) > 1 {
		return false
	}

	rr := isBalanced(root.Right)
	ll := isBalanced(root.Left)
	if !rr || !ll {
		return false
	}

	return true

}

func
main()
{
}
