package main

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Right == nil && root.Left == nil {
		return 1
	}

	r := minDepth(root.Right)
	l := minDepth(root.Left)

	if r == 0 {
		return l+1
	}

	if l == 0 {
		return r +1
	}

	if r < l {
		return r +1
	}

	return l + 1
}

func main() {
}