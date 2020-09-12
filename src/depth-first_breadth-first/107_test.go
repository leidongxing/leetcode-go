package depth_first_breadth_first

//二叉树的层次遍历 II
func levelOrderBottom(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for i := 0; len(queue) > 0; i++ {
		result = append(result, []int{})
		var nextQueue []*TreeNode
		for j := 0; j < len(queue); j++ {
			node := queue[j]
			result[i] = append(result[i], node.Val)
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		queue = nextQueue
	}
	for i := 0; i < len(result)>>1; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	return result
}
