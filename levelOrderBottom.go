package main

import (
	"test/common/number"
	"fmt"
	"html/template"
)

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

//非递归
func levelOrderBottom1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var p *TreeNode
	var curNum int
	var temp []*TreeNode
	temp = append(temp, root)
	var res [][]int
	var r1 [][]int
	for len(temp) != 0 {
		var t []int
		curNum  = len(temp)
		for i:= 0; i < curNum;i++ {
			p = temp[i]
			if p.Left != nil {
				temp = append(temp, p.Left)
			}
			if p.Right != nil {
				temp = append(temp, p.Right)
			}
			t = append(t, p.Val)
		}
		temp = temp[curNum:]
		res = append(res, t)
	}


	r1 = make([][]int, len(res))
	for i:=0; i < len(res);i++ {
		r1[i] = res[len(res) -i -1]
	}

	return r1
}

func level(root []*TreeNode, i int) [][]int {
	if len(root) == 0 {
		return nil
	}

	curPoint := root[len(root) - 1]
	root = root[:len(root) - 1]
	if curPoint.Left != nil {
		root = append(root, curPoint.Left)
	}

	if curPoint.Right != nil {
		root = append(root, curPoint.Right)
	}

	res := level(root, i+1)
	if res != nil {
		res[i] = append(res[i], curPoint.Val)
	}

	return res
}
// 递归 层序遍历
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var temp []*TreeNode
	temp = append(temp, root)
	res := level(root, 0)
	return res
}

func main() {
}
