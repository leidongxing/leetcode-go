package depth_first_breadth_first

//打家劫舍 III
func rob(root *TreeNode) int {
	val := dfsRob(root)
	return max(val[0], val[1])
}

func dfsRob(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}
	left, right := dfsRob(node.Left), dfsRob(node.Right)
	selected := node.Val + left[1] + right[1]
	unselected := max(left[0], left[1]) + max(right[0], right[1])
	return []int{selected, unselected}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
