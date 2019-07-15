package main

func hasPathSum(root *TreeNode, sum,  temp int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		if temp + root.Val == sum {
			return true
		}
	}


	r := hasPathSum(root.Right, temp + root.Val , sum)
	l := hasPathSum(root.Left, temp + root.Val , sum)
	if r || l {
		return true
	}

	return false
}


func main() {
}
