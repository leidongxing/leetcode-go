package tree

import (
	"math"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//二叉树的最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	min := func(a int, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	depth := math.MaxInt32
	if root.Left != nil {
		depth = min(minDepth(root.Left), depth)
	}
	if root.Right != nil {
		depth = min(minDepth(root.Right), depth)
	}
	return depth + 1
}

func Test_111(t *testing.T) {
	var t1 = new(TreeNode)
	t1.Val, t1.Left, t1.Right = 2, nil, nil
	var t2 = new(TreeNode)
	t2.Val, t2.Left, t2.Right = 1, t1, nil
	t.Log(minDepth(t2))
}
