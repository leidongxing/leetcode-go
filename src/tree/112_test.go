package tree

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//路径总和
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}

func hasPathSum1(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return pathSum(root, sum)
}

func pathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return sum == 0
	}
	if root.Left != nil && root.Right != nil {
		return pathSum(root.Left, sum-root.Val) || pathSum(root.Right, sum-root.Val)
	} else if root.Left != nil {
		return pathSum(root.Left, sum-root.Val)
	} else if root.Right != nil {
		return pathSum(root.Right, sum-root.Val)
	}
	return root.Val == sum
}

func Test_112(t *testing.T) {
	//var t1 = new(TreeNode)
	//t1.Val, t1.Left, t1.Right = 7, nil, nil
	//var t2 = new(TreeNode)
	//t2.Val, t2.Left, t2.Right = 2, nil, nil
	//var t3 = new(TreeNode)
	//t3.Val, t3.Left, t3.Right = 1, nil, nil
	//
	//var t4 = new(TreeNode)
	//t4.Val, t4.Left, t4.Right = 11, t1, t2
	//var t5 = new(TreeNode)
	//t5.Val, t5.Left, t5.Right = 13, nil, nil
	//var t6 = new(TreeNode)
	//t6.Val, t6.Left, t6.Right = 4, nil, t3
	//
	//var t7 = new(TreeNode)
	//t7.Val, t7.Left, t7.Right = 4, t4, nil
	//var t8 = new(TreeNode)
	//t8.Val, t8.Left, t8.Right = 8, t5, t6
	//
	//var t9 = new(TreeNode)
	//t9.Val, t9.Left, t9.Right = 5, t7, t8
	//
	//t.Log(hasPathSum(t9,22)) //true

	var t1 *TreeNode
	t.Log(hasPathSum(t1, 0)) //false

}
