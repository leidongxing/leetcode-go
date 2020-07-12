package interview

type ListNode struct {
	Val  int
	Next *ListNode
}

//删除中间节点
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
