package depth_first_breadth_first

//二叉树的层平均值
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	var result []float64
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var nextQueue []*TreeNode
		next := 0
		for j := 0; j < len(queue); j++ {
			node := queue[j]
			next += node.Val
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		result = append(result, float64(next)/float64(len(queue)))
		queue = nextQueue
	}
	return result
}
