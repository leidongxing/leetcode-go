package depth_first_breadth_first

import "testing"

//二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
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
	return result
}

func Test_513(t *testing.T) {
	var t3 = new(TreeNode)
	t3.Val, t3.Left, t3.Right = 3, nil, nil
	var t2 = new(TreeNode)
	t2.Val, t2.Left, t2.Right = 2, nil, nil
	var t1 = new(TreeNode)
	t1.Val, t1.Left, t1.Right = 1, t2, t3
	t.Log(levelOrder(t1))
}
