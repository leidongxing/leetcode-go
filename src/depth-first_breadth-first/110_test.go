package depth_first_breadth_first

//平衡二叉树
//自底向下递归
func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	abs := func(a int) int {
		if a < 0 {
			return -1 * a
		}
		return a
	}

	left := height(root.Left)
	right := height(root.Right)
	//只要有一个不平衡 就是-1
	if left == -1 || right == -1 || abs(left-right) > 1 {
		return -1
	}
	return max(left, right) + 1
}
