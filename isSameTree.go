package main


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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q != nil {
		return false
	}

	if p != nil && q == nil {
		return false
	}

	if p == nil && q == nil {
		return true
	}

	if p.Val != q.Val {
		return false
	}

	res := isSameTree(p.Left, q.Left) || isSameTree(p.Right, q.Right)

	return res
}

func isSameTree1(p *TreeNode, q *TreeNode) bool {

	p1 := p
	q1 := q
	var t1, t2 []*TreeNode
	for {
		if p1 == nil && q1 != nil {
			return false
		}

		if p1 != nil && q1 == nil {
			return false
		}

		if p1 == nil && q1 == nil && len(t1) ==0 0 && len(t2) == 0 {
			break
		}

		if p1 == nil && q1 == nil {
			// 右
			p1 = t1[len(t1) -1]
			q1 = t2[len(t2) -1]

			// 数组回溯
			t1 = t1[:len(t1) -1]
			t2 = t2[:len(t2) -1]

			// 重新赋值
			p1 = p1.Right
			q1 = q1.Right
		} else {
			// 左
			if p1.Val != q1.Val {
				return false
			}
			t1 = append(t1, p1)
			t2 = append(t2, q1)
			p1 = p1.Left
			q1 = q1.Left
		}


	}

	return true
}

func main() {
}
