package depth_first_breadth_first

//层数最深叶子节点的和
func deepestLeavesSum(root *TreeNode) int {
	depth := 0
	depth, max := getDepth(root, depth)
	return max
}

func getDepth(node *TreeNode, depth int) (int, int) {
	leftDepth, leftMax, rightDepth, rightMax := 0, 0, 0, 0
	depth++
	if node.Left != nil {
		leftDepth, leftMax = getDepth(node.Left, depth)
	}
	if node.Right != nil {
		rightDepth, rightMax = getDepth(node.Right, depth)
	}

	if node.Left == nil && node.Right == nil {
		return depth, node.Val
	}

	if leftDepth > rightDepth {
		return leftDepth, leftMax
	} else if leftDepth == rightDepth {
		return leftDepth, leftMax + rightMax
	} else {
		return rightDepth, rightMax
	}
}
