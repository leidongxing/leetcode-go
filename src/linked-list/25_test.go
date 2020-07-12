package linked_list

//K个一组翻转链表
//增加一个虚拟节点hair
func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := head
	for i := 1; i < k && cur != nil; i++ {
		cur = cur.Next
	}
	if cur == nil {
		return head
	}
	next := cur.Next
	cur.Next = nil
	cur = reverse(head)
	head.Next = reverseKGroup(next, k)
	return cur
}

func reverse(head *ListNode) *ListNode {
	//存储前一个节点  head的前节点是nil
	var pre *ListNode = nil
	//记录当前节点
	cur := head
	//当前节点不为空
	for cur != nil {
		//temp保存当前节点的后继节点
		temp := cur.Next
		//反转  当前节点的下一个节点是它的前一个节点
		cur.Next = pre

		//temp的前一个节点变为当前节点
		pre = cur
		//当前节点替换为下一个节点 也就是temp
		cur = temp
	}
	return pre
}
