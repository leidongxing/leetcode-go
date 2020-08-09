package depth_first_breadth_first

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//祖父节点值为偶数的节点和
//深度优先搜索
func sumEvenGrandparent(root *TreeNode) int {
	result := 0
	if root.Left != nil {
		result += dfs(root, root.Left, root.Left.Left)
		result += dfs(root, root.Left, root.Left.Right)
	}
	if root.Right != nil {
		result += dfs(root, root.Right, root.Right.Left)
		result += dfs(root, root.Right, root.Right.Right)
	}
	return result
}
func dfs(grandparent, parent, node *TreeNode) int {
	if node == nil {
		return 0
	}
	result := 0
	if grandparent.Val%2 == 0 {
		result += node.Val
	}
	return result + dfs(parent, node, node.Left) + dfs(parent, node, node.Right)
}

//递归处理
func sumEvenGrandparent1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := 0
	if root.Val%2 == 0 {
		if root.Left != nil {
			if root.Left.Left != nil {
				result += root.Left.Left.Val
			}
			if root.Left.Right != nil {
				result += root.Left.Right.Val
			}
		}

		if root.Right != nil {
			if root.Right.Left != nil {
				result += root.Right.Left.Val
			}
			if root.Right.Right != nil {
				result += root.Right.Right.Val
			}
		}
	}
	return result + sumEvenGrandparent1(root.Left) + sumEvenGrandparent1(root.Right)
}
