package interview

//移除重复节点  双重循环
func removeDuplicateNodes(head *ListNode) *ListNode {
	this := head
	for this != nil {
		tmp := this
		//循环每一个节点的所有后置节点 移除和它val一致的所有节点
		for tmp.Next != nil {
			if tmp.Next.Val == this.Val {
				tmp.Next = tmp.Next.Next
			} else {
				tmp = tmp.Next
			}
		}
		this = this.Next
	}
	return head
}

func removeDuplicateNodes1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	valMap := map[int]bool{head.Val: true}
	this := head
	for this.Next != nil {
		if valMap[this.Next.Val] {
			//直接剔除
			this.Next = this.Next.Next
		} else {
			valMap[this.Next.Val] = true
			this = this.Next
		}
	}
	//处理最后一个节点
	this.Next = nil
	return head
}
