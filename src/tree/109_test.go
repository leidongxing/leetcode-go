package tree

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//有序链表转换二叉搜索树  快慢指针
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	//中间节点为 根节点
	mid := findMiddleElement(head)
	node := &TreeNode{
		Val: mid.Val,
	}
	//只有一个节点
	if head == mid {
		return node
	}
	//左边节点 从head继续开始
	node.Left = sortedListToBST(head)
	//右边节点 从mid下一个开始
	node.Right = sortedListToBST(mid.Next)
	return node
}

//快慢指针找到中间节点
func findMiddleElement(head *ListNode) *ListNode {
	var prev *ListNode
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	if prev != nil {
		prev.Next = nil
	}
	return slow
}
