package tree

func isSymmetric(root *TreeNode) bool {
	return symmetric(root, root)
}

//递归方式求对称二叉树
func symmetric(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && symmetric(p.Left, q.Right) && symmetric(p.Right, q.Left)
}
