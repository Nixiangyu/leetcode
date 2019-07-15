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

// 左旋转
// k2结点不满足平衡性，它的左子树k1比右子树z深两层，k1子树中更深的是k1的左子树x，因此属于左左情况。
/*
	K2		k1
     k1     z   ->  x	     k2
   X    y		  y	 z
 */
func left(t *TreeNode) *TreeNode {
	if t == nil {
		return nil
	}
	var k1 = t.Left
	var k2 = t
	var y = k1.Right

	k1.Right = k2
	k2.Left = y

	return k1
}

// 右旋转
/*
	k2                      K1
    z	     k1      ->     k2       y
          x	 Y       z      x
 */
func right(t *TreeNode) *TreeNode {
	if t == nil {
		return nil
	}
	var k2 = t
	var k1 = t.Right
	var x = k1.Left

	k1.Left = k2
	k2.Right = x

	return k1
}

// 左右旋转
// 把k1作为根，进行一次右右旋转，旋转之后就变成了左左情况，所以第二步再进行一次左左旋转，最后得到了一棵以k2为根的平衡二叉树。
/*
	K3                       k3                      k2
    k1       d              k2       d             k1          k3
a	k2       ->      k1      c      ->     a        b  c         d
     b      c         a      b
 */

func leltRight(t *TreeNode) *TreeNode {
	if t == nil {
		return nil
	}

	tt := left(t)
	if tt == nil {
		return nil
	}

	return right(tt)
}

// 右左旋转
func rightLeft(t *TreeNode) *TreeNode {
	if t == nil {
		return nil
	}

	tt := right(t)
	if tt == nil {
		return nil
	}

	return left(tt)
}

// 判断深度
/*
func level1(root, t *TreeNode) int {
	if t == nil || root == nil {
		return 0
	}

	var l, r int
	if t.Right != nil {
		r = level(root.Right, t)
	}

	if t.Left != nil {
		l = level(root.Left, t)
	}

	if r > l {
		return r + 1
	}

	return l + 1
}
*/

func sortedArrayToBST(nums []int) *TreeNode {
	// 递归的方式实现
	if len(nums) == 0 {
		return nil
	}

	var root = &TreeNode{
		Val:nums[len(nums) / 2],
	}

	// 右子树
	root.Left = sortedArrayToBST(nums[:len(nums) / 2])

	// 左子树
	root.Right = sortedArrayToBST(nums[len(nums) / 2 + 1:])

	return root
}


func main() {


}
