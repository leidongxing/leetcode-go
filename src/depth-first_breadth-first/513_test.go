package depth_first_breadth_first

//找树左下角的值
func findBottomLeftValue1(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	node := root
	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
	}
	return node.Val
}

//递归解决
var maxLevel = -1
var val = 0

func findBottomLeftValue(root *TreeNode) int {
	find(root, 0)
	result := val
	val = 0
	maxLevel = -1
	return result
}

func find(node *TreeNode, level int) {
	if node == nil {
		return
	}
	if level > maxLevel {
		maxLevel = level
		val = node.Val
	}
	find(node.Left, level+1)
	find(node.Right, level+1)
}
