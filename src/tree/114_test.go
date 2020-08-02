package tree

//二叉树展开为链表
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	//优先遍历左子树
	flatten(root.Left)

	//保存right
	right := root.Right
	//展开
	root.Right = root.Left
	root.Left = nil

	for root.Right != nil {
		root = root.Right
	}
	flatten(right)
	root.Right = right
}
